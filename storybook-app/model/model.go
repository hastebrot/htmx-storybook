package model

// some constraints/requirements are
// TypeSection and TypeComponent contain subitems and need to be collapsible in the UI
// each MenuItemType determines how the menu item is rendered in the UI
// the root of the menu is not displayed in the UI
// each MenuItem maps to a path of the page router, i.e. a URL (not considered in the data structure, yet)
// when MenuItems of a MenuItem are collapsed or expanded in the UI, there will be no additional request to the server
// MenuRoot with its MenuItems are transformed into htmx-flavored HTML
// there is a search input that allows to filter MenuItems by title

// so I guess, there is no need to model these constraints in the data structure, it's okay to follow the constraints in the logic that transforms the structs into HTML.

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
