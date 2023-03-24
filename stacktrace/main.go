package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	"github.com/pkg/errors"
)

type argT struct {
	Name string `cli:"name" dft:"world" usage:"tell me your name"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		fmt.Printf("%+v\n", errors.WithStack(errors.New(fmt.Sprintf("Hello, %s!\n", argv.Name))))
		return nil
	}))
}
