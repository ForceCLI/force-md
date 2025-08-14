package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/general"
	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/metadata"
	"github.com/ForceCLI/force-md/metadata/objectTranslations"
	"github.com/ForceCLI/force-md/metadata/objects"
	"github.com/ForceCLI/force-md/registry"
	"github.com/ForceCLI/force-md/repo"
)

var (
	targetDir  string
	formatType string
)

func init() {
	copyCmd.Flags().StringVarP(&targetDir, "target", "t", "", "target directory")
	copyCmd.Flags().StringVarP(&formatType, "format", "f", "", "target format (source or metadata)")
	copyCmd.MarkFlagRequired("target")
	copyCmd.MarkFlagRequired("format")
	RootCmd.AddCommand(copyCmd)
}

var copyCmd = &cobra.Command{
	Use:   "copy [source directory]",
	Short: "Copy metadata between source and metadata formats",
	Long: `Copy and convert metadata between source format (SFDX) and metadata format (MDAPI).

Examples:
  force-md copy src -t sfdx -f source      # Convert from metadata to source format
  force-md copy sfdx -t src -f metadata    # Convert from source to metadata format`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := args[0]

		if formatType != "source" && formatType != "metadata" {
			log.Fatal("Format must be either 'source' or 'metadata'")
		}

		if err := copyMetadata(sourceDir, targetDir, formatType); err != nil {
			log.Fatal(err)
		}
	},
}

func copyMetadata(sourceDir, targetDir, format string) error {
	// Convert format string to metadata.Format
	var targetFormat metadata.Format
	switch format {
	case "source":
		targetFormat = metadata.SourceFormat
	case "metadata":
		targetFormat = metadata.MetadataFormat
	default:
		return fmt.Errorf("invalid format: %s", format)
	}

	// Create a new repo and load all metadata from source directory
	r := repo.NewRepo()

	// Walk the source directory and add all metadata to the repo
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Try to open as metadata
		_, err = r.Open(path)
		if err != nil {
			// Not a metadata file, skip it
			log.Debugf("Skipping %s: %v", path, err)
			return nil
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	// Process metadata based on target format
	if targetFormat == metadata.MetadataFormat {
		// For metadata format, we need to handle CustomObject merging
		return writeMetadataFormat(r, targetDir)
	} else {
		// For source format, process each type individually
		return writeSourceFormat(r, targetDir)
	}
}

// writeSourceFormat writes all metadata in source format
func writeSourceFormat(r *repo.Repo, targetDir string) error {
	for _, metadataType := range r.Types() {
		items := r.Items(metadataType)
		for _, metadataItem := range items {
			name := metadataItem.GetMetadataInfo().Name()
			files, err := getMetadataFiles(metadataItem, name, metadata.SourceFormat)
			if err != nil {
				return fmt.Errorf("failed to get files for %s %s: %w", metadataType, name, err)
			}

			// Write each file
			for relativePath, content := range files {
				targetPath := filepath.Join(targetDir, relativePath)

				// Create the directory if needed
				targetFileDir := filepath.Dir(targetPath)
				if err := os.MkdirAll(targetFileDir, 0755); err != nil {
					return fmt.Errorf("failed to create directory %s: %w", targetFileDir, err)
				}

				// Write the file
				if err := os.WriteFile(targetPath, content, 0644); err != nil {
					return fmt.Errorf("failed to write file %s: %w", targetPath, err)
				}

				log.Debugf("Wrote %s", targetPath)
			}
		}
	}
	return nil
}

// writeMetadataFormat writes all metadata in metadata format (no -meta.xml suffix, merged objects)
func writeMetadataFormat(r *repo.Repo, targetDir string) error {
	// First, handle CustomObjects - merge components back together
	processedObjects := make(map[string]bool)

	// First, collect all object names from child components
	allObjectNames := make(map[string]bool)

	// Check existing CustomObject items
	objectItems := r.Items("CustomObject")
	for _, item := range objectItems {
		objNameStr := string(item.GetMetadataInfo().Name())
		allObjectNames[objNameStr] = true
	}

	// Check all child component types to find objects without base metadata
	childTypes := []string{"CustomField", "RecordType", "ValidationRule", "BigObjectIndex",
		"FieldSet", "WebLink", "CompactLayout", "SharingReason", "BusinessProcess", "ListView"}

	for _, childType := range childTypes {
		items := r.Items(childType)
		for _, item := range items {
			itemNameStr := string(item.GetMetadataInfo().Name())
			// Extract object name from component name (e.g., "Account.Field1" -> "Account")
			if dotIndex := strings.Index(itemNameStr, "."); dotIndex > 0 {
				objName := itemNameStr[:dotIndex]
				allObjectNames[objName] = true
			}
		}
	}

	// Process all objects (including those with only child components)
	for objBaseName := range allObjectNames {
		if processedObjects[objBaseName] {
			continue
		}
		processedObjects[objBaseName] = true

		// Use the ComposeFromChildren helper to create the merged object
		mergedObj := objects.ComposeFromChildren(objBaseName, r)

		// Write merged object
		files, err := mergedObj.Files(metadata.MetadataFormat)
		if err != nil {
			return fmt.Errorf("failed to get files for CustomObject %s: %w", objBaseName, err)
		}

		for relativePath, content := range files {
			targetPath := filepath.Join(targetDir, relativePath)
			targetFileDir := filepath.Dir(targetPath)
			if err := os.MkdirAll(targetFileDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetFileDir, err)
			}
			if err := os.WriteFile(targetPath, content, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
			log.Debugf("Wrote %s", targetPath)
		}
	}

	// Handle CustomObjectTranslation - merge field translations back together
	processedTranslations := make(map[string]bool)

	// First, collect all object translation names from both base translations and field translations
	allObjTransNames := make(map[string]bool)

	// Check existing CustomObjectTranslation items
	objectTranslationItems := r.Items("CustomObjectTranslation")
	for _, item := range objectTranslationItems {
		objTransNameStr := string(item.GetMetadataInfo().Name())
		allObjTransNames[objTransNameStr] = true
	}

	// Check all child translations to find object translations without base metadata
	childTranslationTypes := []string{"CustomFieldTranslation", "RecordTypeTranslation", "ValidationRuleTranslation"}
	for _, childType := range childTranslationTypes {
		childTranslations := r.Items(childType)
		for _, childMeta := range childTranslations {
			childNameStr := string(childMeta.GetMetadataInfo().Name())
			// Extract object translation name from child translation name (e.g., "Account-en_US.Field1" -> "Account-en_US")
			if dotIndex := strings.Index(childNameStr, "."); dotIndex > 0 {
				objTransBaseName := childNameStr[:dotIndex]
				allObjTransNames[objTransBaseName] = true
			}
		}
	}

	// Process all object translations (including those with only field translations)
	for objTransBaseName := range allObjTransNames {
		if processedTranslations[objTransBaseName] {
			continue
		}
		processedTranslations[objTransBaseName] = true

		// Use the ComposeFromChildren helper to create the merged object translation
		mergedObjTrans := objectTranslations.ComposeFromChildren(objTransBaseName, r)

		// Write merged object translation using default Files() method
		files, err := getMetadataFiles(mergedObjTrans, metadata.MetadataObjectName(objTransBaseName), metadata.MetadataFormat)
		if err != nil {
			return fmt.Errorf("failed to get files for CustomObjectTranslation %s: %w", objTransBaseName, err)
		}

		for relativePath, content := range files {
			targetPath := filepath.Join(targetDir, relativePath)
			targetFileDir := filepath.Dir(targetPath)
			if err := os.MkdirAll(targetFileDir, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetFileDir, err)
			}
			if err := os.WriteFile(targetPath, content, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", targetPath, err)
			}
			log.Debugf("Wrote %s", targetPath)
		}
	}

	// Process non-object and non-child types
	for _, metadataType := range r.Types() {
		if metadataType != "CustomObject" && metadataType != "CustomObjectTranslation" && !repo.IsChildType(metadataType) {
			// For non-objects and non-child types, write normally using AllItems to avoid name collisions
			items := r.Items(metadataType)
			for _, metadataItem := range items {
				name := metadataItem.GetMetadataInfo().Name()

				// If the metadata item implements Tidyable, tidy it before writing
				if tidyable, ok := metadataItem.(general.Tidyable); ok {
					tidyable.Tidy()
				}

				files, err := getMetadataFiles(metadataItem, name, metadata.MetadataFormat)
				if err != nil {
					return fmt.Errorf("failed to get files for %s %s: %w", metadataType, name, err)
				}

				for relativePath, content := range files {
					targetPath := filepath.Join(targetDir, relativePath)
					targetFileDir := filepath.Dir(targetPath)
					if err := os.MkdirAll(targetFileDir, 0755); err != nil {
						return fmt.Errorf("failed to create directory %s: %w", targetFileDir, err)
					}
					if err := os.WriteFile(targetPath, content, 0644); err != nil {
						return fmt.Errorf("failed to write file %s: %w", targetPath, err)
					}
					log.Debugf("Wrote %s", targetPath)
				}
			}
		}
	}

	return nil
}

