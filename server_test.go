package main

import (
	"reflect"
	"testing"
)

func TestParseRepos(t *testing.T) {
	if repos, e := ParseRepos(""); repos != nil || e == nil {
		t.Errorf("Expecting error on empty input, got parsed repos %v", repos)
	}

	if repos, e := ParseRepos("a"); repos != nil || e == nil {
		t.Errorf("Expecting error on empty input, got parsed repos %v", repos)
	}

	if repos, e := ParseRepos("a="); repos != nil || e == nil {
		t.Errorf("Expecting error on empty input, got parsed repos %v", repos)
	}

	if repos, e := ParseRepos("=b"); repos != nil || e == nil {
		t.Errorf("Expecting error on empty input, got parsed repos %v", repos)
	}

	alias1 := "tom"
	repo1 := `https://tom:pw@github.com/tom/project`
	alias2 := "bob"
	repo2 := `https://bob:pw2@bitbucket.com/bob/work`
	input := alias1 + "=" + repo1 + "\t " + alias2 + "=" + repo2
	output := map[string]Repo{
		alias1: Repo{URL: repo1},
		alias2: Repo{URL: repo2},
	}
	if repos, e := ParseRepos(input); repos == nil || e != nil {
		t.Errorf("Expecting error on empty input, got parsed repos %v", repos)
	} else if !reflect.DeepEqual(repos, output) {
		t.Errorf(`Expecting %v, got %v`, output, repos)
	}

}
