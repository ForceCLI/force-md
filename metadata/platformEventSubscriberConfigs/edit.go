package platformEventSubscriberConfigs

import (
	. "github.com/ForceCLI/force-md/general"
)

func (o *PlatformEventSubscriberConfig) UpdateUser(user string) {
	o.User = &TextLiteral{
		Text: user,
	}
}

func (o *PlatformEventSubscriberConfig) DeleteUser() {
	o.User = nil
}
