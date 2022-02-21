package main

import (
	assign "github.com/worming004/slack-assign/pkg"
)

func main() {
	conf := assign.GetConfigurationByEnvironmentVariable()
	a, err := assign.NewAssign(conf)
	if err != nil {
		panic(err)
	}

	err = a.Run()

	if err != nil {
		panic(err)
	}
}
