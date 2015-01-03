package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Repo struct {
	URL string
	Dir string
}

func main() {
	flagRepos := flag.String("repos", "", `Git repositories and aliases. `+
		`For example, "bob=https://bob:passwd@github.com/bob/project `+
		`tom=https://tom:pw@bitbucket.com/tom/work"`)
	flagWork := flag.String("work", "/tmp", `Working directory`)
	flagAddr := flag.String("addr", ":6060", "Listening address.")
	flag.Parse()

	if repos, e := ParseRepos(*flagRepos); e != nil {
		log.Fatalf("%v", e)
	} else if e := CloneRepos(repos); e != nil {
		log.Fatalf("Failed clone repos: %v", e)
	} else if h, e := MakeHandler(repos, *flagWork); e != nil {
		log.Fatalf("Failed making handler: %v", e)
	} else {
		http.Handle("/", h)
		log.Fatal(http.ListenAndServe(*flagAddr, nil))
	}
}

func ParseRepos(repos string) (map[string]Repo, error) {
	ret := make(map[string]Repo)
	segments := strings.Fields(repos)
	for _, s := range segments {
		kv := strings.Split(s, "=")
		if len(kv) == 2 && len(kv[0]) > 0 && len(kv[1]) > 0 {
			ret[kv[0]] = Repo{URL: kv[1]}
		} else {
			return nil, fmt.Errorf(`Cannot parse repo and alias: "%s"`, s)
		}
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf(`No repo and alias in "%s"`, repos)
	}
	return ret, nil
}

func CloneRepos(repos map[string]Repo) error {
	return nil
}

func MakeHandler(repos map[string]Repo, local string) (http.Handler, error) {
	return nil, nil
}
