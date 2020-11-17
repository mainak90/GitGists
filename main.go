package main

import (
	"fmt"
	"github.com/mainak90/GitGists/funcs"
	"github.com/mainak90/GitGists/models"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	// define command for our client
	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch the repo details with user. [Usage]: goTool fetch user",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// Github API Logic
					var repos []models.Repo
					user := c.Args()
					var repoUrl string
					repoUrl = fmt.Sprintf("https://api.github.com/users/%s/repos", user.Get(0))
					resp := funcs.GetStats(repoUrl)
					resp.JSON(&repos)
					log.Println(repos)
				} else {
					log.Println("Please give a username. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Creates a gist from the given text. [Usage]: goTool name 'description' sample.txt",
			Action: func(c *cli.Context) error {
				if c.NArg() > 1 {
					// Github API Logic
					var args []string
					argsLocal := c.Args()
					for i := 0; i < c.NArg(); i++ {
						args = append(args, string(argsLocal.Get(i)))
					}
					var postUrl string
					postUrl = "https://api.github.com/gists"
					resp := funcs.CreateGithubGist(postUrl, args)
					log.Println(resp.String())
				} else {
					log.Println("Please give sufficient arguments. See -h to see help")
				}
				return nil
			},
		},
	}

	app.Version = "1.0"
	app.Run(os.Args)
}
