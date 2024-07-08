package shadcn_table

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Table(children ...g.Node) g.Node {
	return h.Div(h.Class("relative w-full overflow-auto"),
		h.Table(h.Class("w-full [caption-side:bottom] text-sm"),
			g.Group(children),
		),
	)
}

func TableCaption(children ...g.Node) g.Node {
	return h.Caption(h.Class("mt-4 text-sm text-[#717179]"),
		g.Group(children),
	)
}

func TableHeader(children ...g.Node) g.Node {
	return h.THead(h.Class("[&_tr]:border-b"),
		g.Group(children),
	)
}

func TableBody(children ...g.Node) g.Node {
	return h.TBody(h.Class("[&_tr:last-child]:border-0"),
		g.Group(children),
	)
}

func TableFooter(children ...g.Node) g.Node {
	return h.TFoot(h.Class("border-t bg-[#F4F4F5]/50 font-medium [&>tr]:last:border-b-0"),
		g.Group(children),
	)
}

func TableRow(children ...g.Node) g.Node {
	return h.Tr(h.Class("border-b transition-colors hover:bg-[#F4F4F5]/50"),
		g.Group(children),
	)
}

func TableHead(children ...g.Node) g.Node {
	return h.Th(h.Class("h-12 px-4 text-left align-middle font-medium text-[#717179] w-[100px]"),
		g.Group(children),
	)
}

func TableCell(children ...g.Node) g.Node {
	return h.Td(h.Class("p-4 align-middle"),
		g.Group(children),
	)
}

func StoryTable() g.Node {
	return h.Div(
		story(
			Table(
				TableCaption(g.Text("A list of your recent invoices.")),
				TableHeader(
					TableRow(
						TableHead(g.Text("Invoice")),
						TableHead(g.Text("Status")),
						TableHead(g.Text("Method")),
						TableHead(h.Div(h.Class("text-right"), g.Text("Amount"))),
					),
				),
				TableBody(
					TableRow(
						TableCell(g.Text("INV001")),
						TableCell(g.Text("Paid")),
						TableCell(g.Text("Credit Card")),
						TableCell(h.Div(h.Class("text-right tabular-nums"), g.Text("$250.00"))),
					),
					TableRow(
						TableCell(g.Text("INV002")),
						TableCell(g.Text("Unpaid")),
						TableCell(g.Text("Bank Transfer")),
						TableCell(h.Div(h.Class("text-right tabular-nums"), g.Text("$350.00"))),
					),
				),
				TableFooter(
					TableRow(
						TableCell(h.ColSpan("3"), g.Text("Total")),
						TableCell(h.Div(h.Class("text-right tabular-nums"), g.Text("$600.00"))),
					),
				),
			),
		),
	)
}

func story(children ...g.Node) g.Node {
	return h.Div(
		h.Class("flex min-h-[150px] w-full items-start justify-center p-10 gap-10"),
		g.Group(children),
	)
}
