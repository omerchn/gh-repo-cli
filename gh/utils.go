package gh

import (
	"log"
	"os/exec"
	"strings"
)

func filter[T any](slice []T, test func(T) bool) (ret []T) {
	for _, item := range slice {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}

func getStringLines(str string) []string {
	orgs := strings.Split(string(str), "\n")
	return filter(orgs, func(s string) bool { return s != "" })
}

func getCommandOutput(name string, arg ...string) []string {
	cmd := exec.Command(name, arg...)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}
	return getStringLines(string(out))
}
