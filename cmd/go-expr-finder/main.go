package main

import (
	"context"

	"github.com/torafugu2929/go-expr-finder/cli"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cli.RootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}
