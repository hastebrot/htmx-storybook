package shadcn_radiogroup

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	"storybook-app/shadcn_label"
)

func RadioGroup(children ...g.Node) g.Node {
	return Div(Class("grid gap-2"),
		Role("radiogroup"),
		g.Group(children),
	)
}

type RadioButtonProps struct {
	Label      string
	IsChecked  bool
	IsDisabled bool
}

func RadioButton(props RadioButtonProps) g.Node {
	return Div(Class("flex items-center space-x-2"),
		Button(
			Role("radio"),
			Classes(
				Class("peer aspect-square h-4 w-4 shrink-0 rounded-full border border-[#18181B]"),
				Class("focus-visible:outline-none"),
				Class("disabled:cursor-not-allowed disabled:opacity-50"),
				g.If(props.IsChecked, Class("bg-[#FAFAFA] text-[#18181B]")),
			),
			g.If(props.IsDisabled, Disabled()),
			g.If(props.IsChecked,
				Span(Class("flex items-center justify-center"),
					lucide.Circle(Class("block h-2.5 w-2.5 flex-shrink-0 [stroke-width:2px] fill-current")),
				),
			),
		),
		shadcn_label.Label(shadcn_label.LabelProps{
			Text: props.Label,
		}),
	)
}

func StoryRadioGroup() g.Node {
	return Div(
		story(
			RadioGroup(
				RadioButton(RadioButtonProps{
					Label: "Default",
				}),
				RadioButton(RadioButtonProps{
					Label:     "Comfortable",
					IsChecked: true,
				}),
				RadioButton(RadioButtonProps{
					Label:      "Compact",
					IsDisabled: true,
				}),
			),
		),
	)
}

func story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[150px] w-full items-start justify-center p-10 gap-10"),
		g.Group(children),
	)
}
