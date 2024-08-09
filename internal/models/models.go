package models

type Model interface {
	AddItem(content string) error
	RemoveItem(id int) error
	ToggleItem(id int) error
	EditItem(id int) error
	ListItems() error
}
