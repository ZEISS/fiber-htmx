package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/accordions"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/paginations"
	pagination "github.com/zeiss/fiber-htmx/components/paginations"
	"github.com/zeiss/fiber-htmx/components/selects"
	"github.com/zeiss/fiber-htmx/components/tables"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	Addr string
}

// DemoRow ...
type DemoRow struct {
	ID   int
	Name string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")

	rootCmd.SilenceUsage = true
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	app := fiber.New()
	app.Get("/", htmx.NewCompHandler(indexPage))

	app.Post("/api/respond", htmx.NewHtmxHandler(func(hx *htmx.Htmx) error {
		if !hx.IsHxRequest() {
			return nil
		}

		_, err := hx.WriteHTML("<div>New Content</div>")
		return err
	}))

	err := app.Listen(cfg.Flags.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

var indexPage = htmx.HTML5(htmx.HTML5Props{
	Title:    "index",
	Language: "en",
	Head: []htmx.Node{
		htmx.Link(htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui@4.7.0/dist/full.min.css"), htmx.Attribute("rel", "stylesheet"), htmx.Attribute("type", "text/css")),
		htmx.Script(htmx.Attribute("src", "https://unpkg.com/htmx.org@1.9.10"), htmx.Attribute("type", "application/javascript")),
		htmx.Script(htmx.Attribute("src", "https://cdn.tailwindcss.com"), htmx.Attribute("type", "application/javascript")),
	},
	Body: []htmx.Node{
		htmx.Div(
			htmx.ClassNames{"bg-base-100": true},
			Navbar(NavbarProps{}), htmx.Button(htmx.Text("Button"), htmx.HxPost("/api/respond"), htmx.HxSwap("outerHTML"), htmx.ClassNames{"btn": true}),
			htmx.Div(
				dropdowns.Dropdown(
					dropdowns.DropdownProps{},
					dropdowns.DropdownButton(
						dropdowns.DropdownButtonProps{},
						htmx.Text("Dropdown"),
					),
					dropdowns.DropdownMenuItems(
						dropdowns.DropdownMenuItemsProps{},
						dropdowns.DropdownMenuItem(
							dropdowns.DropdownMenuItemProps{},
							htmx.A(
								htmx.Text("Item 1"),
							),
						),
						dropdowns.DropdownMenuItem(
							dropdowns.DropdownMenuItemProps{},
							htmx.A(
								htmx.Text("Item 2"),
							),
						),
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				forms.Radio(
					forms.RadioProps{
						Name:    "radio1",
						Value:   "radio1",
						Checked: true,
					},
				),
				forms.Radio(
					forms.RadioProps{
						Name:  "radio1",
						Value: "radio2",
					},
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				forms.Checkbox(
					forms.CheckboxProps{
						Name:    "checkbox1",
						Value:   "checkbox1",
						Checked: true,
					},
				),
				forms.Checkbox(
					forms.CheckboxProps{
						Name:  "checkbox1",
						Value: "checkbox1",
					},
				),
				forms.CheckboxPrimary(
					forms.CheckboxProps{
						Name:  "checkbox1",
						Value: "checkbox1",
					},
				),
				forms.CheckboxSuccess(
					forms.CheckboxProps{
						Name:  "checkbox1",
						Value: "checkbox1",
					},
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				buttons.Button(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Neutral(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Secondary(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Accent(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Ghost(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Link(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
				buttons.Outline(
					buttons.ButtonProps{},
					htmx.Text("Button"),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				collapsible.Collapse(
					collapsible.CollapseProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					collapsible.CollapseTitle(
						collapsible.CollapseTitleProps{},
						htmx.Text("Title"),
					),
					collapsible.CollapseContent(
						collapsible.CollapseContentProps{},
						htmx.Text("Content"),
					),
				),
				collapsible.Collapse(
					collapsible.CollapseProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					collapsible.CollapseTitle(
						collapsible.CollapseTitleProps{},
						htmx.Text("Title with Checkbox"),
					),
					collapsible.CollapseCheckbox(
						collapsible.CollapseCheckboxProps{},
					),
					collapsible.CollapseContent(
						collapsible.CollapseContentProps{},
						htmx.Text("Content"),
					),
				),
				collapsible.CollapseArrow(
					collapsible.CollapseProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					collapsible.CollapseTitle(
						collapsible.CollapseTitleProps{},
						htmx.Text("Title with Checkbox"),
					),
					collapsible.CollapseCheckbox(
						collapsible.CollapseCheckboxProps{},
					),
					collapsible.CollapseContent(
						collapsible.CollapseContentProps{},
						htmx.Text("Content"),
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				avatars.AvatarRoundSmall(
					avatars.AvatarProps{},
					htmx.Img(
						htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
					),
				),
				avatars.AvatarRoundLarge(
					avatars.AvatarProps{},
					htmx.Img(
						htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
					),
				),
				avatars.AvatarRoundedLarge(
					avatars.AvatarProps{},
					htmx.Img(
						htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
					),
				),
				avatars.AvatarRoundedMedium(
					avatars.AvatarProps{},
					htmx.Img(
						htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
					),
				),
				avatars.AvatarRoundedSmall(
					avatars.AvatarProps{},
					htmx.Img(
						htmx.Attribute("src", "https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"),
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				accordions.Accordion(
					accordions.AccordionProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					accordions.AccordionRadio(
						accordions.AccordionRadioProps{
							Name: "accordion1",
						},
					),
					accordions.AccordionTitle(
						accordions.AccordionTitleProps{},
						htmx.Text("Accordion 1"),
					),
					accordions.AccordionContent(
						accordions.AccordionContentProps{},
						htmx.Text("Content 1"),
					),
				),
				accordions.Accordion(
					accordions.AccordionProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					accordions.AccordionRadio(
						accordions.AccordionRadioProps{
							Name: "accordion1",
						},
					),
					accordions.AccordionTitle(
						accordions.AccordionTitleProps{},
						htmx.Text("Accordion 2"),
					),
					accordions.AccordionContent(
						accordions.AccordionContentProps{},
						htmx.Text("Content 2"),
					),
				),
				accordions.Accordion(
					accordions.AccordionProps{
						ClassNames: htmx.ClassNames{
							"mb-4": true,
						},
					},
					accordions.AccordionRadio(
						accordions.AccordionRadioProps{
							Name: "accordion1",
						},
					),
					accordions.AccordionTitle(
						accordions.AccordionTitleProps{},
						htmx.Text("Accordion 1"),
					),
					accordions.AccordionContent(
						accordions.AccordionContentProps{},
						htmx.Text("Content 1"),
					),
				),
				accordions.AccordionArrow(
					accordions.AccordionProps{},
					accordions.AccordionRadio(
						accordions.AccordionRadioProps{
							Name: "accordion1",
						},
					),
					accordions.AccordionTitle(
						accordions.AccordionTitleProps{},
						htmx.Text("Accordion 2"),
					),
					accordions.AccordionContent(
						accordions.AccordionContentProps{},
						htmx.Text("Content 2"),
					),
				),
				accordions.AccordionArrow(
					accordions.AccordionProps{},
					accordions.AccordionRadio(
						accordions.AccordionRadioProps{
							Name: "accordion1",
						},
					),
					accordions.AccordionTitle(
						accordions.AccordionTitleProps{},
						htmx.Text("Accordion 2"),
					),
					accordions.AccordionContent(
						accordions.AccordionContentProps{},
						htmx.Text("Content 2"),
					),
				),
			),
			htmx.Div(
				tables.Table[DemoRow](
					tables.TableProps[DemoRow]{
						Columns: tables.Columns[DemoRow]{
							{
								ID:          "id",
								AccessorKey: "ID",
								Header: func(p tables.TableProps[DemoRow]) htmx.Node {
									return htmx.Th(htmx.Text("ID"))
								},
								Cell: func(p tables.TableProps[DemoRow], row DemoRow) htmx.Node {
									return htmx.Td(
										htmx.Text(strconv.Itoa(row.ID)),
									)
								},
							},
							{
								ID:          "name",
								AccessorKey: "Name",
								Header: func(p tables.TableProps[DemoRow]) htmx.Node {
									return htmx.Th(htmx.Text("Name"))
								},
								Cell: func(p tables.TableProps[DemoRow], row DemoRow) htmx.Node {
									return htmx.Td(htmx.Text(row.Name))
								},
							},
						},
						Rows: tables.NewRows[DemoRow]([]DemoRow{
							{
								ID:   1,
								Name: "Name 1",
							},
							{
								ID:   2,
								Name: "Name 2",
							},
						}),
					},
					htmx.ID("data-table"),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				selects.Select(
					selects.SelectProps{},
					selects.Option(
						selects.OptionProps{
							Disabled: true,
						},
						htmx.Text("Option 1"),
					),
					selects.Option(
						selects.OptionProps{},
						htmx.Text("Option 2"),
					),
					selects.Option(
						selects.OptionProps{
							Selected: true,
						},
						htmx.Text("Option 3"),
					),
				),
			),
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
					"p-4":         true,
				},
				paginations.Pagination(
					pagination.PaginationProps{},
					paginations.Prev(
						pagination.PaginationProps{},
					),
					paginations.Next(
						pagination.PaginationProps{},
					),
				),
			),
		),
	},
}.WithContext(&fiber.Ctx{}))

type NavbarProps struct{}

func Navbar(p NavbarProps) htmx.Node {
	return htmx.Nav(
		htmx.ClassNames{"bg-base-100": true, "navbar": true},
		htmx.Div(htmx.ClassNames{"flex-none": true},
			htmx.Button(htmx.ClassNames{"btn btn-square btn-ghost": true},
				htmx.SVG(htmx.Attribute("xmlns", "http://www.w3.org/2000/svg"),
					htmx.Attribute("fill", "none"), htmx.Attribute("viewBox", "0 0 24 24"),
					htmx.ClassNames{"inline-block w-5 h-5 stroke-current": true},
					htmx.Path(htmx.Attribute("stroke-linecap", "round"),
						htmx.Attribute("stroke-linejoin", "round"), htmx.Attribute("stroke-width", "2"),
						htmx.Attribute("d", "M4 6h16M4 12h16M4 18h16"),
					),
				),
			),
			htmx.Input(htmx.ClassNames{"toggle theme-controller": true}, htmx.Attribute("type", "checkbox"), htmx.Attribute("value", "light")),
		),
	)
}
