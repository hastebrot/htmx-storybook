package router

import (
	g "github.com/maragudk/gomponents"

	"storybook-app/model"
	"storybook-app/shadcn_accordion"
	"storybook-app/shadcn_button"
	"storybook-app/shadcn_checkbox"
	"storybook-app/shadcn_select"
	"storybook-app/shadcn_switch"
	"storybook-app/shadcn_tabs"
	"storybook-app/story_page"
)

func BuildRootNode(page string) g.Node {
	menu := buildMenu()
	var node g.Node
	switch page {
	case "shadcn_accordion":
		node = shadcn_accordion.StoryAccordion()
	case "shadcn_button":
		node = shadcn_button.StoryButton()
	case "shadcn_checkbox":
		node = shadcn_checkbox.StoryCheckbox()
	case "shadcn_select":
		node = shadcn_select.StorySelect()
	case "shadcn_switch":
		node = shadcn_switch.StorySwitch()
	case "shadcn_tabs":
		node = shadcn_tabs.StoryTabs()
	default:
		node = nil
	}
	return story_page.StoryPage(story_page.StoryPageProps{
		CanvasSlot: node,
		MenuRoot:   menu,
	})
}

func buildMenu() model.MenuRoot {
	return model.MenuRoot{
		MenuItems: []model.MenuItem{
			{
				Text: "index",
				Link: "/pages/",
				Type: model.TypeStory,
			},
			{
				Text: "shadcn-ui",
				MenuItems: []model.MenuItem{
					{
						Text: "shadcn_accordion",
						Link: "/pages/shadcn_accordion",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_button",
						Link: "/pages/shadcn_button",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_checkbox",
						Link: "/pages/shadcn_checkbox",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_select",
						Link: "/pages/shadcn_select",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_switch",
						Link: "/pages/shadcn_switch",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_tabs",
						Link: "/pages/shadcn_tabs",
						Type: model.TypeStory,
					},
				},
			},
		},
	}
}
