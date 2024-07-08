package shadcn_checkbox

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	"storybook-app/shadcn_label"
)

type CheckboxProps struct {
	Label     string
	IsChecked bool
}

func Checkbox(props CheckboxProps) g.Node {
	return Div(Class("flex items-center space-x-2"),
		Button(
			Role("checkbox"),
			Classes(
				Class("peer h-4 w-4 shrink-0 rounded-[4px] border border-[#18181B]"),
				Class("focus-visible:outline-none"),
				Class("disabled:cursor-not-allowed disabled:opacity-50"),
				g.If(props.IsChecked, Class("bg-[#18181B] text-[#FAFAFA]")),
			),
			g.If(props.IsChecked,
				Span(Class("flex items-center justify-center"),
					lucide.Check(Class("block h-3.5 w-3.5 flex-shrink-0 [stroke-width:2px] text-current")),
				),
			),
		),
		shadcn_label.Label(shadcn_label.LabelProps{
			Text: props.Label,
		}),
	)
}

func StoryCheckbox() g.Node {
	return Div(
		story(
			Checkbox(CheckboxProps{
				Label: "Label",
			}),
			Checkbox(CheckboxProps{
				Label:     "Label",
				IsChecked: true,
			}),
		),
	)
}

func story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[150px] w-full items-start justify-center p-10 gap-10"),
		g.Group(children),
	)
}
