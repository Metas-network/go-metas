package main

import (
	"fmt"
	"os"

	"github.com/Metas-network/go-metas/cmd/metas/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
