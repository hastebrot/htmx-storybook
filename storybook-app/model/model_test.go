package model

import (
	"testing"
)

func TestMenuRoot(t *testing.T) {
	// when:
	menu := MenuRoot{
		MenuItems: []MenuItem{
			{
				Text: "Document",
				Type: TypeDocument,
			},
			{
				Text: "Category",
				Type: TypeCategory,
				MenuItems: []MenuItem{
					{
						Text: "Component",
						Type: TypeComponent,
						MenuItems: []MenuItem{
							{
								Text: "Story",
								Type: TypeStory,
							},
						},
					},
				},
			},
		},
	}

	// then:
	if len(menu.MenuItems) != 2 {
		t.Fatal("len is not 2, menu.MenuItems")
	}
	if len(menu.MenuItems[0].MenuItems) != 0 {
		t.Fatal("len is not 0, menu.MenuItems[0].MenuItems")
	}
	if len(menu.MenuItems[1].MenuItems) != 1 {
		t.Fatal("len is not 1, menu.MenuItems[1].MenuItems")
	}
	if len(menu.MenuItems[1].MenuItems[0].MenuItems) != 1 {
		t.Fatal("len is not 1, menu.MenuItems[1].MenuItems[0].MenuItems")
	}
	if len(menu.MenuItems[1].MenuItems[0].MenuItems[0].MenuItems) != 0 {
		t.Fatal("len is not 0, menu.MenuItems[1].MenuItems[0].MenuItems[0].MenuItems")
	}
}
