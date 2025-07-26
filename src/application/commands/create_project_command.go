package commands

import (
	"context"
	"fmt"
	"github.com/Lim0os/builder/src/domain"
	"log/slog"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Lim0os/builder/src/common/decorator"
)

type createProjectCommand struct {
	logger   *slog.Logger
	template *ProjectTemplate
}

type CreateProjectCommand decorator.CommandDecorator[string]

type ProjectTemplate struct {
	Files map[string]string
}

func NewProjectTemplate(projectName string) *ProjectTemplate {
	return &ProjectTemplate{
		Files: map[string]string{
			"src/application/app.go":                           domain.ApplicationAppGoTemplate,
			"src/domain/entity.go":                             domain.DomainEntityGoTemplate,
			"src/domain/repository.go":                         domain.DomainRepositoryGoTemplate,
			"src/ports_adapters/primary/http_server/server.go": domain.HttpServerGoTemplate,
			"go.mod":     domain.GoModTemplate,
			"main.go":    domain.MainGoTemplate,
			"Makefile":   domain.MakefileTemplate,
			"Dockerfile": domain.DockerfileTemplate,
			"README.md":  domain.ReadmeTemplate,
			".gitignore": domain.GitignoreTemplate,
		},
	}
}

func NewCreateProjectCommand(logger *slog.Logger) decorator.CommandDecorator[string] {
	handler := createProjectCommand{
		logger: logger,
	}
	return decorator.ApplyCommandDecorator[string](handler, logger)
}

func (c createProjectCommand) Handle(ctx context.Context, projectName string) error {
	defer ctx.Done()
	if err := validateProjectName(projectName); err != nil {
		return fmt.Errorf("invalid project name: %w", err)
	}

	c.template = NewProjectTemplate(projectName)

	dirs := []string{
		"src/application/commands",
		"src/common/config",
		"src/common/decorator",
		"src/domain",
		"src/ports_adapters/primary/http_server/dto",
		"src/ports_adapters/secondary/persistence",
		"src/ports_adapters/secondary/service",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(projectName, dir), 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	for filePath, content := range c.template.Files {
		fullPath := filepath.Join(projectName, filePath)
		if err := generateFileFromTemplate(fullPath, content, projectName); err != nil {
			return fmt.Errorf("failed to generate file %s: %w", filePath, err)
		}
	}

	c.logger.Info("project created successfully", "name", projectName)
	return nil
}

func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	return nil
}

func generateFileFromTemplate(path, content string, data interface{}) error {
	tmpl, err := template.New(filepath.Base(path)).Parse(content)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}
