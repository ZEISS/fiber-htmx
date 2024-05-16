package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/zeiss/fiber-htmx/components/accordions"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/collapsible"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/kbds"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/fiber-htmx/components/modals"
	"github.com/zeiss/fiber-htmx/components/progress"
	"github.com/zeiss/fiber-htmx/components/skeletons"
	"github.com/zeiss/fiber-htmx/components/stats"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/toasts"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
	htmx "github.com/zeiss/fiber-htmx"
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

var demoRows = []DemoRow{
	{
		ID:   1,
		Name: "Name 1",
	},
	{
		ID:   2,
		Name: "Name 2",
	},
	{
		ID:   3,
		Name: "Name 3",
	},
	{
		ID:   4,
		Name: "Name 4",
	},
	{
		ID:   5,
		Name: "Name 5",
	},
	{
		ID:   6,
		Name: "Name 6",
	},
	{
		ID:   7,
		Name: "Name 7",
	},
	{
		ID:   8,
		Name: "Name 8",
	},
	{
		ID:   9,
		Name: "Name 9",
	},
	{
		ID:   10,
		Name: "Name 10",
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":3000", "addr")

	rootCmd.SilenceUsage = true
}

type exampleController struct {
	htmx.UnimplementedController
}

func (c *exampleController) Error(err error) error {
	return htmx.RenderComp(
		c.Ctx(),
		htmx.HTML5(
			htmx.HTML5Props{
				Title:    "error",
				Language: "en",
				Head:     []htmx.Node{},
			},
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
				},
				htmx.Text(err.Error()),
			),
		),
		htmx.RenderStatusCode(err),
	)
}

