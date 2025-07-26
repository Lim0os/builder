package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Lim0os/builder/src/ports_adapters/primary/cmd"
	"github.com/Lim0os/builder/src/ports_adapters/service/application"
	"log/slog"
	"os"

	"github.com/Lim0os/builder/src/application/commands"
	"github.com/Lim0os/builder/src/common/config"
)

type CommandsArgs struct {
	ProjectName string
}

const (
	LvlDebug      = "debug"
	CommandCreate = "create"
)

func main() {
	logger := config.InitLogger(LvlDebug)
	args := parseCommandLineArgs()

	ctx := context.WithValue(context.Background(), "logger", logger)
	app := application.NewApp(logger)
	cmdCommand := cmd.NewCmd(app)

	commandRouter := map[string]func(context.Context, string) error{
		CommandCreate: cmdCommand.CreateApp,
	}

	if len(flag.Args()) < 1 {
		logger.Error("command is required")
		os.Exit(1)
	}
	command := flag.Arg(0)

	if handler, exists := commandRouter[command]; exists {

		if err := handler(ctx, args.ProjectName); err != nil {
			logger.Error("command failed", "error", err)
			os.Exit(1)
		}
	} else {
		logger.Error("unknown command", "command", command)
		os.Exit(1)
	}

	logger.Info("command executed successfully", "command", command)
}

func parseCommandLineArgs() CommandsArgs {
	args := CommandsArgs{}
	flag.Parse()

	if len(flag.Args()) >= 2 {
		args.ProjectName = flag.Arg(1)
	}
	return args
}

func handleCreateCommand(ctx context.Context, args CommandsArgs) error {
	logger := ctx.Value("logger").(*slog.Logger)

	if args.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}

	logger.Info("creating project", "name", args.ProjectName)

	cmd := commands.NewCreateProjectCommand(logger)
	return cmd.Handle(ctx, args.ProjectName)
}
