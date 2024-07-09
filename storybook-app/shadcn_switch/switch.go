package shadcn_switch

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type SwitchProps struct {
	Text       string
	IsChecked  bool
	IsDisabled bool
}

func Switch(props SwitchProps) g.Node {
	return Div(Class("flex items-center space-x-2"),
		Label(
			Role("switch"),
			Classes(
				Class("peer inline-flex h-6 w-11 shrink-0 cursor-pointer items-center rounded-full"),
				Class("border-2 border-transparent transition-colors"),
				Class("disabled:cursor-not-allowed disabled:opacity-50"),
				g.If(props.IsDisabled,
					Class("!cursor-not-allowed opacity-50"),
				),
				Class("bg-[#E4E4E7] [&:has(input:checked)]:bg-[#18181B]"),
			),
			g.If(props.IsDisabled, Disabled()),
			Input(Class("hidden peer"),
				Type("checkbox"),
				g.If(props.IsDisabled, Disabled()),
				g.If(props.IsChecked, Checked()),
			),
			Span(Classes(
				Class("pointer-events-none block h-5 w-5 rounded-full"),
				Class("bg-[#FFFFFF] shadow-lg ring-0"),
				Class("transition-transform"),
				Class("translate-x-0 peer-checked:translate-x-5"),
			)),
		),
		Label(Class("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),
			g.Text(props.Text),
		),
	)
}

func StorySwitch() g.Node {
	return Div(
		story(
			Switch(SwitchProps{
				Text:      "Label",
				IsChecked: false,
			}),
			Switch(SwitchProps{
				Text:      "Label",
				IsChecked: true,
			}),
		),
		story(
			Switch(SwitchProps{
				Text:       "Label",
				IsChecked:  false,
				IsDisabled: true,
			}),
			Switch(SwitchProps{
				Text:       "Label",
				IsChecked:  true,
				IsDisabled: true,
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
