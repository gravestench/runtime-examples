package twitch_soundboard

import (
	"github.com/gravestench/runtime"

	"github.com/gravestench/runtime-examples/services/config_file"
	"github.com/gravestench/runtime-examples/services/desktop_notification"
)

func (s *Service) ResolveDependencies(rt runtime.R) {
	for _, service := range rt.Services() {
		if instance, ok := service.(config_file.Manager); ok {
			s.configManager = instance
		}

		if instance, ok := service.(desktop_notification.SendsNotifications); ok {
			s.notification = instance
		}
	}
}

func (s *Service) DependenciesResolved() bool {
	if s.configManager == nil {
		return false
	}

	return true
}
