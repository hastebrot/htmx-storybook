package shadcn_select

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type SelectProps struct {
	Placeholder string
}

func Select(props SelectProps, children ...g.Node) g.Node {
	return h.Div(h.Class("shadcn-select flex flex-col _w-full w-[180px]"),
		h.Button(Classes(
			h.Class("flex h-10 items-center justify-between rounded-md"),
			h.Class("border border-[#E4E4E7] bg-[#FFFFFF] px-3 py-2 text-sm"),
			h.Class("placeholder:text-[#71717A] disabled:cursor-not-allowed disabled:opacity-50"),
		),
			h.Span(h.Class("pointer-events-none"), g.Text(props.Placeholder)),
			h.Span(h.Class("opacity-50"),
				lucide.ChevronDown(h.Class("block h-4 w-4 flex-shrink-0")),
			),
		),
		SelectPopover(
			g.Group(children),
		),
	)
}

func SelectPopover(children ...g.Node) g.Node {
	return h.Div(Classes(
		h.Class("relative z-50 max-h-96 min-w-[8rem] overflow-hidden rounded-md"),
		h.Class("border bg-[#FFFFFF] text-[#09090B] shadow-md mt-[4px]"),
	),
		h.Div(h.Class("p-1 w-full"),
			g.Group(children),
		),
	)
}

type SelectItemProps struct {
	Text       string
	IsSelected bool
}

func SelectItem(props SelectItemProps) g.Node {
	return h.Div(
		h.Role("option"),
		Classes(
			h.Class("relative flex w-full cursor-default select-none items-center"),
			h.Class("rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none"),
			h.Class("hover:bg-[#F4F4F5] focus:bg-[#F4F4F5] focus:text-[#18181B]"),
			h.Class("disabled:pointer-events-none disabled:opacity-50"),
		),
		h.Span(h.Class("absolute left-2 flex h-3.5 w-3.5 items-center justify-center"),
			g.If(props.IsSelected,
				lucide.Check(h.Class("block h-4 w-4 flex-shrink-0 [stroke-width:2px]")),
			),
		),
		h.Span(g.Text(props.Text)),
	)
}

func StorySelect() g.Node {
	return story(
		Select(SelectProps{
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

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[350px] w-full items-start justify-center p-10"),
		g.Group(children),
	)
}
