package main

import (
	"github.com/palasthotel/freistilbox/commands"
	"github.com/palasthotel/freistilbox/help"
	"github.com/palasthotel/freistilbox/version"
	"github.com/palasthotel/freistilbox/credentials"
	"github.com/palasthotel/freistilbox/api"
	"github.com/palasthotel/freistilbox/clusters"
)

func Fill() {
	commands.List=[]commands.Command{
		help.Cmd{},
		version.Cmd{},
		credentials.Cmd{},
		api.Cmd{},
		clusters.Cmd{},
	}
}
