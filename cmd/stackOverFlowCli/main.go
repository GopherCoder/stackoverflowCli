package main

import (
	"fmt"
	"log"
	"os"
	"stackOverFlowClient/app/domain/repository"
	"stackOverFlowClient/app/infrastructure"
	"stackOverFlowClient/app/uerInterface"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "stackoverflowcli"
	app.Usage = "An application for stackoverflow web site"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}
	app.CommandNotFound = func(context *cli.Context, s string) {
		fmt.Println("command not found")
	}
	app.Commands = commandStorager()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func commandStorager() []cli.Command {
	return []cli.Command{
		{
			Name:        "user",
			Usage:       "User Commands",
			Subcommands: userSubCommand(),
		},
		{
			Name:        "job",
			Usage:       "Job Commands",
			Subcommands: jobSubCommand(),
		},
		{
			Name:        "tag",
			Usage:       "Tag Commands",
			Subcommands: tagSubCommand(),
		},
		{
			Name:        "question",
			Usage:       "Question Commands",
			Subcommands: questionSubCommand(),
		},
	}
}

func userSubCommand() []cli.Command {
	return []cli.Command{
		{
			Name:    "editor",
			Aliases: []string{"e"},
			Usage:   "Show stackoverflow user editor",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var editor uerInterface.Editors
				repo = &editor
				repo.FindAll(infrastructure.API["editors"])
			},
		}, {
			Name:    "moderators",
			Aliases: []string{"m"},
			Usage:   "Show stackoverflow user moderators",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var moderators uerInterface.Moderators
				repo = &moderators
				repo.FindAll(infrastructure.API["moderators"])
			},
		}, {
			Name:    "reputation",
			Aliases: []string{"r"},
			Usage:   "Show stackoverflow user reputation",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var reputation uerInterface.Reputation
				repo = &reputation
				repo.FindAll(infrastructure.API["reputation"])
			},
		}, {
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Show stackoverflow user userNew",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var userNew uerInterface.UserNew
				repo = &userNew
				repo.FindAll(infrastructure.API["newUser"])
			},
		}, {
			Name:    "voters",
			Aliases: []string{"v"},
			Usage:   "Show stackoverflow user voters",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var voters uerInterface.Voters
				repo = &voters
				repo.FindAll(infrastructure.API["voters"])
			},
		},
	}

}
func jobSubCommand() []cli.Command {
	return []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Show stackoverflow jobs",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var job uerInterface.Job
				repo = &job
				repo.FindAll(infrastructure.API["jobs"])
			},
		},
	}

}
func tagSubCommand() []cli.Command {
	return []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Show stackoverflow tag",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var tags uerInterface.Tags
				repo = &tags
				repo.FindAll(infrastructure.API["tags"])
			},
		},
	}

}
func questionSubCommand() []cli.Command {
	return []cli.Command{
		{
			Name:    "search",
			Aliases: []string{"q"},
			Usage:   "Show stackoverflow question",
			Action: func(c *cli.Context) {
				var repo repository.StackOverFlowRepository
				var question uerInterface.Question
				repo = &question
				repo.FindAll(fmt.Sprintf(infrastructure.API["search"], c.Args().Get(0)))
			},
		},
	}

}
