package sopt

// Command definition.
type Command struct {
	// Name of the command.
	Name string
	// Help text of the command.
	Help string
	// Func to execute the command.
	Func ToolCommand
	// Options for this command.
	Options []*Option
	// Aliases for this command.
	Aliases []string
}

// ToolCommand function signature.
type ToolCommand func(args []string) error

// SetCommand to a group.
func (opt *Options) SetCommand(name, help, group string, fn ToolCommand, aliases []string) *Command {
	cmd := &Command{
		Name:    name,
		Help:    help,
		Func:    fn,
		Aliases: aliases,
	}

	opt.commands[name] = cmd
	g := opt.GetGroup(group)
	if g == nil {
		g = opt.AddGroup(group)
	}
	g.commands = append(g.commands, cmd.Name)
	return cmd
}
