package cmd

import "context"

func (c *Cmd) CreateApp(ctx context.Context, projectName string) error {
	err := c.app.Commands.CreateProject.Handle(ctx, projectName)
	if err != nil {
		return err
	}
	return nil
}
