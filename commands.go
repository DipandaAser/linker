package linker

type Command struct {
	Text, Description string
}

var commands = []Command{
	{Text: "help", Description: "Display bot help."},
	{Text: "config", Description: "Get actual group/channel Linker Code" +
		" \nIf you are not an admin of the group where you call this command you can't get Linker Group/Channel Code"},
	{Text: "link", Description: "This command help you to link two groups/channels," +
		" \nall messages created between this two groups/channels will be share in both two groups/channels" +
		" \nPlease provide to Linker Code of those groups/channels." +
		" \nExample: link yg64y iuh78" +
		" \nPlease use this command only in private message with the bot" +
		" \nDon't use in a group/channel because anyone can view your Linker Group/Channel Code"},
	{Text: "diffuse", Description: "This command help diffuse/broadcast one Group/Channel to another Group/Channel," +
		" \nonly the message of the first group/channel will be share on the second group/channel." +
		" \nNote that with this command the oder of the groups/channels is important" +
		" \nPlease provide to Linker Code of those groups/channels." +
		" \nExample: diffuse yg64y iuh78" +
		" \nThe message of the group/channel yg64y will be share on iuh78. The reverse is not possible." +
		" \nPlease use this command only in private message with the bot" +
		" \nDon't use in a group because anyone can view your Linker Groups Code"},
	{Text: "list", Description: "List all active links and diffusion of the group/channel"},
}

//GetCommands get the list of available commands
func GetCommands() []Command {
	return commands
}
