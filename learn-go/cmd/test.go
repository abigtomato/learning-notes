package main

import (
	"os"
	"fmt"
	"log"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	
	var language string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "lang",
			Value: "english",
			Usage: "language for the greeting",
			Destination: &language,
		},
	}
	
	app.Commands = []cli.Command{
		{
			Name: 	 "add",
			Aliases: []string{"a"},
			Usage: 	 "add a task to the list",
			Action:  func(c *cli.Context) error {
				fmt.Println(c.Command.Name)
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action:  func(c *cli.Context) error {
				fmt.Printf("completed task: %q\n", c.Args().First())
				return nil
			},
		}, 
		{	
			Name:        "template",
			Aliases:     []string{"t"},
			Usage:       "options go r task templates",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		}, 
	}

	app.Action = func(c * cli.Context) error {
		// name := "Nefertiti"

		// if c.NArg() > 0 {
		// 	name = c.Args().Get(0)
		// }

		/*if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}*/

		// switch language {
		// case "chinese":
		// 	fmt.Println("你好", name)
		// case "spanish":
		// 	fmt.Println("Hola", name)
		// default:
		// 	fmt.Println("Hello", name)
		// }

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
