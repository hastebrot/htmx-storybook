package shadcn_input

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type InputProps struct {
	Type        string
	Placeholder string
}

func Input(props InputProps) g.Node {
	return h.Input(
		Classes(
			h.Class("flex h-10 w-full rounded-md"),
			h.Class("border border-[#E4E4E7] bg-[#FFFFFF]"),
			h.Class("px-3 py-2 text-sm"),
			h.Class("file:border-0 file:bg-transparent file:text-sm file:font-medium"),
			h.Class("placeholder:text-[#717179] focus-visible:outline-none"),
			h.Class("disabled:cursor-not-allowed disabled:opacity-50"),
		),
		h.Type(props.Type),
		h.Placeholder(props.Placeholder),
	)
}

func StoryInput() g.Node {
	return h.Div(
		story(
			Input(InputProps{
				Type:        "email",
				Placeholder: "Email",
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
