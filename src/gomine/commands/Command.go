package commands

type Command struct {
	name string
	permission string
	aliases []string
}

func NewCommand(name string, permission string, aliases []string) *Command {
	return &Command{name, permission, aliases}
}

/**
 * Returns the command name.
 */
func (command *Command) GetName() string {
	return command.name
}

/**
 * Returns the command permission string.
 */
func (command *Command) GetPermission() string {
	return command.permission
}

/**
 * Returns the aliases of this command.
 */
func (command *Command) GetAliases() []string {
	return command.aliases
}

func (command *Command) Parse(commandText string) (string, bool) {
	return commandText, true
}