func (c *exampleController) Get() error {
	return htmx.RenderComp(
		c.Ctx(),
		htmx.HTML5(
			htmx.HTML5Props{
				Title:    "index",
				Language: "en",
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "/static/output.css"),
						htmx.Attribute("rel", "stylesheet"),
						htmx.Attribute("type", "text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "/static/main.js"),
						htmx.Attribute("type", "application/javascript"),
					),
				},
			},
			htmx.Div(
				htmx.ClassNames{
					"bg-base-100": true,
				},
				Navbar(
					NavbarProps{
						Title: "Navbar",
					},
				),
				htmx.Button(
					htmx.Text("Button"),
					htmx.HxPost("/api/respond"),
					htmx.HxSwap("outerHTML"),
					htmx.ClassNames{"btn": true},
				),
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
					htmx.ClassNames{},
					// htmx.Range(
					// 	htmx.Div(htmx.Text("Item 1")),
					// 	htmx.Div(htmx.Text("Item 2")),
					// ).
					// 	Filter(func(int) bool { return true }).
					// 	Map(func(int) htmx.Node { return htmx.Div(htmx.Text("Hello")) }).
					// 	Group(),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					menus.Menu(
						menus.MenuProps{},
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 1"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 2"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 3"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					menus.Menu(
						menus.MenuProps{},
						menus.MenuTitle(
							menus.MenuTitleProps{},
							htmx.Text("Title"),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 1"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 2"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 3"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					menus.Menu(
						menus.MenuProps{},
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 1"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuCollapsible(
								menus.MenuCollapsibleProps{
									Open: true,
								},
								menus.MenuCollapsibleSummary(
									menus.MenuCollapsibleSummaryProps{},
									htmx.Text("Collapsible"),
								),
								menus.MenuItem(
									menus.MenuItemProps{},
									htmx.A(
										htmx.Attribute("href", "#"),
										htmx.Text("Item 1"),
									),
									htmx.A(
										htmx.Attribute("href", "#"),
										htmx.Text("Item 2"),
									),
									htmx.A(
										htmx.Attribute("href", "#"),
										htmx.Text("Item 3"),
									),
								),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href:   "#",
									Active: true,
								},
								htmx.Text("Item 1"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							menus.MenuLink(
								menus.MenuLinkProps{
									Href: "#",
								},
								icons.UserOutline(
									icons.IconProps{},
								),
								htmx.Text("Item 2"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 3"),
							),
						),
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
							htmx.Text("Collapsible"),
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
							Pagination: tables.TablePagination(
								tables.TablePaginationProps[DemoRow]{
									Pagination: tables.Pagination(
										tables.PaginationProps{
											Total:  len(demoRows),
											Offset: 0,
											Limit:  10,
										},
										// tables.Select(
										// 	tables.SelectProps{
										// 		Total:  len(demoRows),
										// 		Offset: 0,
										// 		Limit:  10,
										// 		Limits: tables.DefaultLimits,
										// 		URL:    "/api/data",
										// 	},
										// ),
									),
								},
							),
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
							Rows: tables.NewRows(demoRows),
						},
						htmx.ID("data-table"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					forms.Select(
						forms.SelectProps{},
						forms.Option(
							forms.OptionProps{
								Disabled: true,
							},
							htmx.Text("Option 1"),
						),
						forms.Option(
							forms.OptionProps{},
							htmx.Text("Option 2"),
						),
						forms.Option(
							forms.OptionProps{
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
					tables.Pagination(
						tables.PaginationProps{},
						tables.Prev(
							tables.PaginationProps{
								URL: "/api/data",
							},
						),
						tables.Next(
							tables.PaginationProps{
								URL: "/api/data",
							},
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					forms.TextInput(
						forms.TextInputProps{},
					),
					forms.TextInputBordered(
						forms.TextInputProps{},
					),
					forms.TextInputGhost(
						forms.TextInputProps{},
					),
					forms.FileInput(
						forms.FileInputProps{},
					),
					forms.FileInputAccent(
						forms.FileInputProps{},
					),
					forms.FileInputPrimary(
						forms.FileInputProps{},
					),
					forms.FileInputSuccess(
						forms.FileInputProps{},
					),
					forms.FileInputWarning(
						forms.FileInputProps{},
					),
					forms.FileInputError(
						forms.FileInputProps{},
					),
					forms.TextInputGhost(
						forms.TextInputProps{},
					),
					forms.TextInputGhost(
						forms.TextInputProps{},
					),
					forms.TextInputWithIcon(
						forms.TextInputProps{
							Icon: icons.SearchOutline(
								icons.IconProps{},
							),
						},
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					forms.Toggle(
						forms.ToggleProps{},
					),
					forms.ToggleSuccess(
						forms.ToggleProps{
							Checked:  true,
							Disabled: true,
						},
					),
					forms.ToggleWarning(
						forms.ToggleProps{
							Checked: true,
						},
					),
					forms.ToggleError(
						forms.ToggleProps{
							Checked: true,
						},
					),
					forms.ToggleInfo(
						forms.ToggleProps{},
					),
					forms.ToggleLabel(
						forms.ToggleProps{},
						htmx.Text("Toggle"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					toasts.ToastEnd(
						toasts.ToastProps{},
						toasts.ToastAlertInfo(
							htmx.Text("Info"),
						),
						toasts.ToastAlertError(
							htmx.Text("Error"),
						),
						toasts.ToastAlertSuccess(
							htmx.Text("Success"),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					stats.Stats(
						stats.StatsProps{},
						stats.Stat(
							stats.StatProps{},
							stats.Figure(
								stats.FigureProps{},
								icons.BoltSlashOutline(
									icons.IconProps{},
								),
							),
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Page Views"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("89,400"),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text("Total page views in the last 30 days"),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Figure(
								stats.FigureProps{},
								icons.HeartOutline(
									icons.IconProps{},
								),
							),
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Likes"),
							),
							stats.Value(
								stats.ValueProps{
									ClassNames: htmx.ClassNames{
										"text-primary": true,
									},
								},
								htmx.Text("25.6K"),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text("21% more than last month"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					stats.StatsVertical(
						stats.StatsProps{},
						stats.Stat(
							stats.StatProps{},
							stats.Figure(
								stats.FigureProps{},
								icons.BoltSlashOutline(
									icons.IconProps{},
								),
							),
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Page Views"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("89,400"),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text("Total page views in the last 30 days"),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Figure(
								stats.FigureProps{},
								icons.HeartOutline(
									icons.IconProps{},
								),
							),
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Likes"),
							),
							stats.Value(
								stats.ValueProps{
									ClassNames: htmx.ClassNames{
										"text-primary": true,
									},
								},
								htmx.Text("25.6K"),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text("21% more than last month"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					stats.Stats(
						stats.StatsProps{},
						stats.Stat(
							stats.StatProps{},
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Account balance"),
							),
							stats.Value(
								stats.ValueProps{},
								htmx.Text("$89,400"),
							),
							stats.Actions(
								stats.ActionsProps{},
								buttons.Success(
									buttons.ButtonProps{
										ClassNames: htmx.ClassNames{
											"btn-sm": true,
										},
									},
									htmx.Text("Add funds"),
								),
							),
						),
						stats.Stat(
							stats.StatProps{},
							stats.Figure(
								stats.FigureProps{},
								icons.HeartOutline(
									icons.IconProps{},
								),
							),
							stats.Title(
								stats.TitleProps{},
								htmx.Text("Total Likes"),
							),
							stats.Value(
								stats.ValueProps{
									ClassNames: htmx.ClassNames{
										"text-primary": true,
									},
								},
								htmx.Text("25.6K"),
							),
							stats.Description(
								stats.DescriptionProps{},
								htmx.Text("21% more than last month"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					cards.Card(
						cards.CardProps{},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Card Title"),
							),
							htmx.Text("Card Body"),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Primary(
									buttons.ButtonProps{},
									htmx.Text("Button"),
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
					forms.Textarea(
						forms.TextareaProps{
							Placeholder: "Textarea",
						},
					),
					forms.Textarea(
						forms.TextareaProps{
							Placeholder: "Textarea",
							Disabled:    true,
						},
					),
					forms.TextareaBordered(
						forms.TextareaProps{
							Placeholder: "Textarea",
						},
					),
					forms.TextareaGhost(
						forms.TextareaProps{
							Placeholder: "Textarea",
						},
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					progress.Progress(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
							Max:   100,
							Value: 50,
						},
					),
					progress.ProgressPrimary(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
							Max:   100,
							Value: 50,
						},
					),
					progress.ProgressSecondary(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
							Max:   100,
							Value: 50,
						},
					),
					progress.ProgressSuccess(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
							Max:   100,
							Value: 50,
						},
					),
					progress.ProgressInfo(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
							Max:   100,
							Value: 50,
						},
					),
					progress.ProgressIntermediate(
						progress.ProgressProps{
							ClassNames: htmx.ClassNames{
								"block": true,
								"m-4":   true,
							},
						},
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					skeletons.Skeleton(
						skeletons.SkeletonProps{
							Width:  32,
							Height: 32,
						},
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					menus.Menu(
						menus.MenuProps{},
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 1"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 2"),
							),
						),
						menus.MenuItem(
							menus.MenuItemProps{},
							htmx.A(
								htmx.Attribute("href", "#"),
								htmx.Text("Item 3"),
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					kbds.Kbd(
						kbds.KbdProps{},
						htmx.Text("Ctrl"),
					),
					kbds.Kbd(
						kbds.KbdProps{},
						htmx.Text("Alt"),
					),
					kbds.Kbd(
						kbds.KbdProps{},
						htmx.Text("Shift"),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					buttons.Button(
						buttons.ButtonProps{},
						htmx.Text("Open Modal"),
						htmx.OnClick("my_modal_1.showModal()"),
					),
					modals.Modal(
						modals.ModalProps{
							ID: "my_modal_1",
						},
						modals.ModalAction(
							modals.ModalActionProps{},
							modals.ModalCloseButton(
								modals.ModalCloseButtonProps{},
							),
						),
					),
				),
				htmx.Div(
					htmx.ClassNames{
						"bg-base-100": true,
						"p-4":         true,
					},
					drawers.Drawer(
						drawers.DrawerProps{
							ID: "drawer",
						},
						drawers.DrawerContent(
							drawers.DrawerContentProps{
								ID: "drawer",
							},
							drawers.DrawerOpenButton(
								drawers.DrawerOpenProps{
									ID: "drawer",
								},
								htmx.Text("Open Drawer"),
							),
						),
						drawers.DrawerSide(
							drawers.DrawerSideProps{
								ID: "drawer",
							},
							drawers.DrawerSideMenu(
								drawers.DrawerSideMenuProps{},
								drawers.DrawerSideMenu(
									drawers.DrawerSideMenuProps{},
									drawers.DrawerSideMenuItem(
										drawers.DrawerSideMenuItemProps{},
										htmx.Text("Item 1"),
									),
									drawers.DrawerSideMenuItem(
										drawers.DrawerSideMenuItemProps{},
										htmx.Text("Item 2"),
									),
									drawers.DrawerSideMenuItem(
										drawers.DrawerSideMenuItemProps{},
										htmx.Text("Item 3"),
									),
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
					drawers.Drawer(
						drawers.DrawerProps{
							ID: "drawer2",
						},
						drawers.DrawerContent(
							drawers.DrawerContentProps{
								ID: "drawer2",
							},
							drawers.DrawerOpenButton(
								drawers.DrawerOpenProps{
									ID: "drawer2",
								},
								htmx.Text("Open Drawer 2"),
							),
						),
						drawers.DrawerSide(
							drawers.DrawerSideProps{
								ID: "drawer2",
							},
							drawers.DrawerSideMenu(
								drawers.DrawerSideMenuProps{},
								collapsible.CollapseArrow(
									collapsible.CollapseProps{
										TabIndex: 2,
										ClassNames: htmx.ClassNames{
											"mb-4": true,
										},
									},
									collapsible.CollapseTitle(
										collapsible.CollapseTitleProps{},
										htmx.Text("Title with Checkbox"),
									),
									collapsible.CollapseContent(
										collapsible.CollapseContentProps{},
										htmx.Text("Content"),
									),
								),
								collapsible.CollapseArrow(
									collapsible.CollapseProps{
										TabIndex: 3,
										ClassNames: htmx.ClassNames{
											"mb-4": true,
										},
									},
									collapsible.CollapseTitle(
										collapsible.CollapseTitleProps{},
										htmx.Text("Title with Checkbox"),
									),
									collapsible.CollapseContent(
										collapsible.CollapseContentProps{},
										htmx.Text("Content"),
									),
								),
							),
						),
					),
				),
			),
		),
	)
}

func run(_ context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(htmx.Static()),
	}))

	app.All("/", htmx.NewHxControllerHandler(&exampleController{}))

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

type NavbarProps struct {
	Title string
}

func Navbar(props NavbarProps, children ...htmx.Node) htmx.Node {
	return htmx.Nav(
		htmx.ClassNames{"bg-base-100": true, "navbar": true},
		htmx.Div(
			htmx.H3(
				htmx.ClassNames{"text-lg font-bold": true},
				htmx.Text(props.Title),
			),
		),
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