// getMetadataFiles returns the files for a metadata item, either from its Files() method or a default implementation
func getMetadataFiles(m metadata.RegisterableMetadata, name metadata.MetadataObjectName, format metadata.Format) (map[string][]byte, error) {
	// Check if this metadata implements the FilesGenerator interface
	if generator, ok := m.(metadata.FilesGenerator); ok {
		return generator.Files(format)
	}

	// Default implementation for types without Files() method
	return getDefaultFiles(m, name, format)
}

// getDefaultFiles provides a default implementation for metadata types that don't implement FilesGenerator
func getDefaultFiles(m metadata.RegisterableMetadata, name metadata.MetadataObjectName, format metadata.Format) (map[string][]byte, error) {
	metadataType := m.Type()
	dirName := registry.GetCanonicalDirectoryName(metadataType)

	// Marshal the metadata to XML using internal.Marshal to get proper formatting
	xmlContent, err := internal.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s: %w", metadataType, err)
	}

	files := make(map[string][]byte)
	fileName := string(name)

	switch format {
	case metadata.SourceFormat:
		// Source format: add -meta.xml suffix
		if !strings.HasSuffix(fileName, "-meta.xml") {
			// Add the appropriate extension and -meta.xml
			ext := registry.GetMetadataSuffix(metadataType)
			if ext != "" {
				fileName = strings.TrimSuffix(fileName, "."+ext)
				fileName = fileName + "." + ext + "-meta.xml"
			} else {
				fileName = fileName + "-meta.xml"
			}
		}
		files[filepath.Join(dirName, fileName)] = xmlContent

	case metadata.MetadataFormat:
		// Metadata format: no -meta.xml suffix
		if strings.HasSuffix(fileName, "-meta.xml") {
			fileName = strings.TrimSuffix(fileName, "-meta.xml")
		}
		// Ensure proper extension for metadata format
		ext := registry.GetMetadataSuffix(metadataType)
		if ext != "" && !strings.HasSuffix(fileName, "."+ext) {
			fileName = fileName + "." + ext
		}
		files[filepath.Join(dirName, fileName)] = xmlContent

	default:
		return nil, fmt.Errorf("unsupported format: %v", format)
	}

	return files, nil
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
