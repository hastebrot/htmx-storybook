package router

import (
	g "github.com/maragudk/gomponents"

	"storybook-app/model"
	"storybook-app/shadcn_accordion"
	"storybook-app/shadcn_select"
	"storybook-app/shadcn_tabs"
	"storybook-app/story_page"
)

func BuildRootNode(page string) g.Node {
	menu := buildMenu()
	var node g.Node
	switch page {
	case "shadcn_accordion":
		node = story_page.StoryPage(story_page.StoryPageProps{
			CanvasSlot: shadcn_accordion.AccordionStory(),
			MenuRoot:   menu,
		})
	case "shadcn_select":
		node = story_page.StoryPage(story_page.StoryPageProps{
			CanvasSlot: shadcn_select.SelectStory(),
			MenuRoot:   menu,
		})
	case "shadcn_tabs":
		node = story_page.StoryPage(story_page.StoryPageProps{
			CanvasSlot: shadcn_tabs.TabsStory(),
			MenuRoot:   menu,
		})
	default:
		node = story_page.StoryPage(story_page.StoryPageProps{
			CanvasSlot: nil,
			MenuRoot:   menu,
		})
	}
	return node
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
						Text: "shadcn_select",
						Link: "/pages/shadcn_select",
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
