package shadcn_card

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Card(children ...g.Node) g.Node {
	return Div(Class("rounded-lg border bg-[#FEFEFE] text-[#09090B] shadow-sm w-[350px]"),
		g.Group(children),
	)
}

type CardHeaderProps struct {
	Title       string
	Description string
}

func CardHeader(props CardHeaderProps) g.Node {
	return Div(Class("flex flex-col space-y-1.5 p-6"),
		H3(Class("text-2xl font-semibold leading-none _tracking-tight"),
			g.Text(props.Title),
		),
		P(Class("text-sm text-[#717179]"),
			g.Text(props.Description),
		),
	)
}

func CardFooter(children ...g.Node) g.Node {
	return Div(Class("items-center p-6 pt-0 flex justify-between"),
		g.Group(children),
	)
}

func CardContent(children ...g.Node) g.Node {
	return Div(Class("p-6 pt-0"),
		g.Group(children),
	)
}

func StoryCard() g.Node {
	return Div(
		story(
			Card(
				CardHeader(CardHeaderProps{
					Title:       "Title",
					Description: "Description",
				}),
				CardContent(
					P(g.Text("Card content")),
				),
				CardFooter(
					P(g.Text("Card footer")),
				),
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
