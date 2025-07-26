package cmd

import (
	"context"
	"github.com/Lim0os/builder/src/application"
	"reflect"
)

type Cmd struct {
	app application.Application
}

func NewCmd(app application.Application) *Cmd {
	return &Cmd{app: app}
}

type CommandHandler[T any] interface {
	Handle(ctx context.Context, arg T) error
}

func (c *Cmd) Execute(methodName string, ctx context.Context, arg any) {
	method := reflect.ValueOf(c).MethodByName(methodName)
	if !method.IsValid() {
		panic("method not found")
	}

	result := method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(arg),
	})

	if err := result[0].Interface(); err != nil {
		panic(err.(error))
	}
}
