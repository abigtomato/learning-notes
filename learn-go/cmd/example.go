package main

import (
	"os"
	"fmt"
	"log"
	"bufio"	
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) {
		if c.NArg() != 0 {
			fmt.Printf("未找到命令: %s\n运行命令 %s help 获取帮助\n", c.Args().Get(0), app.Name)
			return
		}

		var prompt string
		prompt = app.Name + " > " 
		
		L:
		for {
			fmt.Println(prompt)

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input := scanner.Text()

			switch input {
			case "close":
				fmt.Println("close.")
				break L
			default:
			}
 
			cmdArgs := strings.Split(input, " ")
			if len(cmdArgs) == 0 {
				continue
			}

			s := []string{app.Name}
			s = append(s, cmdArgs...)

			c.App.Run(s)
		}

		return
	}

	app.Commands = []cli.Command{
		{
			Name: 	 "add",
			Aliases: []string{"a"},
			Usage: 	 "add a task to the list",
			Action:  func(c *cli.Context) error {
				fmt.Printf("addedh task: %q\n", c.Args().Get(0))
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
