package shadcn_button

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type ButtonProps struct {
	Text       string
	Variant    ButtonVariant
	IsDisabled bool
}

type ButtonVariant int

const (
	ButtonVariantPrimary ButtonVariant = iota
	ButtonVariantSecondary
	ButtonVariantDestructive
	ButtonVariantOutline
	ButtonVariantGhost
	ButtonVariantLink
)

func Button(props ButtonProps) g.Node {
	return h.Button(Classes(
		h.Class("inline-flex items-center justify-center whitespace-nowrap rounded-md"),
		h.Class("text-sm font-medium transition-colors"),
		h.Class("h-10 px-4 py-2"),
		h.Class("focus-visible:outline-none"),
		h.Class("disabled:pointer-events-none disabled:opacity-50"),
		g.If(props.Variant == ButtonVariantPrimary,
			h.Class("bg-[#18181B] text-[#FAFAFA] hover:bg-[#18181B]/90"),
		),
		g.If(props.Variant == ButtonVariantSecondary,
			h.Class("bg-[#F4F4F5] text-[#18181B] hover:bg-[#F4F4F5]/80"),
		),
		g.If(props.Variant == ButtonVariantDestructive,
			h.Class("bg-[#DD524C] text-[#FAFAFA] hover:bg-[#DD524C]/90"),
		),
		g.If(props.Variant == ButtonVariantOutline,
			h.Class("border border-[#E4E4E7] bg-[#FFFFFF] text-[#09090B] hover:bg-[#F4F4F5] hover:text-[#18181B]"),
		),
		g.If(props.Variant == ButtonVariantGhost,
			h.Class("text-[#09090B] hover:text-[#18181B] hover:bg-[#F4F4F5]"),
		),
		g.If(props.Variant == ButtonVariantLink,
			h.Class("text-[#09090B] hover:text-[#18181B] hover:underline"),
		),
	),
		g.If(props.IsDisabled, h.Disabled()),
		g.Text(props.Text),
	)
}

func StoryButton() g.Node {
	return h.Div(
		story(
			Button(ButtonProps{
				Text: "Button",
			}),
			Button(ButtonProps{
				Text:    "Button",
				Variant: ButtonVariantSecondary,
			}),
			Button(ButtonProps{
				Text:    "Button",
				Variant: ButtonVariantDestructive,
			}),
			Button(ButtonProps{
				Text:    "Button",
				Variant: ButtonVariantOutline,
			}),
			Button(ButtonProps{
				Text:    "Button",
				Variant: ButtonVariantGhost,
			}),
			Button(ButtonProps{
				Text:    "Button",
				Variant: ButtonVariantLink,
			}),
		),
		story(
			Button(ButtonProps{
				Text:       "Button",
				IsDisabled: true,
			}),
			Button(ButtonProps{
				Text:       "Button",
				Variant:    ButtonVariantSecondary,
				IsDisabled: true,
			}),
			Button(ButtonProps{
				Text:       "Button",
				Variant:    ButtonVariantDestructive,
				IsDisabled: true,
			}),
			Button(ButtonProps{
				Text:       "Button",
				Variant:    ButtonVariantOutline,
				IsDisabled: true,
			}),
			Button(ButtonProps{
				Text:       "Button",
				Variant:    ButtonVariantGhost,
				IsDisabled: true,
			}),
			Button(ButtonProps{
				Text:       "Button",
				Variant:    ButtonVariantLink,
				IsDisabled: true,
			}),
		),
	)
}

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[350px] w-full items-start justify-center p-10 gap-10 flex-wrap"),
		g.Group(children),
	)
}
