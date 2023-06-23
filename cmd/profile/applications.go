package profile

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/ForceCLI/force-md/internal"
	"github.com/ForceCLI/force-md/profile"
)

var (
	applicationName    string
	isDefault, visible bool
)

func init() {
	deleteApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	deleteApplicationCmd.MarkFlagRequired("application")

	addApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	addApplicationCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "make default application")
	addApplicationCmd.MarkFlagRequired("application")

	listApplicationsCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "default applications only")
	listApplicationsCmd.Flags().BoolVarP(&visible, "visible", "v", false, "visible applications only")

	showApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	showApplicationCmd.MarkFlagRequired("application")

	editApplicationCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	editApplicationCmd.Flags().BoolP("default", "d", false, "is default application")
	editApplicationCmd.Flags().BoolP("visible", "v", false, "is visible")
	editApplicationCmd.Flags().BoolP("no-default", "D", false, "is not default")
	editApplicationCmd.Flags().BoolP("no-visible", "V", false, "is not visible")
	editApplicationCmd.MarkFlagRequired("application")

	tableApplicationsCmd.Flags().StringVarP(&applicationName, "application", "a", "", "application name")
	tableApplicationsCmd.Flags().BoolP("default", "d", false, "is default application")
	tableApplicationsCmd.Flags().BoolP("visible", "v", false, "is visible")
	tableApplicationsCmd.Flags().BoolP("no-default", "D", false, "is not default")
	tableApplicationsCmd.Flags().BoolP("no-visible", "V", false, "is not visible")

	ApplicationCmd.AddCommand(addApplicationCmd)
	ApplicationCmd.AddCommand(deleteApplicationCmd)
	ApplicationCmd.AddCommand(editApplicationCmd)
	ApplicationCmd.AddCommand(listApplicationsCmd)
	ApplicationCmd.AddCommand(showApplicationCmd)
	ApplicationCmd.AddCommand(tableApplicationsCmd)
}

var ApplicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage application visibility",
}

var deleteApplicationCmd = &cobra.Command{
	Use:   "delete -a ApplicationName [flags] [filename]...",
	Short: "Delete application visibility",
	Long:  "Delete application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			deleteApplicationVisibility(file, applicationName)
		}
	},
}

var showApplicationCmd = &cobra.Command{
	Use:   "show -a ApplicationName [filename]...",
	Short: "Show application visibility",
	Long:  "Show application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			showApplicationVisibility(file, applicationName)
		}
	},
}

var editApplicationCmd = &cobra.Command{
	Use:   "edit -a ApplicationName [flags] [filename]...",
	Short: "Update application visibility",
	Long:  "Update application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := applicationVisibiltyFromFlags(cmd)
		for _, file := range args {
			updateApplicationVisibility(file, perms)
		}
	},
}

var addApplicationCmd = &cobra.Command{
	Use:   "add -a ApplicationName [flags] [filename]...",
	Short: "Add application visibility",
	Long:  "Add application visibility in profiles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			addApplicationVisibility(file, applicationName, isDefault)
		}
	},
}

var listApplicationsCmd = &cobra.Command{
	Use:   "list [flags] [filename]...",
	Short: "List applications assigned",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			listApplications(file)
		}
	},
}

var tableApplicationsCmd = &cobra.Command{
	Use:   "table [flags] [filename]...",
	Short: "List Applications in a table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		perms := applicationVisibiltyFromFlags(cmd)
		tableApplications(args, perms)
	},
}

func deleteApplicationVisibility(file string, applicationName string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.DeleteApplicationVisibility(applicationName)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func addApplicationVisibility(file string, applicationName string, isDefault bool) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.AddApplicationVisibility(applicationName, isDefault)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func listApplications(file string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	var filters []profile.ApplicationFilter
	if isDefault {
		filters = append(filters, func(v profile.ApplicationVisibility) bool {
			return v.Default.Text == "true"
		})
	}
	if visible {
		filters = append(filters, func(v profile.ApplicationVisibility) bool {
			return v.Visible.Text == "true"
		})
	}
	apps := p.GetApplications(filters...)
	for _, a := range apps {
		defaultString := ""
		visibleString := "not visible"
		if a.Default.Text == "true" {
			defaultString = " (default)"
		}
		if a.Visible.Text == "true" {
			visibleString = "visible"
		}
		fmt.Printf("%s%s: %s\n", a.Application, defaultString, visibleString)
	}
}

func showApplicationVisibility(file string, application string) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile set failed: " + err.Error())
		return
	}
	filter := func(v profile.ApplicationVisibility) bool {
		return strings.ToLower(v.Application) == strings.ToLower(application)
	}
	apps := p.GetApplications(filter)
	b, err := xml.MarshalIndent(apps[0], "", "    ")
	if err != nil {
		log.Warn("marshal failed: " + err.Error())
		return
	}
	fmt.Println(string(b))
}

func tableApplications(files []string, toApply profile.ApplicationVisibility) {
	var filters []profile.ApplicationFilter
	if applicationName != "" {
		filters = append(filters, func(f profile.ApplicationVisibility) bool {
			return strings.ToLower(f.Application) == strings.ToLower(applicationName)
		})
	}
	if toApply.Visible.Text != "" {
		filters = append(filters, func(f profile.ApplicationVisibility) bool {
			return toApply.Visible.ToBool() == f.Visible.ToBool()
		})
	}
	if toApply.Default.Text != "" {
		filters = append(filters, func(f profile.ApplicationVisibility) bool {
			return toApply.Default.ToBool() == f.Default.ToBool()
		})
	}
	type perm struct {
		apps    profile.ApplicationVisibilityList
		profile string
	}
	var perms []perm
	for _, file := range files {
		p, err := profile.Open(file)
		if err != nil {
			log.Warn("parsing profile failed: " + err.Error())
			return
		}
		profileName := strings.TrimSuffix(path.Base(file), ".profile")
		perms = append(perms, perm{apps: p.GetApplications(filters...), profile: profileName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Application", "Visible", "Default"})
	table.SetRowLine(true)
	for _, perm := range perms {
		for _, t := range perm.apps {
			table.Append([]string{perm.profile, t.Application,
				t.Visible.String(),
				t.Default.String(),
			})
		}
	}
	if table.NumLines() > 0 {
		table.Render()
	}
}

func updateApplicationVisibility(file string, perms profile.ApplicationVisibility) {
	p, err := profile.Open(file)
	if err != nil {
		log.Warn("parsing profile failed: " + err.Error())
		return
	}
	err = p.SetApplicationVisibility(applicationName, perms)
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(p, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}

func applicationVisibiltyFromFlags(cmd *cobra.Command) profile.ApplicationVisibility {
	perms := profile.ApplicationVisibility{}
	perms.Default = textValue(cmd, "default")
	perms.Visible = textValue(cmd, "visible")
	return perms
}
