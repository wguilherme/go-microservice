package entity

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
	ID                   int
	UserID               int
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}
