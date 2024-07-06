package model

import (
	g "github.com/maragudk/gomponents"
)

type MenuRoot struct {
	MenuItems []MenuItem
}

type MenuItem struct {
	Text      string
	Link      string
	Type      MenuItemType
	MenuItems []MenuItem
	Component func() g.Node
}

type MenuItemType int

const (
	TypeCategory MenuItemType = iota
	TypeFolder
	TypeComponent
	TypeDocument
	TypeStory
)
