package momentummodel

type IModel interface {
	Id() string
	SetId(string)
}

type Model struct {
	IModel

	id string
}

func (m *Model) Id() string {
	return m.id
}

func (m *Model) SetId(id string) {
	m.id = id
}
