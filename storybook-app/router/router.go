package router

import (
	g "github.com/maragudk/gomponents"

	"storybook-app/model"
	"storybook-app/shadcn_accordion"
	"storybook-app/shadcn_alert"
	"storybook-app/shadcn_badge"
	"storybook-app/shadcn_breadcrumb"
	"storybook-app/shadcn_button"
	"storybook-app/shadcn_card"
	"storybook-app/shadcn_checkbox"
	"storybook-app/shadcn_command"
	"storybook-app/shadcn_input"
	"storybook-app/shadcn_label"
	"storybook-app/shadcn_radiogroup"
	"storybook-app/shadcn_select"
	"storybook-app/shadcn_switch"
	"storybook-app/shadcn_table"
	"storybook-app/shadcn_tabs"
	"storybook-app/story_page"
)

func BuildRootNode(page string) g.Node {
	menu := buildMenu()
	var node g.Node
	switch page {
	case "shadcn_accordion":
		node = shadcn_accordion.StoryAccordion()
	case "shadcn_alert":
		node = shadcn_alert.StoryAlert()
	case "shadcn_badge":
		node = shadcn_badge.StoryBadge()
	case "shadcn_breadcrumb":
		node = shadcn_breadcrumb.StoryBreadcrumb()
	case "shadcn_button":
		node = shadcn_button.StoryButton()
	case "shadcn_card":
		node = shadcn_card.StoryCard()
	case "shadcn_checkbox":
		node = shadcn_checkbox.StoryCheckbox()
	case "shadcn_command":
		node = shadcn_command.StoryCommand()
	case "shadcn_input":
		node = shadcn_input.StoryInput()
	case "shadcn_label":
		node = shadcn_label.StoryLabel()
	case "shadcn_radiogroup":
		node = shadcn_radiogroup.StoryRadioGroup()
	case "shadcn_select":
		node = shadcn_select.StorySelect()
	case "shadcn_switch":
		node = shadcn_switch.StorySwitch()
	case "shadcn_table":
		node = shadcn_table.StoryTable()
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
						Text: "shadcn_alert",
						Link: "/pages/shadcn_alert",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_badge",
						Link: "/pages/shadcn_badge",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_breadcrumb",
						Link: "/pages/shadcn_breadcrumb",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_button",
						Link: "/pages/shadcn_button",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_card",
						Link: "/pages/shadcn_card",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_checkbox",
						Link: "/pages/shadcn_checkbox",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_command",
						Link: "/pages/shadcn_command",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_input",
						Link: "/pages/shadcn_input",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_label",
						Link: "/pages/shadcn_label",
						Type: model.TypeStory,
					},
					{
						Text: "shadcn_radiogroup",
						Link: "/pages/shadcn_radiogroup",
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
						Text: "shadcn_table",
						Link: "/pages/shadcn_table",
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
