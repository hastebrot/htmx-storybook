package model

type MenuRoot struct {
	MenuItems []MenuItem
}

type MenuItem struct {
	Text      string
	Type      MenuItemType
	MenuItems []MenuItem
}

type MenuItemType int

const (
	TypeCategory MenuItemType = iota
	TypeFolder
	TypeComponent
	TypeDocument
	TypeStory
)
