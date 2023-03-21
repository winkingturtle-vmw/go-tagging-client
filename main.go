package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/winkingturtle-vmw/go-tagging/lib"
)

func main() {
	app := lib.New("john")
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
