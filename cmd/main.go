package main

import (
	assign "github.com/worming004/slack-assign/pkg"
)

func main() {
	conf := assign.GetConfigurationByEnvironmentVariable()
	a := assign.NewAssign(conf)

	err := a.Run()

	if err != nil {
		panic(err)
	}
}
