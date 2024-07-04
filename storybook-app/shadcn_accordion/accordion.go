package shadcn_accordion

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func Accordion() g.Node {
	return Div(Classes(Class("flex flex-col w-full")),
		AccordionItem(AccordionItemProps{
			Name: "shadcn_accordion",
			Text: "Section one",
		}, AccordionContent(g.Text("Content"))),
		AccordionItem(AccordionItemProps{
			Name: "shadcn_accordion",
			Text: "Section two",
		}, AccordionContent(g.Text("Content"))),
		AccordionItem(AccordionItemProps{
			Name: "shadcn_accordion",
			Text: "Section three",
		}, AccordionContent(g.Text("Content"))),
	)
}

type AccordionItemProps struct {
	Name string
	Text string
}

func AccordionItem(props AccordionItemProps, children ...g.Node) g.Node {
	return Div(Class("border-b [&:has(>h3_input:checked)>section]:block"),
		H3(Class("flex"),
			Label(Class("flex flex-1 items-center justify-between py-4 font-[400] hover:underline cursor-pointer"),
				Input(Class("peer hidden"), Type("radio"), Name(props.Name)),
				g.Text(props.Text),
				Span(Class("[stroke-width:2px] peer-checked:rotate-180"),
					lucide.ChevronDown(Class("block h-4 w-4 flex-shrink-0")),
				),
			),
		),
		Section(Class("hidden"), g.Group(children)),
	)
}

func AccordionContent(children ...g.Node) g.Node {
	return Div(Class("overflow-hidden text-sm font-[400]"),
		Div(Class("pb-4 pt-0"), g.Group(children)),
	)
}

func AccordionStory() g.Node {
	return Story(Accordion())
}

func Story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[350px] w-full items-center justify-center p-10"),
		g.Group(children),
	)
}
