package shadcn_dialog

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func Dialog(children ...g.Node) g.Node {
	return g.Group(children)
}

func DialogContent(children ...g.Node) g.Node {
	return h.Div(Classes(
		// h.Class("fixed left-[50%] top-[50%] translate-x-[-50%] translate-y-[-50%]"),
		h.Class("relative z-50 grid w-full max-w-lg gap-4 p-6"),
		h.Class("border border-[#E5E7EB] bg-[#FFFFFF] shadow-lg"),
		h.Class("sm:rounded-lg sm:max-w-md"),
	),
		g.Group(children),
		h.Button(Classes(
			h.Class("absolute right-4 top-4 rounded-sm opacity-70"),
			h.Class("hover:opacity-100 transition-opacity"),
			h.Class("focus:outline-none"),
			h.Class("disabled:pointer-events-none"),
		),
			h.Span(h.Class("[stroke-width:2px]"),
				lucide.X(h.Class("block h-4 w-4 flex-shrink-0")),
			),
		),
	)
}

type DialogHeaderProps struct {
	Title       string
	Description string
}

func DialogHeader(props DialogHeaderProps) g.Node {
	return h.Div(h.Class("flex flex-col space-y-1.5 text-center sm:text-left"),
		h.H2(h.Class("text-lg font-semibold leading-none _tracking-tight"),
			g.Text(props.Title),
		),
		h.P(h.Class("text-sm text-[#717179]"),
			g.Text(props.Description),
		),
	)
}

func DialogFooter(children ...g.Node) g.Node {
	return h.Div(h.Class("flex flex-col-reverse sm:flex-row sm:space-x-2 sm:justify-start"),
		g.Group(children),
	)
}

func StoryDialog() g.Node {
	return h.Div(
		story(
			Dialog(
				DialogContent(
					DialogHeader(DialogHeaderProps{
						Title:       "Title",
						Description: "Description",
					}),
					h.Div(
						g.Text("Dialog content"),
					),
					DialogFooter(
						g.Text("Dialog footer"),
					),
				),
			),
		),
	)
}

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[150px] w-full items-start justify-center p-10 gap-10"),
		g.Group(children),
	)
}
