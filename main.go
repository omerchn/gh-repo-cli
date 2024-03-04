package main

import (
	"fmt"
	"gh-repo-cli/cli"
	"gh-repo-cli/gh"
	"log"
	"os"
	"strings"
)

func main() {
	s := cli.Spinner()

	orgs := gh.GetOrgs()
	items := make([]cli.ListItem, len(orgs))
	for i, org := range orgs {
		items[i] = cli.ListItem{
			TitleText: org,
		}
	}

	s.ReleaseTerminal()

	org, err := cli.List("Select Org", items)
	if err != nil {
		log.Panic(err)
	}
	if org == "" {
		os.Exit(0)
	}

	s.RestoreTerminal()

	repos := gh.GetReposForOrg(org)
	items = make([]cli.ListItem, len(repos))
	for i, repo := range repos {
		items[i] = cli.ListItem{
			TitleText:       strings.Split(repo, "/")[1],
			DescriptionText: repo,
		}
	}

	s.ReleaseTerminal()

	repo, err := cli.List("Select Repo", items)
	if err != nil {
		log.Panic(err)
	}
	if repo == "" {
		os.Exit(0)
	}

	url := fmt.Sprintf("https://github.com/%s/%s", org, repo)

	const (
		openInBrowser = "Open in browser"
		clone         = "Clone"
	)

	action, err := cli.List(fmt.Sprintf("What to do with %s?", repo), []cli.ListItem{
		{TitleText: openInBrowser},
		{TitleText: clone},
	})
	if err != nil {
		log.Panic(err)
	}
	if action == "" {
		os.Exit(0)
	}

	if action == openInBrowser {
		fmt.Printf("open %s in browser\n", url)
	}
	if action == clone {
		fmt.Printf("clone %s\n", url)
	}
}
