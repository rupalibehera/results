package main

import (
	"context"

	"github.com/tektoncd/results/tools/tkn-results/cmd"
	"github.com/tektoncd/results/tools/tkn-results/internal/flags"
)

func main() {
	params := &flags.Params{}
	cmd.Root(params).ExecuteContext(context.Background())
}
