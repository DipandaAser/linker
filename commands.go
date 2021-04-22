package linker

type Command struct {
	Text, Description string
}

var commands = []Command{
	{Text: "help", Description: "Display bot help."},
	{Text: "config", Description: "Get actual group/channel Linker Code" +
		" \nIf you are not an admin of the group where you call this command you can't get Linker Group/Channel Code"},
	{Text: "link", Description: "This command help you to link two groups/channels," +
		" \nall messages will be synced between this two groups/channels" +
		" \n\tExample: link yg64y iuh78" +
		" \nPlease use this command only in private message with the bot"},
	{Text: "diffuse", Description: "This command help diffuse/broadcast one Group/Channel to another Group/Channel," +
		" \nonly the message of the first group/channel will be share on the second group/channel." +
		" \n\tExample: diffuse yg64y iuh78" +
		" \nPlease use this command only in private message with the bot"},
	{Text: "start", Description: "Start an active link or diffusion group/channel, " +
		"Example: " +
		"\n\t\tstart yg64y"},
	{Text: "stop", Description: "Stop an active link or diffusion group/channel, " +
		"Example: " +
		"\n\t- Stop a specified link or diffusion" +
		"\n\t\tstop yg64y" +
		"\n" +
		"\n\t- Stop all active link or diffusion NB: use this one on a group" +
		"\n\t\tstop all"},
	{Text: "list", Description: "List all active links and diffusion of the group/channel"},
}

//GetCommands get the list of available commands
func GetCommands() []Command {
	return commands
}
