package shadcn_input

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	"storybook-app/shadcn_label"
)

type InputProps struct {
	Type        string
	Placeholder string
	Value       string
	IsDisabled  bool
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
		g.If(props.IsDisabled, h.Disabled()),
		g.If(props.Value != "", h.Value(props.Value)),
		h.Type(props.Type),
		h.Placeholder(props.Placeholder),
	)
}

func StoryInput() g.Node {
	return h.Div(
		story(
			h.Div(h.Class("grid w-full max-w-sm items-center gap-1.5"),
				shadcn_label.Label(shadcn_label.LabelProps{
					Text: "Username",
				}),
				Input(InputProps{
					Type:        "text",
					Placeholder: "Username",
				}),
			),
		),
		story(
			h.Div(h.Class("grid w-full max-w-sm items-center gap-1.5"),
				shadcn_label.Label(shadcn_label.LabelProps{
					Text: "Password",
				}),
				Input(InputProps{
					Type:        "password",
					Placeholder: "Password",
					Value:       "password",
				}),
			),
		),
		story(
			h.Div(h.Class("grid w-full max-w-sm items-center gap-1.5"),
				shadcn_label.Label(shadcn_label.LabelProps{
					Text: "Email",
				}),
				Input(InputProps{
					Type:        "email",
					Placeholder: "Email",
				}),
			),
		),
		story(
			h.Div(h.Class("grid w-full max-w-sm items-center gap-1.5"),
				Input(InputProps{
					Type:        "text",
					Placeholder: "Disabled",
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
