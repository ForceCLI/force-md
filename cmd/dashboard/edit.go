package dashboard

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/ForceCLI/force-md/dashboard"
	"github.com/ForceCLI/force-md/internal"
)

type DashboardType enumflag.Flag

var (
	runningUser   string
	dashboardType DashboardType
)

const (
	NoneDashboardType DashboardType = iota
	SpecifiedUser
	LoggedInUser
	MyTeamUser
)

var DashboardTypeIds = map[DashboardType][]string{
	NoneDashboardType: {"None"},
	SpecifiedUser:     {"SpecifiedUser"},
	LoggedInUser:      {"LoggedInUser"},
	MyTeamUser:        {"MyTeamUser"},
}

func init() {
	EditCmd.Flags().VarP(enumflag.New(&dashboardType, "dashboard-type", DashboardTypeIds, enumflag.EnumCaseInsensitive),
		"dashboard-type", "t", "dashboard-type; can be 'SpecifiedUser', 'LoggedInUser', or 'MyTeamUser'")
	EditCmd.Flags().StringVarP(&runningUser, "running-user", "r", "", "user dashboard runs as")
}

var EditCmd = &cobra.Command{
	Use:   "edit [flags] [filename]...",
	Short: "Edit dashboard",
	Long:  "Edit dashboard",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range args {
			updateDashboard(file)
		}
	},
}

func updateDashboard(file string) {
	a, err := dashboard.Open(file)
	if err != nil {
		log.Warn("parsing dashboard failed: " + err.Error())
		return
	}
	if runningUser != "" {
		a.UpdateRunningUser(runningUser)
	}
	switch dashboardType {
	case LoggedInUser, SpecifiedUser, MyTeamUser:
		a.UpdateDashboardType(DashboardTypeIds[dashboardType][0])
	}
	if err != nil {
		log.Warn(fmt.Sprintf("update failed for %s: %s", file, err.Error()))
		return
	}
	err = internal.WriteToFile(a, file)
	if err != nil {
		log.Warn("update failed: " + err.Error())
		return
	}
}
