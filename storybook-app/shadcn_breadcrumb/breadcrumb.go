package shadcn_breadcrumb

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

type BreadcrumbProps struct {
}

func Breadcrumb(props BreadcrumbProps, children ...g.Node) g.Node {
	return Nav(g.Group(children))
}

func BreadcrumbList(children ...g.Node) g.Node {
	return Ol(Classes(
		Class("flex flex-wrap items-center gap-1.5 sm:gap-2.5"),
		Class("break-words text-sm text-[#717179]"),
	),
		g.Group(children),
	)
}

func BreadcrumbItem(children ...g.Node) g.Node {
	return Li(Class("inline-flex items-center gap-1.5"),
		g.Group(children),
	)
}

type BreadcrumbItemLinkProps struct {
	Text string
	Href string
}

func BreadcrumbItemLink(props BreadcrumbItemLinkProps) g.Node {
	return A(Class("transition-colors hover:text-[#09090B]"),
		Href(props.Href),
		g.Text(props.Text),
	)
}

type BreadcrumbItemPageProps struct {
	Text string
}

func BreadcrumbItemPage(props BreadcrumbItemPageProps) g.Node {
	return Span(Class("font-normal text-[#09090B]"),
		Role("link"),
		g.Text(props.Text),
	)
}

func BreadcrumbSeparator() g.Node {
	return Li(
		Role("presentation"),
		Aria("hidden", "true"),
		Span(Class("[stroke-width:2px]"),
			lucide.ChevronRight(Class("block h-[24px] w-[24px] flex-shrink-0")),
		),
	)
}

func StoryBreadcrumb() g.Node {
	return story(
		Breadcrumb(BreadcrumbProps{},
			BreadcrumbList(
				BreadcrumbItem(BreadcrumbItemLink(BreadcrumbItemLinkProps{
					Text: "Home",
				})),
				BreadcrumbSeparator(),
				BreadcrumbItem(BreadcrumbItemLink(BreadcrumbItemLinkProps{
					Text: "Components",
				})),
				BreadcrumbSeparator(),
				BreadcrumbItem(BreadcrumbItemPage(BreadcrumbItemPageProps{
					Text: "Breadcrumb",
				})),
			),
		),
	)
}

func story(children ...g.Node) g.Node {
	return Div(
		Class("flex min-h-[350px] w-full items-start justify-center p-10 gap-10 flex-wrap"),
		g.Group(children),
	)
}
