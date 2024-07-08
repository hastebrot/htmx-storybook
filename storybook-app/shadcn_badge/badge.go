package shadcn_badge

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type BadgeProps struct {
	Text    string
	Variant BadgeVariant
}

type BadgeVariant int

const (
	BadgeVariantPrimary BadgeVariant = iota
	BadgeVariantSecondary
	BadgeVariantDestructive
	BadgeVariantOutline
)

func Badge(props BadgeProps) g.Node {
	return h.Div(Classes(
		h.Class("inline-flex items-center rounded-full"),
		h.Class("text-xs font-semibold transition-colors"),
		h.Class("px-2.5 py-0.5"),
		h.Class("focus-visible:outline-none"),
		g.If(props.Variant == BadgeVariantPrimary,
			h.Class("bg-[#18181B] text-[#FAFAFA]"),
		),
		g.If(props.Variant == BadgeVariantSecondary,
			h.Class("bg-[#F4F4F5] text-[#18181B]"),
		),
		g.If(props.Variant == BadgeVariantDestructive,
			h.Class("bg-[#DD524C] text-[#FAFAFA]"),
		),
		g.If(props.Variant == BadgeVariantOutline,
			h.Class("border border-[#E4E4E7] bg-[#FFFFFF] text-[#09090B]"),
		),
	),
		g.Text(props.Text),
	)
}

func StoryBadge() g.Node {
	return story(
		Badge(BadgeProps{
			Text: "Badge",
		}),
		Badge(BadgeProps{
			Text:    "Badge",
			Variant: BadgeVariantSecondary,
		}),
		Badge(BadgeProps{
			Text:    "Badge",
			Variant: BadgeVariantDestructive,
		}),
		Badge(BadgeProps{
			Text:    "Badge",
			Variant: BadgeVariantOutline,
		}),
	)
}

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[350px] w-full items-start justify-center p-10 gap-10 flex-wrap"),
		g.Group(children),
	)
}
