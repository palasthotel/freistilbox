package main

import (
	"github.com/palasthotel/freistilbox/commands"
	"github.com/palasthotel/freistilbox/help"
	"github.com/palasthotel/freistilbox/version"
)

func Fill() {
	commands.List=[]commands.Command{
		help.Cmd{},
		version.Cmd{},
	}
}
