package main

import (
	"flag"

	assign "github.com/worming004/slack-assign/pkg"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var userIdToIgnore arrayFlags
	flag.Var(&userIdToIgnore, "ignore", "list of userid to ignore")
	flag.Parse()
	conf := assign.GetConfigurationByEnvironmentVariable()
	conf.UserIdToIgnore = userIdToIgnore
	a, err := assign.NewAssign(conf)
	if err != nil {
		panic(err)
	}

	err = a.Run()

	if err != nil {
		panic(err)
	}
}
