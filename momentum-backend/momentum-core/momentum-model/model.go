package momentummodel

type IModel interface {
	Id() string
	SetId(string)
}

type Model struct {
	IModel

	id string
}
