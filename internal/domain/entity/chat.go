package entity

import "errors"

type ChatConfig struct {
	Model            *Model
	Temperature      float32 // 0.0 - 1.0
	TopP             float32 // 0.0 - 1.0 - to a low value of 0.1 or lower will give more creative results, higher will give more conservative results.
	N                int     // number of tokens to generate
	Stop             string  // list of tokens to stop on
	MaxTokens        int     // number of tokens to generate
	PresencePenalty  float32 // [optional] -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalizea new tokens based on whether they appear in the text so far. Negative values encourage the model to talk about new topics.
	FrequencyPenalty float32 // [optional] -2.0 to 2.0 - Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far. Negative values encourage the model to talk about less frequent topics.
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
		return errors.New("invalid user id")
	}
	if c.Status != "active" && c.Status != "ended" {
		return errors.New("invalid status")
	}
	if c.Config.Temperature < 0 || c.Config.Temperature > 2 {
		return errors.New("invalid temperature")
	}
	// more validations
	return nil
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is ended. no more messages allowed")
	}
	for {
		if c.Config.Model.GetMaxTokens() >= m.Model.GetMaxTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, m)
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.RefreshTokenUsage()
	}
	return nil
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) CountMessages() int {
	return len(c.Messages)
}

func (c *Chat) End() {
	c.Status = "ended"
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].GetQtdTokens()
	}
}
