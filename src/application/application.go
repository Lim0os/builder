package application

import "github.com/Lim0os/builder/src/application/commands"

type Command struct {
	CreateProject commands.CreateProjectCommand
}

type Application struct {
	Commands Command
}
