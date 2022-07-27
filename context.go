package dgc

import (
	"github.com/bwmarrin/discordgo"
)

// Ctx represents the context for a command event
type Ctx struct {
	Session       *discordgo.Session
	Event         *discordgo.MessageCreate
	Arguments     *Arguments
	CustomObjects *ObjectsMap
	Router        *Router
	Command       *Command
	Message       *discordgo.Message
	Locals        *map[string]interface{}
}

// ExecutionHandler represents a handler for a context execution
type ExecutionHandler func(*Ctx)

// RespondText responds with the given text message
func (ctx *Ctx) RespondText(text string) error {
	m, err := ctx.Session.ChannelMessageSend(ctx.Event.ChannelID, text)
	ctx.Message = m
	return err
}

// RespondEmbed responds with the given embed message
func (ctx *Ctx) RespondEmbed(embed *discordgo.MessageEmbed) error {
	m, err := ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embed)
	ctx.Message = m
	return err
}

// RespondTextEmbed responds with the given text and embed message
func (ctx *Ctx) RespondTextEmbed(text string, embed *discordgo.MessageEmbed) error {
	m, err := ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Content: text,
		Embed:   embed,
	})
	ctx.Message = m
	return err
}

// ReplyText replies to the message with the given text message
func (ctx *Ctx) ReplyText(text string) error {
	m, err := ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Content:   text,
		Reference: ctx.Message.Reference(),
	})
	ctx.Message = m
	return err
}

// ReplyEmbed replies to the message with the given embed message
func (ctx *Ctx) ReplyEmbed(embed *discordgo.MessageEmbed) error {
	m, err := ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Embed:     embed,
		Reference: ctx.Message.Reference(),
		Content:   "",
	})
	ctx.Message = m
	return err
}

// ReplyTextEmbed replies to the message with the given text and embed message
func (ctx *Ctx) ReplyTextEmbed(text string, embed *discordgo.MessageEmbed) error {
	m, err := ctx.Session.ChannelMessageSendComplex(ctx.Event.ChannelID, &discordgo.MessageSend{
		Content:   text,
		Embed:     embed,
		Reference: ctx.Message.Reference(),
	})
	ctx.Message = m
	return err
}
