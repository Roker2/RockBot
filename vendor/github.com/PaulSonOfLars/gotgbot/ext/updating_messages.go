package ext

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/parsemode"
)

func (b Bot) EditMessageText(chatId int, messageId int, text string) (*Message, error) {
	return b.EditMessageTextMarkup(chatId, messageId, text, "", nil)
}

func (b Bot) EditMessageHTML(chatId int, messageId int, text string) (*Message, error) {
	return b.EditMessageTextMarkup(chatId, messageId, text, parsemode.Html, nil)
}

func (b Bot) EditMessageMarkdown(chatId int, messageId int, text string) (*Message, error) {
	return b.EditMessageTextMarkup(chatId, messageId, text, parsemode.Markdown, nil)
}

func (b Bot) EditMessageMarkdownV2(chatId int, messageId int, text string) (*Message, error) {
	return b.EditMessageTextMarkup(chatId, messageId, text, parsemode.MarkdownV2, nil)
}

func (b Bot) EditMessage(chatId int, messageId int, text string, parseMode string) (*Message, error) {
	return b.EditMessageTextMarkup(chatId, messageId, text, parseMode, nil)
}

func (b Bot) EditMessageTextMarkup(chatId int, messageId int, text string, parseMode string, markup ReplyMarkup) (*Message, error) {
	msg := b.NewSendableEditMessageText(chatId, messageId, text)
	msg.ParseMode = parseMode
	msg.ReplyMarkup = markup
	return msg.Send()
}

func (b Bot) EditMessageTextInline(inlineMessageId string, text string) (*Message, error) {
	return b.EditMessageInline(inlineMessageId, text, "")
}

func (b Bot) EditMessageHTMLInline(inlineMessageId string, text string) (*Message, error) {
	return b.EditMessageInline(inlineMessageId, text, parsemode.Html)
}

func (b Bot) EditMessageMarkdownInline(inlineMessageId string, text string) (*Message, error) {
	return b.EditMessageInline(inlineMessageId, text, parsemode.Markdown)
}

func (b Bot) EditMessageMarkdownV2Inline(inlineMessageId string, text string) (*Message, error) {
	return b.EditMessageInline(inlineMessageId, text, parsemode.MarkdownV2)
}

func (b Bot) EditMessageInline(inlineMessageId string, text string, parseMode string) (*Message, error) {
	msg := b.NewSendableEditMessageText(0, 0, text)
	msg.InlineMessageId = inlineMessageId
	msg.ParseMode = parseMode
	return msg.Send()
}

func (b Bot) EditMessageCaption(chatId int, messageId int, caption string) (*Message, error) {
	return b.editMessageCaption(chatId, messageId, caption, nil, "")
}

func (b Bot) EditMessageCaptionMarkup(chatId int, messageId int, caption string, markup ReplyMarkup) (*Message, error) {
	return b.editMessageCaption(chatId, messageId, caption, markup, "")
}

func (b Bot) EditMessageCaptionParseMode(chatId int, messageId int, caption string, parseMode string) (*Message, error) {
	return b.editMessageCaption(chatId, messageId, caption, nil, parseMode)
}

func (b Bot) editMessageCaption(chatId int, messageId int, caption string, markup ReplyMarkup, parseMode string) (*Message, error) {
	msg := b.NewSendableEditMessageCaption(chatId, messageId, caption)
	msg.ReplyMarkup = markup
	msg.ParseMode = parseMode
	return msg.Send()
}

func (b Bot) EditMessageCaptionInline(inlineMessageId string, caption string) (*Message, error) {
	msg := b.NewSendableEditMessageCaption(0, 0, caption)
	msg.InlineMessageId = inlineMessageId
	return msg.Send()
}

func (b Bot) EditMessageReplyMarkup(chatId int, messageId int, replyMarkup InlineKeyboardMarkup) (*Message, error) {
	msg := b.NewSendableEditMessageReplyMarkup(chatId, messageId, &replyMarkup)
	return msg.Send()
}

func (b Bot) EditMessageReplyMarkupInline(inlineMessageId string, replyMarkup InlineKeyboardMarkup) (*Message, error) {
	msg := b.NewSendableEditMessageReplyMarkup(0, 0, &replyMarkup)
	msg.InlineMessageId = inlineMessageId
	return msg.Send()
}

func (b Bot) DeleteMessage(chatId int, messageId int) (bool, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(chatId))
	v.Add("message_id", strconv.Itoa(messageId))

	return b.boolSender("deleteMessage", v)
}

func (b Bot) boolSender(meth string, v url.Values) (bb bool, err error) {
	r, err := b.Get(meth, v)
	if err != nil {
		return false, err
	}

	return bb, json.Unmarshal(r, &bb)
}
