package application

import (
	"github.com/Lim0os/builder/src/application"
	"github.com/Lim0os/builder/src/application/commands"
	"log/slog"
)

func NewApp(logger *slog.Logger) application.Application {
	return application.Application{
		Commands: application.Command{
			CreateProject: commands.NewCreateProjectCommand(logger),
		},
	}
}
