package main

import (
	"fmt"

	"github.com/Masterminds/cookoo"
	"github.com/pkg/errors"
)

func main() {
	// Build a new Cookoo app.
	registry, router, context := cookoo.Cookoo()
	// Fill the registry.
	registry.AddRoutes(
		cookoo.Route{
			Name: "TEST",
			Help: "A test route",
			Does: cookoo.Tasks{
				cookoo.Cmd{
					Name: "hi",
					Fn:   HelloWorld,
				},
			},
		},
	)
	// Execute the route.
	router.HandleRequest("TEST", context, false)
}

func HelloWorld(cxt cookoo.Context, params *cookoo.Params) (interface{}, cookoo.Interrupt) {
	fmt.Printf("%+v\n", errors.New("Hello World"))
	return true, nil
}
