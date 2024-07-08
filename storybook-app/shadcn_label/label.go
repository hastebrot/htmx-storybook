package shadcn_label

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

type LabelProps struct {
	Text string
	For  string
}

func Label(props LabelProps) g.Node {
	return h.Label(h.Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),
		g.If(props.For != "", h.For(props.For)),
		g.Text(props.Text),
	)
}

func StoryLabel() g.Node {
	return h.Div(
		story(
			Label(LabelProps{
				Text: "Label",
			}),
		),
	)
}

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[150px] w-full items-start justify-center p-10 gap-10"),
		g.Group(children),
	)
}
