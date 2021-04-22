package linker

type Command struct {
	Text, Option, Description string
}

var commands = []Command{
	{Text: "help", Description: "Display bot help."},
	{Text: "config", Description: "Get actual group/channel Linker Code" +
		" \nOnly admin can use this command"},
	{Text: "list", Description: "List all links and diffusion of the group/channel" +
		" \nOnly admin can use this command"},
	{Text: "link", Option: "<group/channel_id> <group/channel_id>", Description: "Link two groups/channels," +
		" \nPlease use this command only in private message with the bot"},
	{Text: "diffuse", Option: "<group/channel_id> <group/channel_id>", Description: "Diffuse/broadcast the first Group/Channel in the second Group/Channel," +
		" \nPlease use this command only in private message with the bot"},
	{Text: "start", Option: "<link_or_diffusion_id>", Description: "Start an existing stopped link or diffusion" +
		" \nPlease use this command only in private message with the bot"},
	{Text: "stop", Option: "<link_or_diffusion_id>", Description: "Stop an existing active link or diffusion" +
		" \nPlease use this command only in private message with the bot"},
}

//GetCommands get the list of available commands
func GetCommands() []Command {
	return commands
}
