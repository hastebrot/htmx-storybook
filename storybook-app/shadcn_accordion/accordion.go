package shadcn_accordion

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func Accordion(children ...g.Node) g.Node {
	return Div(Classes(Class("shadcn-accordion flex flex-col w-full")),
		g.Group(children),
	)
}

type AccordionItemProps struct {
	Name string
	Text string
}

func AccordionItem(props AccordionItemProps, children ...g.Node) g.Node {
	return Div(Class("shadcn-accordion-item border-b [&:has(>header_input:checked)>section]:block"),
		Header(Class("flex"),
			Label(Class("flex flex-1 items-center justify-between py-4 font-medium hover:underline cursor-pointer"),
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
	return Div(Class("shadcn-accordion-content overflow-hidden text-sm font-normal"),
		Div(Class("pb-4 pt-0"), g.Group(children)),
	)
}

func AccordionStory() g.Node {
	return Story(
		Accordion(
			AccordionItem(AccordionItemProps{
				Name: "shadcn_accordion",
				Text: "Section one",
			},
				AccordionContent(g.Text("Content one")),
			),
			AccordionItem(AccordionItemProps{
				Name: "shadcn_accordion",
				Text: "Section two",
			},
				AccordionContent(g.Text("Content two")),
			),
			AccordionItem(AccordionItemProps{
				Name: "shadcn_accordion",
				Text: "Section three",
			},
				AccordionContent(g.Text("Content three")),
			),
		),
	)
}

func Story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[350px] w-full items-start justify-center p-10"),
		g.Group(children),
	)
}
