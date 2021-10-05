package handler

import (
	"github.com/cruffinoni/Digona/src/commands"
	"github.com/cruffinoni/Digona/src/commands/parser"
	"github.com/cruffinoni/Digona/src/commands/privateMessage"
	"github.com/cruffinoni/Digona/src/commands/reaction"
	"github.com/cruffinoni/Digona/src/commands/role"
	"strings"
)

type CommandHandler func(*parser.MessageParser) error

type CommandPair struct {
	Name    string
	Command CommandHandler
}

var commandsListing = map[string]CommandHandler{
	"delete":       commands.RedirectDelete,
	"react":        reaction.Role,
	"qr-code":      privateMessage.GenerateQrCode,
	"default-role": role.SetDefaultRole,
}

func GetCommandFromArgs(args []string) *CommandPair {
	for _, word := range args {
		if function, exists := commandsListing[strings.ToLower(word)]; exists {
			return &CommandPair{
				Name:    word,
				Command: function,
			}
		}
	}
	return nil
}
