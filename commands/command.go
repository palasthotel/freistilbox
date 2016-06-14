package commands

type Command interface {
	Command() string
	Description() string
	Help(args []string)
	Execute(args []string)
}
