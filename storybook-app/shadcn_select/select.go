package shadcn_select

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type SelectButtonProps struct {
	Placeholder string
}

func SelectButton(props SelectButtonProps, children ...g.Node) g.Node {
	return Div(Class("shadcn-select flex flex-col _w-full w-[180px]"),
		Button(Classes(
			Class("flex h-10 items-center justify-between rounded-md"),
			Class("border border-[#E4E4E7] bg-[#FFFFFF] px-3 py-2 text-sm"),
			Class("placeholder:text-[#71717A] disabled:cursor-not-allowed disabled:opacity-50"),
		),
			Span(Class("pointer-events-none"), g.Text(props.Placeholder)),
			Span(Class("opacity-50"),
				lucide.ChevronDown(Class("block h-4 w-4 flex-shrink-0")),
			),
		),
		SelectPopover(
			g.Group(children),
		),
	)
}

func SelectPopover(children ...g.Node) g.Node {
	return Div(Classes(
		Class("relative z-50 max-h-96 min-w-[8rem] overflow-hidden rounded-md"),
		Class("border bg-[#FFFFFF] text-[#09090B] shadow-md mt-[4px]"),
	),
		Div(Class("p-1 w-full"),
			g.Group(children),
		),
	)
}

func SelectStory() g.Node {
	return Story(
		SelectButton(SelectButtonProps{
			Placeholder: "Select a fruit",
		},
			SelectItem(SelectItemProps{Text: "Apple"}),
			SelectItem(SelectItemProps{Text: "Banana", IsSelected: true}),
			SelectItem(SelectItemProps{Text: "Blueberry"}),
			SelectItem(SelectItemProps{Text: "Grapes"}),
			SelectItem(SelectItemProps{Text: "Pineapple"}),
		),
	)
}

type SelectItemProps struct {
	Text       string
	IsSelected bool
}

func SelectItem(props SelectItemProps) g.Node {
	return Div(
		Role("option"),
		Classes(
			Class("relative flex w-full cursor-default select-none items-center"),
			Class("rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none"),
			Class("hover:bg-[#F4F4F5] focus:bg-[#F4F4F5] focus:text-[#18181B]"),
			Class("disabled:pointer-events-none disabled:opacity-50"),
		),
		Span(Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
			g.If(props.IsSelected,
				lucide.Check(Class("block h-4 w-4 flex-shrink-0 [stroke-width:2px]")),
			),
		),
		Span(g.Text(props.Text)),
	)
}

func Story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[350px] w-full items-start justify-center p-10"),
		g.Group(children),
	)
}
