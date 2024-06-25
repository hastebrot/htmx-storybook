package onyx_button

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func StoryPage() g.Node {
	return Div(Classes(Class("grid grid-cols-[auto_1fr] bg-[#F7F9FC]")),
		Div(Class("flex flex-col w-[250px] overflow-y-scroll"),
			sidebar(),
		),
		Div(Class("p-[10px] pl-0 flex flex-col"),
			Div(
				Class("h-full bg-[#FFFFFF] [box-shadow:rgba(0,0,0,0.1)_0_1px_5px_0] rounded-[4px] overflow-hidden flex flex-col"),
				toolbar(),
				Div(Class("flex-1")),
			),
		),
	)
}

func sidebar() g.Node {
	return Div(Class("p-[20px] w-full h-full flex flex-col"),
		sidebarHeader(),
		sidebarSearchField(),
		sidebarExplorer(),
	)
}

func sidebarHeader() g.Node {
	return Header(Div(Class("flex items-center justify-between"),
		Div(Class("text-[14px] font-[700] text-[rgb(51,51,51)] flex w-full items-center min-h-[22px]"),
			g.Text("storybook-app"),
		),
		Div(Class("inline-block cursor-pointer"),
			Button(Class("rounded-[3em] cursor-pointer flex items-center justify-center text-[rgba(51,51,51,0.7)] p-[7px] [box-shadow:rgba(51,51,51,0.2)_0_0_0_1px_inset] hover:[box-shadow:rgba(51,51,51,0.5)_0_0_0_1px_inset]"),
				lucide.MoreHorizontal(Class("w-[14px] [stroke-width:3px]")),
			),
		),
	))
}

func sidebarSearchField() g.Node {
	return Div(Class("relative flex flex-col mt-[16px]"),
		Div(Class("absolute top-[8px] left-[10px] z-[1] pointer-events-none text-[rgb(102,102,102)]"),
			lucide.Search(Class("w-[12px]")),
		),
		Input(Class("appearance-none h-[28px] px-[28px] border border-[rgba(153,153,153,0.4)] bg-transparent rounded-[28px] text-[12px] text-[rgb(51,51,51)] placeholder:text-[rgb(102,102,102)] focus:outline-none focus:border-[rgb(30,167,253)] focus:bg-[rgb(246,249,252)]"),
			Type("search"),
			Placeholder("Find components"),
		),
	)
}

func sidebarExplorer() g.Node {
	return Nav(Class("mt-[16px]"),
		Div(
			Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] p-[3px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
				lucide.FileText(Class("block w-[14px] p-[1px] ml-[3px] mr-[5px] flex-shrink-0 text-[rgb(255,131,0)]")),
				Span(g.Text("Document")),
			)),

			sidebarExplorerSection(),

			Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] p-[3px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
				Span(Class("mt-[4px] mr-[7px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent rotate-90")),
				lucide.Box(Class("block w-[14px] p-[1px] ml-[3px] mr-[5px] flex-shrink-0 text-[rgb(30,167,253)]")),
				Span(g.Text("Component")),
			)),

			Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] p-[3px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
				Span(Class("mt-[4px] mr-[7px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent")),
				lucide.Box(Class("block w-[14px] p-[1px] ml-[3px] mr-[5px] flex-shrink-0 text-[rgb(30,167,253)]")),
				Span(g.Text("Component")),
			)),

			Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] p-[3px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
				Span(Class("ml-[28px]")),
				lucide.Bookmark(Class("block w-[14px] p-[1px] ml-[3px] mr-[5px] flex-shrink-0 text-[rgb(30,167,253)]")),
				Span(g.Text("Example")),
			)),
		),
	)
}

func sidebarExplorerSection() g.Node {
	return Div(Class("flex items-center mt-[16px] mb-[4px] _px-[20px] min-h-[20px]"),
		Span(Class("mt-[4px] mr-[7px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent rotate-90")),
		Button(Class("text-[rgb(153,153,153)] uppercase tracking-[0.35em] text-[11px] leading-[16px] font-[900]"), g.Text("Components")),
	)
}

func toolbar() g.Node {
	return Div(Class("flex text-[rgb(153,153,153)] bg-[rgb(255,255,255)] h-[40px] [box-shadow:rgba(0,0,0,0.1)_0_-1px_0_0_inset]"),
		Div(Class("flex w-full justify-between flex-nowrap"),
			Div(Class("flex whitespace-nowrap shrink-0"),
				toolbarButton(toolbarButtonProps{text: "Canvas", isActive: true}),
				toolbarButton(toolbarButtonProps{text: "Docs", isActive: false}),
				toolbarDivider(),
				toolbarIconButton(lucide.RefreshCw(Class("w-[16px]"))),
				toolbarIconButton(lucide.ZoomIn(Class("w-[16px]"))),
				toolbarIconButton(lucide.ZoomOut(Class("w-[16px]"))),
				toolbarIconButton(lucide.Search(Class("w-[16px]"))),
			),
			Div(Class("flex whitespace-nowrap shrink-0 ml-[30px] mr-[3px]"),
				toolbarIconButton(lucide.ExternalLink(Class("w-[16px]"))),
			),
		),
	)
}

type toolbarButtonProps struct {
	text     string
	isActive bool
}

func toolbarButton(props toolbarButtonProps) g.Node {
	if props.isActive {
		return Button(Class("border-y-[3px] border-t-transparent border-b-[rgb(30,167,253)] text-[rgb(30,167,253)] px-[15px] inline-flex items-center text-[13px] leading-[12px] font-semibold cursor-pointer"), g.Text(props.text))
	}

	return Button(Class("border-y-[3px] border-t-transparent border-b-transparent text-[rgb(153,153,153)] px-[15px] inline-flex items-center text-[13px] leading-[12px] font-semibold cursor-pointer"), g.Text(props.text))
}

func toolbarDivider() g.Node {
	return Span(Class("w-[1px] h-[20px] bg-[rgba(0,0,0,0.1)] mt-[10px] ml-[6px] mr-[2px]"))
}

func toolbarIconButton(children ...g.Node) g.Node {
	return Button(Class("inline-flex items-center justify-center bg-transparent ml-[4px] mt-[6px] p-[8px_7px] cursor-pointer rounded-[4px] h-[28px] hover:bg-[rgba(30,167,253,0.1)] hover:text-[rgb(30,167,253)]"), g.Group(children))
}
