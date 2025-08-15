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
	"github.com/ForceCLI/force-md/metadata/pkg"
	"github.com/ForceCLI/force-md/registry"
	"github.com/ForceCLI/force-md/repo"
)

var (
	targetDir   string
	formatType  string
	packageFile string
)

func init() {
	copyCmd.Flags().StringVarP(&targetDir, "target", "t", "", "target directory")
	copyCmd.Flags().StringVarP(&formatType, "format", "f", "", "target format (source or metadata)")
	copyCmd.Flags().StringVarP(&packageFile, "package", "x", "", "package.xml file to filter metadata")
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
  force-md copy sfdx -t src -f metadata    # Convert from source to metadata format
  force-md copy src -t sfdx -f source -x package.xml  # Convert with package.xml filter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sourceDir := args[0]

		if formatType != "source" && formatType != "metadata" {
			log.Fatal("Format must be either 'source' or 'metadata'")
		}

		if err := CopyMetadata(sourceDir, targetDir, formatType, packageFile); err != nil {
			log.Fatal(err)
		}
	},
}

func CopyMetadata(sourceDir, targetDir, format, packageFile string) error {
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

	// If package.xml is specified, filter the repo
	if packageFile != "" {
		filteredRepo, err := filterRepoByPackage(r, packageFile, targetFormat)
		if err != nil {
			return fmt.Errorf("failed to filter by package.xml: %w", err)
		}
		r = filteredRepo
	}

	// Ensure target directory exists
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// Copy package.xml and destructiveChanges files if they exist
	if packageFile != "" {
		// Copy the package.xml file to the target directory
		packageContent, err := os.ReadFile(packageFile)
		if err != nil {
			return fmt.Errorf("failed to read package.xml: %w", err)
		}
		targetPackagePath := filepath.Join(targetDir, "package.xml")
		if err := os.WriteFile(targetPackagePath, packageContent, 0644); err != nil {
			return fmt.Errorf("failed to write package.xml to target: %w", err)
		}

		// Look for destructiveChanges files in the same directory as package.xml
		packageDir := filepath.Dir(packageFile)
		entries, err := os.ReadDir(packageDir)
		if err == nil {
			for _, entry := range entries {
				if !entry.IsDir() && strings.HasPrefix(entry.Name(), "destructiveChanges") && strings.HasSuffix(entry.Name(), ".xml") {
					destructivePath := filepath.Join(packageDir, entry.Name())
					destructiveContent, err := os.ReadFile(destructivePath)
					if err == nil {
						targetDestructivePath := filepath.Join(targetDir, entry.Name())
						os.WriteFile(targetDestructivePath, destructiveContent, 0644)
					}
				}
			}
		}
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
	childTypes := []string{"CustomField", "RecordType", "ValidationRule", "Index",
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

// filterRepoByPackage filters the repo contents based on package.xml specifications
func filterRepoByPackage(r *repo.Repo, packageFile string, targetFormat metadata.Format) (*repo.Repo, error) {
	// Parse the package.xml file
	p, err := pkg.Open(packageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse package.xml: %w", err)
	}

	// Create a new filtered repo
	filteredRepo := repo.NewRepo()
	requiredComponents := make(map[string]map[string]bool) // metadataType -> componentName -> found
	hasWildcards := make(map[string]bool)                  // metadataType -> hasWildcard

	// Process package.xml types to identify wildcards and specific components
	for _, typeEntry := range p.Types {
		metadataType := typeEntry.Name

		// Special handling for Settings type
		if metadataType == "Settings" {
			// Settings in package.xml are listed as "Settings" with members like "Security"
			// But internally they're registered as "SecuritySettings", "AccountSettings", etc.
			for _, member := range typeEntry.Members {
				memberName := string(member)
				if memberName == "*" {
					// For Settings wildcard, we need to find all Settings types in the repo
					// We'll mark a special flag to process all Settings types later
					hasWildcards["__ALL_SETTINGS__"] = true
					log.Debugf("Settings wildcard found - will include all Settings types")
				} else {
					// Convert Settings.Security to SecuritySettings type
					actualMetadataType := memberName + "Settings"
					// Mark this Settings type as having a wildcard to include ALL instances of this settings type
					// (there's typically only one instance per settings type)
					hasWildcards[actualMetadataType] = true
					log.Debugf("Settings member %s converted to type %s with wildcard", memberName, actualMetadataType)
				}
			}
		} else {
			requiredComponents[metadataType] = make(map[string]bool)

			log.Debugf("Package.xml type: %s", metadataType)
			for _, member := range typeEntry.Members {
				memberName := string(member)
				log.Debugf("  Member: %s", memberName)
				if memberName == "*" {
					hasWildcards[metadataType] = true
				} else {
					requiredComponents[metadataType][memberName] = false
				}
			}
		}
	}

	// Temporary map to store items that should be included
	itemsToInclude := make(map[string]metadata.RegisterableMetadata) // path -> item

	// Process each metadata type from both the repo AND the package.xml requirements
	log.Debugf("Available metadata types: %v", r.Types())
	allTypesToProcess := make(map[string]bool)

	// Add all types from repo
	for _, metadataType := range r.Types() {
		allTypesToProcess[metadataType] = true
	}

	// Add all types from package.xml requirements
	for metadataType := range requiredComponents {
		allTypesToProcess[metadataType] = true
	}
	for metadataType := range hasWildcards {
		if metadataType != "__ALL_SETTINGS__" {
			allTypesToProcess[metadataType] = true
		}
	}

	// If we have the __ALL_SETTINGS__ wildcard, add all Settings types from the repo
	if hasWildcards["__ALL_SETTINGS__"] {
		for _, metadataType := range r.Types() {
			// Check if this is a Settings type (ends with "Settings")
			if strings.HasSuffix(metadataType, "Settings") {
				hasWildcards[metadataType] = true
				allTypesToProcess[metadataType] = true
				log.Debugf("Including Settings type %s due to Settings wildcard", metadataType)
			}
		}
	}

	log.Debugf("Processing types: %v", allTypesToProcess)

	for metadataType := range allTypesToProcess {
		typeRequirements, hasRequirements := requiredComponents[metadataType]
		hasWildcard := hasWildcards[metadataType]

		// Skip types not in package.xml
		if !hasRequirements && !hasWildcard {
			continue
		}

		log.Debugf("Processing metadata type: %s", metadataType)

		// For foldered metadata types, check if any required components are actually folder paths
		// If so, we need to include the corresponding folder metadata
		if isFolderedMetadataType(metadataType) && hasRequirements {
			folderType := getFolderTypeForMetadata(metadataType)
			log.Debugf("  Foldered type %s -> folder type %s", metadataType, folderType)
			if folderType != "" {
				folderItems := r.Items(folderType)
				log.Debugf("  Found %d folder items of type %s", len(folderItems), folderType)
				for _, folderItem := range folderItems {
					// Get the folder name for package.xml matching
					folderName := getMetadataNameForPackageXML(folderItem)
					log.Debugf("  Checking folder: %s (type: %s) against requirements", folderName, folderType)
					if _, isFolderRequired := typeRequirements[folderName]; isFolderRequired {
						// This folder is required, include the folder metadata
						folderPath := string(folderItem.GetMetadataInfo().Path())
						itemsToInclude[folderPath] = folderItem
						typeRequirements[folderName] = true
						log.Debugf("  Found required folder: %s", folderName)
					}
				}
			}
		}

		items := r.Items(metadataType)
		for _, item := range items {
			// Use the proper name for package.xml matching (includes folder for foldered metadata)
			itemName := getMetadataNameForPackageXML(item)
			itemPath := string(item.GetMetadataInfo().Path())

			// Include the item if:
			// 1. There's a wildcard for this type
			// 2. The item is specifically named
			// 3. This item is in a folder that is specifically named
			if hasWildcard {
				// Add to filtered repo
				itemsToInclude[itemPath] = item
				// Mark as found if it was also specifically named
				if _, isNamed := typeRequirements[itemName]; isNamed {
					typeRequirements[itemName] = true
				}
			} else if _, isNamed := typeRequirements[itemName]; isNamed {
				// Add specifically named item
				itemsToInclude[itemPath] = item
				typeRequirements[itemName] = true
			}

			// Handle child components for objects
			if metadataType == "CustomObject" || metadataType == "CustomObjectTranslation" {
				// When a CustomObject or CustomObjectTranslation is included, we need to handle its children
				childTypes := getChildTypesForMetadataType(metadataType)
				objectBaseName := string(item.GetMetadataInfo().Name())

				// For CustomObject/CustomObjectTranslation in metadata format, we ALWAYS need to include ALL children
				// to properly merge them into the parent, regardless of whether they're in package.xml
				if targetFormat == metadata.MetadataFormat {
					// Include ALL child components for this object/translation
					for _, childType := range childTypes {
						childItems := r.Items(childType)
						for _, childItem := range childItems {
							childName := getMetadataNameForPackageXML(childItem)
							// Check if this child belongs to the current object/translation
							if strings.HasPrefix(childName, objectBaseName+".") {
								childPath := string(childItem.GetMetadataInfo().Path())
								itemsToInclude[childPath] = childItem
							}
						}
					}
				} else {
					// For CustomObject or source format, only include children that are explicitly requested
					for _, childType := range childTypes {
						childTypeRequirements, hasChildRequirements := requiredComponents[childType]
						childHasWildcard := hasWildcards[childType]

						if !hasChildRequirements && !childHasWildcard {
							continue
						}

						childItems := r.Items(childType)
						for _, childItem := range childItems {
							childName := getMetadataNameForPackageXML(childItem)
							childPath := string(childItem.GetMetadataInfo().Path())
							// Check if this child belongs to the current object
							if strings.HasPrefix(childName, objectBaseName+".") {
								if childHasWildcard {
									itemsToInclude[childPath] = childItem
									if _, isNamed := childTypeRequirements[childName]; isNamed {
										childTypeRequirements[childName] = true
									}
								} else if _, isNamed := childTypeRequirements[childName]; isNamed {
									itemsToInclude[childPath] = childItem
									childTypeRequirements[childName] = true
								}
							}
						}
					}
				}
			}
		}
	}

	// Check for missing named components
	var missingComponents []string
	for metadataType, typeRequirements := range requiredComponents {
		for componentName, found := range typeRequirements {
			if !found {
				missingComponents = append(missingComponents, fmt.Sprintf("%s.%s", metadataType, componentName))
			}
		}
	}

	if len(missingComponents) > 0 {
		return nil, fmt.Errorf("missing required components: %s", strings.Join(missingComponents, ", "))
	}

	// Add all the included items to the filtered repo
	for _, item := range itemsToInclude {
		filteredRepo.AddItem(item)
	}

	return filteredRepo, nil
}

// getChildTypesForMetadataType returns the child metadata types for a given parent type
func getChildTypesForMetadataType(metadataType string) []string {
	switch metadataType {
	case "CustomObject":
		return []string{"CustomField", "RecordType", "ValidationRule", "Index",
			"FieldSet", "WebLink", "CompactLayout", "SharingReason", "BusinessProcess", "ListView"}
	case "CustomObjectTranslation":
		return []string{"CustomFieldTranslation", "RecordTypeTranslation", "ValidationRuleTranslation"}
	default:
		return []string{}
	}
}

// getFolderTypeForMetadata returns the folder metadata type for a foldered metadata type
func getFolderTypeForMetadata(metadataType string) string {
	switch metadataType {
	case "Dashboard":
		return "DashboardFolder"
	case "Document":
		return "DocumentFolder"
	case "EmailTemplate":
		return "EmailFolder"
	case "Report":
		return "ReportFolder"
	default:
		return ""
	}
}

// extractFolderName extracts the folder name from a foldered metadata item name
// For example, "Images/SomeFile.png" returns "Images"
func extractFolderName(itemName string) string {
	if strings.Contains(itemName, "/") {
		return strings.Split(itemName, "/")[0]
	}
	return ""
}

// isFolderedMetadataType returns true if the metadata type uses folders
func isFolderedMetadataType(metadataType string) bool {
	switch metadataType {
	case "Dashboard", "Document", "EmailTemplate", "Report":
		return true
	default:
		return false
	}
}

// isFolderMetadataType returns true if the metadata type IS a folder
func isFolderMetadataType(metadataType string) bool {
	switch metadataType {
	case "DashboardFolder", "DocumentFolder", "EmailFolder", "ReportFolder":
		return true
	default:
		return false
	}
}

// getMetadataNameForPackageXML returns the name as it should appear in package.xml
// For foldered metadata, this includes the folder path (e.g., "FolderName/ItemName")
// For folder metadata in subfolders, this includes the path (e.g., "ParentFolder/SubFolder")
func getMetadataNameForPackageXML(item metadata.RegisterableMetadata) string {
	itemName := string(item.GetMetadataInfo().Name())
	metadataType := string(item.Type())
	path := string(item.GetMetadataInfo().Path())

	// For folder metadata types (e.g., DashboardFolder), extract the folder path
	if isFolderMetadataType(metadataType) {
		// Get the directory name for this folder metadata type
		var dirName string
		switch metadataType {
		case "DashboardFolder":
			dirName = "dashboards"
		case "DocumentFolder":
			dirName = "documents"
		case "EmailFolder":
			dirName = "email"
		case "ReportFolder":
			dirName = "reports"
		}

		// Extract the folder structure from the path
		if dirName != "" && strings.Contains(path, dirName+"/") {
			// Split the path to get everything after the metadata directory
			parts := strings.Split(path, dirName+"/")
			if len(parts) > 1 {
				// Get the relative path within the metadata directory
				relativePath := parts[1]
				// Remove the file extension and -meta.xml suffix
				relativePath = strings.TrimSuffix(relativePath, "-meta.xml")
				relativePath = strings.TrimSuffix(relativePath, ".dashboardFolder-meta.xml")
				relativePath = strings.TrimSuffix(relativePath, ".reportFolder-meta.xml")
				relativePath = strings.TrimSuffix(relativePath, ".documentFolder-meta.xml")
				relativePath = strings.TrimSuffix(relativePath, ".emailFolder-meta.xml")
				relativePath = strings.TrimSuffix(relativePath, ".dashboardFolder")
				relativePath = strings.TrimSuffix(relativePath, ".reportFolder")
				relativePath = strings.TrimSuffix(relativePath, ".documentFolder")
				relativePath = strings.TrimSuffix(relativePath, ".emailFolder")

				// For folder metadata, return the full path including subfolders
				if relativePath != "" {
					return relativePath
				}
			}
		}
	}

	// For foldered metadata types, extract the folder from the path
	if isFolderedMetadataType(metadataType) {
		// Get the directory name for this metadata type
		var dirName string
		switch metadataType {
		case "Dashboard":
			dirName = "dashboards"
		case "Document":
			dirName = "documents"
		case "EmailTemplate":
			dirName = "email"
		case "Report":
			dirName = "reports"
		}

		// Extract the folder structure from the path
		if dirName != "" && strings.Contains(path, dirName+"/") {
			// Split the path to get everything after the metadata directory
			parts := strings.Split(path, dirName+"/")
			if len(parts) > 1 {
				// Get the relative path within the metadata directory
				relativePath := parts[1]
				// Remove the file extension and -meta.xml suffix
				relativePath = strings.TrimSuffix(relativePath, "-meta.xml")

				// For Document metadata, we need to find the actual document file
				// because the package.xml references the actual file, not the metadata file
				if metadataType == "Document" {
					// Remove .document suffix to get base name
					baseName := strings.TrimSuffix(relativePath, ".document")
					// The actual document file should be in the same directory
					metadataDir := filepath.Dir(path)

					// Look for files in the same directory that start with the same base name
					entries, err := os.ReadDir(metadataDir)
					if err == nil {
						baseFileName := filepath.Base(baseName)
						for _, entry := range entries {
							if !entry.IsDir() && strings.HasPrefix(entry.Name(), baseFileName+".") && !strings.HasSuffix(entry.Name(), "-meta.xml") {
								// Found the actual document file
								folderPart := filepath.Dir(baseName)
								if folderPart == "." {
									return entry.Name()
								} else {
									return folderPart + "/" + entry.Name()
								}
							}
						}
					}
					// Fallback to the base name if we can't find the actual file
					return baseName
				}

				relativePath = strings.TrimSuffix(relativePath, ".dashboard")
				relativePath = strings.TrimSuffix(relativePath, ".report")
				relativePath = strings.TrimSuffix(relativePath, ".email")

				// If there's a folder structure, use it as the name
				if strings.Contains(relativePath, "/") {
					return relativePath
				}
			}
		}
	}

	return itemName
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
