package gh

func GetOrgs() []string {
	user := make(chan []string)
	orgs := make(chan []string)

	go func() {
		user <- getCommandOutput("gh", "api", "user", "-q", ".login")
	}()

	go func() {
		orgs <- getCommandOutput("gh", "org", "list")
	}()

	return append(<-user, <-orgs...)
}

func GetReposForOrg(org string) []string {
	return getCommandOutput("gh", "repo", "list", org, "--limit", "999", "--json", "nameWithOwner", "--template", "{{range .}}{{.nameWithOwner}}{{\"\\n\"}}{{end}}")
}
