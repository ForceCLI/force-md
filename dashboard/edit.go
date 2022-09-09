package dashboard

import (
	. "github.com/octoberswimmer/force-md/general"
)

func (o *Dashboard) UpdateRunningUser(user string) {
	o.RunningUser = &TextLiteral{
		Text: user,
	}
}
