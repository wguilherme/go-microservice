package entity

// struct
type ChatConfig struct {
	Model *Model
}

type Chat struct {
	ID                   int
	UserID               int
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}
