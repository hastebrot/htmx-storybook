package shadcn_alert

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type AlertProps struct {
	Title       string
	Description string
	IconSlot    g.Node
	Variant     AlertVariant
}

type AlertVariant int

const (
	AlertVariantDefault AlertVariant = iota
	AlertVariantDestructive
)

func Alert(props AlertProps) g.Node {
	return Div(Classes(
		Class("relative w-full rounded-lg p-4 border"),
		g.If(props.Variant == AlertVariantDefault,
			Class("border-[#E4E4E7] bg-[#FFFFFF] text-[#09090B]"),
		),
		g.If(props.Variant == AlertVariantDestructive,
			Class("border-[#DD524C]/50 bg-[#FFFFFF] text-[#DD524C]"),
		),
	),
		Role("alert"),
		Div(Class("absolute left-4 top-4 text-current"),
			Span(Class("[stroke-width:2px]"),
				props.IconSlot,
			),
		),
		H5(Class("pl-7 mb-1 font-medium leading-none _tracking-tight"),
			g.Text(props.Title),
		),
		P(Class("pl-7 text-sm"),
			g.Text(props.Description),
		),
	)
}

func StoryAlert() g.Node {
	return Div(
		story(
			Alert(AlertProps{
				Title:       "Title",
				Description: "Description",
				IconSlot:    lucide.Terminal(Class("block h-4 w-4 flex-shrink-0")),
			}),
		),
		story(
			Alert(AlertProps{
				Title:       "Title",
				Description: "Description",
				IconSlot:    lucide.AlertCircle(Class("block h-4 w-4 flex-shrink-0")),
				Variant:     AlertVariantDestructive,
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
