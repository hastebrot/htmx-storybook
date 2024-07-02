package story_page

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	. "storybook-app/model"
)

func StoryPage() g.Node {
	menu := MenuRoot{
		MenuItems: []MenuItem{
			{
				Text: "Document",
				Type: TypeDocument,
			},
			{
				Text: "Category",
				Type: TypeCategory,
				MenuItems: []MenuItem{
					{
						Text: "Component",
						Type: TypeComponent,
						MenuItems: []MenuItem{
							{
								Text: "Story",
								Type: TypeStory,
							},
						},
					},
					{
						Text: "Folder",
						Type: TypeFolder,
						MenuItems: []MenuItem{
							{
								Text: "Document",
								Type: TypeDocument,
							},
							{
								Text: "Component",
								Type: TypeComponent,
								MenuItems: []MenuItem{
									{
										Text: "Story",
										Type: TypeStory,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return Div(Classes(Class("grid grid-cols-[auto_1fr] bg-[#F7F9FC]")),
		Div(Class("flex flex-col w-[250px] overflow-y-scroll"),
			sidebar(menu),
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

func sidebar(menu MenuRoot) g.Node {
	return Div(Class("p-[20px] w-full h-full flex flex-col"),
		sidebarHeader(),
		sidebarSearchField(),
		sidebarExplorer(menu),
	)
}

func sidebarHeader() g.Node {
	return Header(Div(Class("flex items-center justify-between"),
		Div(Class("text-[14px] font-[700] text-[rgb(51,51,51)] flex w-full items-center min-h-[22px]"),
			g.Text("storybook-app"),
		),
		Div(Class("inline-block cursor-pointer"),
			Button(Class("rounded-[3em] cursor-pointer flex items-center justify-center text-[rgba(51,51,51,0.7)] p-[7px] [box-shadow:rgba(51,51,51,0.2)_0_0_0_1px_inset] hover:[box-shadow:rgba(51,51,51,0.5)_0_0_0_1px_inset]"),
				Type("button"),
				g.Attr("popovertarget", "sidebarExplorerPopover"),
				hx.On("click", "htmx.showPopover(event, htmx.find('#sidebarExplorerPopover'))"),
				lucide.MoreHorizontal(Class("w-[14px] [stroke-width:3px]")),
			),
		),
		Div(
			ID("sidebarExplorerPopover"),
			Classes(
				Class("rounded-[8px] absolute z-20 m-0 p-0 min-w-[220px] max-h-[440px]"),
				Class("backdrop:bg-transparent bg-[rgba(255,255,255,0.95)]"),
			),
			Style("box-shadow: 0px 5px 5px rgba(0,0,0,0.05), 0 1px 3px rgba(0,0,0,0.1);"),
			// Style("filter: drop-shadow(0px 5px 5px rgba(0,0,0,0.05)) drop-shadow(0 1px 3px rgba(0,0,0,0.1));"),
			g.Attr("popover"),
			Div(
				sidebarHeaderMenuItem(sidebarHeaderMenuItemProps{
					text:     "About your Storybook",
					iconSlot: lucide.Info(Class("w-[14px]")),
				}),
				sidebarHeaderMenuItem(sidebarHeaderMenuItemProps{
					text:     "Documentation",
					iconSlot: lucide.ExternalLink(Class("w-[14px]")),
				}),
				Div(Class("border-b-[4px] border-[rgb(201,214,224)]")),
				sidebarHeaderMenuItem(sidebarHeaderMenuItemProps{
					text:       "Show sidebar",
					iconSlot:   lucide.Check(Class("w-[14px]")),
					isSelected: true,
				}),
				sidebarHeaderMenuItem(sidebarHeaderMenuItemProps{
					text:       "Show toolbar",
					iconSlot:   lucide.Check(Class("w-[14px]")),
					isSelected: true,
				}),
			),
		),
	))
}

type sidebarHeaderMenuItemProps struct {
	text       string
	iconSlot   g.Node
	isSelected bool
}

func sidebarHeaderMenuItem(props sidebarHeaderMenuItemProps) g.Node {
	return A(Classes(
		Class("text-[12px] leading-[18px] cursor-pointer select-none flex items-center justify-between p-[7px_10px]"),
		Class("hover:bg-[rgba(2,156,253,0.1)] text-[#5C6870]"),
		g.If(props.isSelected, Class("text-[#029CFD] font-[600]")),
	),
		hx.On("click", "htmx.hidePopover(event, htmx.find('#sidebarExplorerPopover'))"),
		Span(Class("flex"),
			props.iconSlot,
		),
		Span(Class("flex flex-col flex-1 pl-[10px]"),
			g.Text(props.text),
		),
	)
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

func sidebarExplorer(menu MenuRoot) g.Node {
	menuItems := collectMenuItemsFromRoot(menu)

	return Nav(Class("mt-[16px]"),
		Div(
			g.Group(g.Map(menuItems, func(item MenuItem) g.Node {
				return transformMenuItem(item)
			})),
		),
	)
}

func collectMenuItemsFromRoot(root MenuRoot) []MenuItem {
	var menuItems []MenuItem
	for _, menuItem := range root.MenuItems {
		collectMenuItems(menuItem, &menuItems)
	}
	return menuItems
}

func collectMenuItems(item MenuItem, resultMenuItems *[]MenuItem) {
	*resultMenuItems = append(*resultMenuItems, item)
	for _, menuItem := range item.MenuItems {
		collectMenuItems(menuItem, resultMenuItems)
	}
}

func transformMenuItem(item MenuItem) g.Node {
	if item.Type == TypeCategory {
		return sidebarExplorerCategory(sidebarExplorerCategoryProps{
			text:       item.Text,
			isExpanded: false,
		})
	}
	if item.Type == TypeComponent {
		return sidebarExplorerItemComponent(sidebarExplorerItemComponentProps{
			text:       item.Text,
			isExpanded: false,
		})
	}
	if item.Type == TypeDocument {
		return sidebarExplorerItemDocument(sidebarExplorerItemDocumentProps{
			text: item.Text,
		})
	}
	if item.Type == TypeFolder {
		return sidebarExplorerItemFolder(sidebarExplorerItemFolderProps{
			text:       item.Text,
			isExpanded: false,
		})
	}
	if item.Type == TypeStory {
		return sidebarExplorerItemStory(sidebarExplorerItemStoryProps{
			text:     item.Text,
			isActive: false,
		})
	}
	return nil
}

type sidebarExplorerCategoryProps struct {
	text       string
	isExpanded bool
}

func sidebarExplorerCategory(props sidebarExplorerCategoryProps) g.Node {
	return Div(Class("flex items-center mt-[16px] mb-[4px] min-h-[20px]"),
		Span(Classes(
			g.If(props.isExpanded, Class("rotate-90")),
			Class("mt-[4px] ml-[-2px] mr-[2px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent")),
		),
		Button(Class("ml-[6px] text-[rgb(153,153,153)] uppercase tracking-[0.35em] text-[11px] leading-[16px] font-[900]"),
			g.Text(props.text),
		),
	)
}

type sidebarExplorerItemDocumentProps struct {
	text string
}

func sidebarExplorerItemDocument(props sidebarExplorerItemDocumentProps) g.Node {
	return Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] py-[2px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
		Span(Class("invisible mt-[4px] ml-[-2px] mr-[2px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent")),
		Span(Class("p-[1px] ml-[5px] mr-[6px] text-[rgb(255,131,0)]"),
			lucide.FileText(Class("block w-[14px] flex-shrink-0")),
		),
		Span(g.Text(props.text)),
	))
}

type sidebarExplorerItemComponentProps struct {
	text       string
	isExpanded bool
}

func sidebarExplorerItemComponent(props sidebarExplorerItemComponentProps) g.Node {
	return Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] py-[2px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
		Span(Classes(
			g.If(props.isExpanded, Class("rotate-90")),
			Class("mt-[4px] ml-[-2px] mr-[2px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent")),
		),
		Span(Class("p-[1px] ml-[5px] mr-[6px] text-[rgb(30,167,253)]"),
			lucide.Grid2x2(Class("block w-[14px] flex-shrink-0")),
		),
		Span(g.Text(props.text)),
	))
}

type sidebarExplorerItemStoryProps struct {
	text     string
	isActive bool
}

func sidebarExplorerItemStory(props sidebarExplorerItemStoryProps) g.Node {
	return Div(A(Classes(
		Class("flex items-center text-[13px] text-[rgb(51,51,51)] py-[2px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
		g.If(props.isActive, Class("bg-[rgb(30,167,253)] hover:!bg-[rgb(30,167,253)] text-[rgb(255,255,255)] font-[600]")),
	),
		Span(Class("ml-[22px]")),
		Span(Classes(
			Class("p-[1px] ml-[5px] mr-[6px] text-[rgb(55,213,211)]"),
			g.If(props.isActive, Class("text-[rgb(255,255,255)]")),
		),
			lucide.Bookmark(Class("block w-[14px] flex-shrink-0")),
		),
		Span(g.Text(props.text)),
	))
}

type sidebarExplorerItemFolderProps struct {
	text       string
	isExpanded bool
}

func sidebarExplorerItemFolder(props sidebarExplorerItemFolderProps) g.Node {
	return Div(A(Class("flex items-center text-[13px] text-[rgb(51,51,51)] py-[2px] cursor-pointer hover:bg-[rgba(30,167,253,0.07)] hover:outline-none"),
		Span(Classes(
			g.If(props.isExpanded, Class("rotate-90")),
			Class("mt-[4px] ml-[-2px] mr-[2px] inline-block border-l-[3px] border-l-[rgba(153,153,153,0.6)] border-y-[3px] border-y-transparent")),
		),
		Span(Class("p-[1px] ml-[5px] mr-[6px] text-[rgb(119,77,215)]"),
			lucide.Folder(Class("block w-[14px] flex-shrink-0")),
		),
		Span(g.Text(props.text)),
	))
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
				toolbarIconButton(lucide.Fullscreen(Class("w-[16px]"))),
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
