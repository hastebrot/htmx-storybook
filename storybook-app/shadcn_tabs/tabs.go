package shadcn_tabs

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func Tabs(children ...g.Node) g.Node {
	tabsClass := `
		[&:has(.shadcn-tabs-item:nth-of-type(1)_input:checked)_.shadcn-tabs-content:nth-of-type(1)]:block
		[&:has(.shadcn-tabs-item:nth-of-type(2)_input:checked)_.shadcn-tabs-content:nth-of-type(2)]:block
		[&:has(.shadcn-tabs-item:nth-of-type(3)_input:checked)_.shadcn-tabs-content:nth-of-type(3)]:block
		[&:has(.shadcn-tabs-item:nth-of-type(4)_input:checked)_.shadcn-tabs-content:nth-of-type(4)]:block
		[&:has(.shadcn-tabs-item:nth-of-type(5)_input:checked)_.shadcn-tabs-content:nth-of-type(5)]:block
	`

	return Div(Class(fmt.Sprintf("shadcn-tabs flex flex-col w-full %s", tabsClass)),
		g.Group(children),
	)
}

func TabsList(children ...g.Node) g.Node {
	tabsColumns := len(children)

	return Div(Classes(
		Class("shadcn-tabs-list grid grid-cols-[repeat(var(--tabs-columns),1fr)] items-center justify-center"),
		Class("w-full h-10 rounded-md p-1 text-[#717179] bg-[#F4F4F5]"),
	),
		Style(fmt.Sprintf("--tabs-columns:%x", tabsColumns)),
		g.Group(children),
	)
}

type TabsItemProps struct {
	Name     string
	Text     string
	IsActive bool
}

func TabsItem(props TabsItemProps) g.Node {
	return Div(Class("shadcn-tabs-item"),
		Header(Class("flex [&:has(input:checked)]:bg-[#FFFFFF] [&:has(input:checked)]:text-[#09090B] [&:has(input:checked)]:shadow-sm"),
			Label(Classes(
				Class("flex flex-1 items-center justify-center whitespace-nowrap"),
				Class("rounded-sm px-3 py-1.5 text-sm font-medium"),
				Class("disabled:pointer-events-none disabled:opacity-50"),
				Class("cursor-pointer select-none"),
			),
				Input(Class("peer hidden"), Type("radio"), g.If(props.IsActive, g.Attr("checked")), Name(props.Name)),
				g.Text(props.Text),
			),
		),
	)
}

func TabsContent(children ...g.Node) g.Node {
	return Section(Class("shadcn-tabs-content hidden overflow-hidden text-sm font-normal"),
		Div(Class("mt-2"), g.Group(children)),
	)
}

func TabsStory() g.Node {
	return Story(
		Tabs(
			TabsList(
				TabsItem(TabsItemProps{
					Name:     "shadcn_tabs",
					Text:     "Section one",
					IsActive: true,
				}),
				TabsItem(TabsItemProps{
					Name: "shadcn_tabs",
					Text: "Section two",
				}),
			),
			TabsContent(g.Text("Content one")),
			TabsContent(g.Text("Content two")),
		),
	)
}

func Story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[350px] w-full items-start justify-center p-10"),
		g.Group(children),
	)
}
