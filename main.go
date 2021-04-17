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
			Usage:   "Creates a gist from the given text. [Usage]: goTool create <name> <description> <sample.txt>",
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
		{
			Name:    "newrepo",
			Aliases: []string{"n"},
			Usage:   "create new repo. [Usage]: goTool newrepo <repo-name>",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// Github API Logic
					argset := c.Args()
					var repoUrl string
					repoUrl = fmt.Sprintf("https://api.github.com/user/repos")
					log.Println("Repourl: ", repoUrl)
					reponame := fmt.Sprintf("%s", argset.Get(0))
					log.Println("Reponame: ", reponame)
					resp := funcs.NewRepo(repoUrl, reponame)
					log.Println(resp)
				} else {
					log.Println("Please give a username and a reponame. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "newrepos",
			Aliases: []string{"s"},
			Usage:   "create a file of repos. [Usage]: goTool newrepos <repo-file-path>",
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					// Github API Logic
					argset := c.Args()
					var repoUrl string
					repoUrl = fmt.Sprintf("https://api.github.com/user/repos")
					log.Println("Repourl: ", repoUrl)
					repopath := fmt.Sprintf("%s", argset.Get(0))
					repolist := funcs.MakeList(repopath)
					repos := funcs.RepoList{Repos: repolist}
					repos.NewRepos(repoUrl)
				} else {
					log.Println("Please give a username and a reponame. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "pushwebhook",
			Aliases: []string{"p"},
			Usage:   "Push a webhook into an existing repo. [Usage]: goTool pushwebhook <owner> <repo> <payload>",
			Action: func(c *cli.Context) error {
				if c.NArg() == 3 {
					// Github API Logic
					argset := c.Args()
					repoUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", argset.Get(0), argset.Get(1))
					log.Println("Repourl: ", repoUrl)
					log.Println("Webhook: ", argset.Get(2))
					var c funcs.HookRequest
					inited:= c.InitConfig(argset.Get(2), argset.Get(0), argset.Get(1))
					inited.CreateWebHook(repoUrl)
				} else {
					log.Println("Please provide a <owner> <repo> <payload>. See -h to see help")
				}
				return nil
			},
		},
		{
			Name:    "pushwebhooks",
			Aliases: []string{"ps"},
			Usage:   "Push a json object of multiple webhooks into an existing repo. [Usage]: goTool pushwebhooks <path>",
			Action: func(c *cli.Context) error {
				if c.NArg() == 1 {
					// Github API Logic
					argset := c.Args()
					var c funcs.HookRequest
					c.CreateWebHooks(argset.Get(0))
				} else {
					log.Println("Please provide a <owner> <repo> <payload>. See -h to see help")
				}
				return nil
			},
		},
	}

	app.Version = "1.0"
	app.Run(os.Args)
}
