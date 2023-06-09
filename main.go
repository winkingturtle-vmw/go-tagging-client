package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/winkingturtle-vmw/go-tagging/lib"
)

func main() {
	lib := lib.New("john")
	lib.WithUsage("to do something")
	lib.WithUsageText("usage text")
	lib.WithArgsUsage("args")
	lib.App.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := lib.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
