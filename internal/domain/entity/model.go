package entity

// struct
type Model struct {
	Name      string
	MaxTokens int
}

// construtctor func
func NewModel(name string, maxTokens int) *Model {
	return &Model{
		Name:      name,
		MaxTokens: maxTokens,
	}
}

// method to get the model
func (m *Model) GetMaxTokens() int {
	return m.MaxTokens
}

func (m *Model) GetName() string {
	return m.Name
}
