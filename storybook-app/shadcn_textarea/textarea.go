package shadcn_textarea

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	"storybook-app/shadcn_label"
)

type TextareaProps struct {
	Placeholder string
	Value       string
	IsDisabled  bool
}

func Textarea(props TextareaProps) g.Node {
	return h.Textarea(
		Classes(
			h.Class("flex min-h-[80px] w-full rounded-md "),
			h.Class("border border-[#E4E4E7] bg-[#FFFFFF]"),
			h.Class("px-3 py-2 text-sm"),
			h.Class("file:border-0 file:bg-transparent file:text-sm file:font-medium"),
			h.Class("placeholder:text-[#717179] focus-visible:outline-none"),
			h.Class("disabled:cursor-not-allowed disabled:opacity-50"),
		),
		g.If(props.IsDisabled, h.Disabled()),
		g.If(props.Value != "", h.Value(props.Value)),
		h.Placeholder(props.Placeholder),
	)
}

func StoryTextarea() g.Node {
	return h.Div(
		story(
			h.Div(h.Class("grid w-full gap-1.5"),
				shadcn_label.Label(shadcn_label.LabelProps{
					Text: "Your message",
				}),
				Textarea(TextareaProps{
					Placeholder: "Type your message here.",
				}),
			),
		),
		story(
			h.Div(h.Class("grid w-full gap-1.5"),
				shadcn_label.Label(shadcn_label.LabelProps{
					Text: "Your message",
				}),
				Textarea(TextareaProps{
					Placeholder: "Type your message here.",
					IsDisabled:  true,
				}),
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
