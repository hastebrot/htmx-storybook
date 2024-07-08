package shadcn_command

import (
	"github.com/hastebrot/gomponents-lucide/lucide"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
)

func Command(children ...g.Node) g.Node {
	return Div(Classes(
		Class("flex h-full w-full flex-col overflow-hidden"),
		Class("bg-[#FFFFFF] text-[#09090B]"),
		Class("rounded-lg border shadow-md"),
	),
		g.Group(children),
	)
}

type CommandInputProps struct {
	Placeholder string
	Value       string
}

func CommandInput(props CommandInputProps) g.Node {
	return Div(Class("flex items-center border-b border-[#E4E4E7] px-3"),
		Span(Class("[stroke-width:2px] mr-2"),
			lucide.Search(Class("block h-4 w-4 flex-shrink-0 text-[#717179]")),
		),
		Input(Classes(
			Class("flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none"),
			Class("placeholder:text-[#717179] disabled:cursor-not-allowed disabled:opacity-50"),
		),
			Type("text"),
			Placeholder(props.Placeholder),
			Value(props.Value),
		),
	)
}

type CommandEmptyProps struct {
	Text string
}

func CommandEmpty(props CommandEmptyProps) g.Node {
	return Div(Class("py-6 text-center text-sm"),
		g.Text(props.Text),
	)
}

func CommandList(children ...g.Node) g.Node {
	return Div(Class("max-h-[300px] overflow-y-auto overflow-x-hidden"),
		Role("listbox"),
		g.Group(children),
	)
}

type CommandGroupProps struct {
	Title string
}

func CommandGroup(props CommandGroupProps, children ...g.Node) g.Node {
	return Div(Class("overflow-hidden p-1 text-[#09090B]"),
		Div(Class("px-2 py-1.5 text-xs font-medium text-[#717179]"), g.Text(props.Title)),
		Div(
			Role("group"),
			g.Group(children),
		),
	)
}

type CommandItemProps struct {
	Text     string
	IconSlot g.Node
	Shortcut string
}

func CommandItem(props CommandItemProps) g.Node {
	return Div(Classes(
		Class("relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5"),
		Class("text-sm outline-none"),
		Class("hover:bg-[#F4F4F5] hover:text-[#18181B]"),
		Class("disabled:pointer-events-none disabled:opacity-50"),
	),
		Role("option"),
		g.If(props.IconSlot != nil,
			Span(Class("[stroke-width:2px] mr-2 flex-shrink-0"),
				props.IconSlot,
			),
		),
		g.Text(props.Text),
		g.If(props.Shortcut != "",
			Span(Class("ml-auto text-xs tracking-widest text-[#717179]"),
				g.Text(props.Shortcut),
			),
		),
	)
}

func CommandSeparator() g.Node {
	return Div(Class("-mx-1 h-px bg-[#E4E4E7]"),
		Role("separator"),
	)
}

func StoryCommand() g.Node {
	return Div(
		story(
			Command(
				CommandInput(CommandInputProps{
					Placeholder: "Type a command or search...",
					Value:       "Value",
				}),
				CommandList(
					CommandEmpty(CommandEmptyProps{
						Text: "No results found.",
					}),
				),
			),
		),
		story(
			Command(
				CommandInput(CommandInputProps{
					Placeholder: "Type a command or search...",
				}),
				CommandList(
					CommandGroup(CommandGroupProps{Title: "Suggestions"},
						CommandItem(CommandItemProps{
							Text:     "Calendar",
							IconSlot: lucide.Calendar(Class("block h-4 w-4")),
						}),
						CommandItem(CommandItemProps{
							Text:     "Search Emoji",
							IconSlot: lucide.Smile(Class("block h-4 w-4")),
						}),
						CommandItem(CommandItemProps{
							Text:     "Calculator",
							IconSlot: lucide.Calculator(Class("block h-4 w-4")),
						}),
					),
					CommandSeparator(),
					CommandGroup(CommandGroupProps{Title: "Settings"},
						CommandItem(CommandItemProps{
							Text:     "Profile",
							IconSlot: lucide.User(Class("block h-4 w-4")),
							Shortcut: "⌘P",
						}),
						CommandItem(CommandItemProps{
							Text:     "Billing",
							IconSlot: lucide.CreditCard(Class("block h-4 w-4")),
							Shortcut: "⌘B",
						}),
						CommandItem(CommandItemProps{
							Text:     "Settings",
							IconSlot: lucide.Settings(Class("block h-4 w-4")),
							Shortcut: "⌘S",
						}),
					),
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
