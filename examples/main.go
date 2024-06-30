package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alerts"
	"github.com/zeiss/fiber-htmx/components/avatars"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/drawers"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/menus"
	"github.com/zeiss/fiber-htmx/components/mockups"
	"github.com/zeiss/fiber-htmx/components/navbars"
	"github.com/zeiss/fiber-htmx/components/swap"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/fiber-htmx/components/toasts"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
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

func (c *exampleController) Post() error {
	return c.Render(
		htmx.Fragment(
			htmx.Div(
				htmx.HxSwapOob("true"),
				htmx.ID("field_error"),
				htmx.Text("This field is required."),
			),
		),
	)
}

func (c *exampleController) Get() error {
	msg := htmx.MessagesFromContext(c.Ctx())
	msg.Add(htmx.HtmxMessage{
		Message: "Hello, World!",
	})
	msg.Add(htmx.HtmxMessage{
		Message: "Hello, World!",
	})

	return c.Render(
		htmx.HTML5(
			htmx.HTML5Props{
				Title:    "index",
				Language: "en",
				Attributes: []htmx.Node{
					htmx.DataAttribute("theme", "light"),
				},
				Head: []htmx.Node{
					htmx.Link(
						htmx.Attribute("href", "https://cdn.jsdelivr.net/npm/daisyui/dist/full.css"),
						htmx.Attribute("rel", "stylesheet"),
						htmx.Attribute("type", "text/css"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://unpkg.com/htmx.org@2.0.0"),
						htmx.Attribute("type", "application/javascript"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://cdn.tailwindcss.com"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://unpkg.com/hyperscript.org@0.9.12"),
						htmx.Attribute("type", "application/javascript"),
					),
					htmx.Script(
						htmx.Attribute("src", "https://unpkg.com/htmx-ext-sse@2.0.0/sse.js"),
						htmx.Attribute("type", "application/javascript"),
					),
				},
			},
			htmx.Body(
				htmx.HxBoost(true),
				htmx.Fragment(
					htmx.Div(),
					htmx.Div(),
				),
				htmx.Div(
					htmx.ClassNames{},
					drawers.Drawer(
						drawers.DrawerProps{
							ID: "global-drawer",
							ClassNames: htmx.ClassNames{
								"lg:drawer-open": true,
							},
						},
						drawers.DrawerContent(
							drawers.DrawerContentProps{
								ID: "drawer",
							},
							htmx.Div(
								htmx.ClassNames{
									tailwind.OverflowAuto: true,
									tailwind.WFull:        true,
									tailwind.HScreen:      true,
									tailwind.MaxWFull:     true,
								},
								htmx.Div(
									htmx.ClassNames{
										tailwind.Flex:    true,
										tailwind.HFull:   true,
										tailwind.FlexCol: true,
										"bg-base-200":    true,
									},
									navbars.Navbar(
										navbars.NavbarProps{
											ClassNames: htmx.ClassNames{
												"navbar":      true,
												"z-10":        true,
												"border-b":    true,
												"px-3":        true,
												"bg-base-100": true,
											},
										},
										navbars.NavbarStart(
											navbars.NavbarStartProps{
												ClassNames: htmx.ClassNames{
													"gap-3": true,
												},
											},
											drawers.DrawerOpenButton(
												drawers.DrawerOpenProps{
													ID: "global-drawer",
													ClassNames: htmx.ClassNames{
														"btn-sm":      true,
														"btn-square":  true,
														"btn-primary": false,
													},
												},
												icons.Bars3Outline(
													icons.IconProps{},
												),
											),
										),
										navbars.NavbarEnd(
											navbars.NavbarEndProps{},
											swap.Swap(
												swap.SwapProps{
													ClassNames: htmx.ClassNames{
														"swap-rotate": true,
													},
												},
												htmx.Input(
													htmx.Class("theme-controller"),
													htmx.Value("dark"),
													htmx.Attribute("type", "checkbox"),
												),
												swap.SwapOn(
													swap.SwapProps{},
													icons.MoonOutlineSmall(
														icons.IconProps{},
													),
												),
												swap.SwapOff(
													swap.SwapProps{},
													icons.SunOutlineSmall(
														icons.IconProps{},
													),
												),
											),
											buttons.CircleSmall(
												buttons.ButtonProps{},
												icons.BellAlertOutlineSmall(
													icons.IconProps{},
												),
											),
											dropdowns.Dropdown(
												dropdowns.DropdownProps{
													ClassNames: htmx.ClassNames{
														"dropdown-end": true,
													},
												},
												dropdowns.DropdownButton(
													dropdowns.DropdownButtonProps{
														ClassNames: htmx.ClassNames{
															"btn-sm":     true,
															"btn-circle": true,
															"btn-ghost":  true,
														},
													},
													avatars.AvatarRoundSmall(
														avatars.AvatarProps{},
														htmx.Img(
															htmx.Attribute("src", "https://avatars.githubusercontent.com/u/570959?v=4"),
														),
													),
												),
												dropdowns.DropdownMenuItems(
													dropdowns.DropdownMenuItemsProps{},
													dropdowns.DropdownMenuItem(
														dropdowns.DropdownMenuItemProps{},
														htmx.A(
															htmx.Text("Profile"),
														),
													),
												),
											),
										),
									),
									htmx.Div(
										mockups.Code(
											mockups.CodeProps{},
										),
										swap.Swap(
											swap.SwapProps{
												ClassNames: htmx.ClassNames{
													"swap-rotate": true,
												},
											},
											// Write a Hyperscript that toggles a checkbox inside the current element
											htmx.HyperScript(`on click
    writeText(my previousElementSibling's innerText) on navigator.clipboard
    set <input/> in me to checked to not checked
    wait 1s
    set <input/> in me to checked to true`),
											htmx.Input(
												htmx.Value("copy"),
												htmx.Attribute("type", "checkbox"),
												htmx.Checked(),
											),
											swap.SwapOn(
												swap.SwapProps{},
												icons.MoonOutlineSmall(
													icons.IconProps{},
												),
											),
											swap.SwapOff(
												swap.SwapProps{},
												icons.SunOutlineSmall(
													icons.IconProps{},
												),
											),
										),
									),
									alerts.Info(
										alerts.AlertProps{},
										htmx.Text("Hello, World!"),
									),
									htmx.Div(
										htmx.HxSSE(),
										htmx.HxSSEConnect("/sse"),
										htmx.HxSSESwap("message"),
									),
									htmx.Div(
										htmx.ID("events"),
									),
									htmx.Div(
										forms.TextInput(
											forms.TextInputProps{},
										),
										htmx.Div(
											htmx.Merge(
												htmx.ClassNames{
													"text-bg-error": true,
												},
											),
											htmx.ID("field_error"),
										),
										htmx.Form(
											htmx.Method("POST"),
											htmx.HxSwap("none"),
											buttons.Ghost(
												buttons.ButtonProps{},
												htmx.HxPost("/error"),
												htmx.Type("submit"),
												htmx.Text("Submit"),
											),
										),
									),
									htmx.Div(
										toasts.Toast(
											toasts.ToastProps{},
											toasts.ToastAlertInfo(
												toasts.ToastProps{},
												htmx.Text("message 1"),
											),
											toasts.ToastAlertInfo(
												toasts.ToastProps{},
												htmx.Text("message 2"),
											),
											toasts.ToastAlertInfo(
												toasts.ToastProps{},
												htmx.Text("message 3"),
											),
										),
									),
									htmx.Div(
										tables.Table(
											tables.TableProps{
												Pagination: tables.TablePagination(
													tables.TablePaginationProps{},
													tables.Pagination(
														tables.PaginationProps{
															Total:  len(demoRows),
															Offset: 0,
															Limit:  10,
														},
													),
												),
											},
											tables.Columns[DemoRow]{
												{
													ID:          "id",
													AccessorKey: "ID",
													Header: func(p tables.TableProps) htmx.Node {
														return htmx.Th(htmx.Text("ID"))
													},
													Cell: func(p tables.TableProps, row DemoRow) htmx.Node {
														return htmx.Td(
															htmx.Text(strconv.Itoa(row.ID)),
														)
													},
												},
												{
													ID:          "name",
													AccessorKey: "Name",
													Header: func(p tables.TableProps) htmx.Node {
														return htmx.Th(htmx.Text("Name"))
													},
													Cell: func(p tables.TableProps, row DemoRow) htmx.Node {
														return htmx.Td(htmx.Text(row.Name))
													},
												},
											},
											demoRows,
										),
									),
								),
							),
						),
						drawers.DrawerSide(
							drawers.DrawerSideProps{
								ID: "drawer",
							},
							drawers.DrawerSideMenu(
								drawers.DrawerSideMenuProps{
									ClassNames: htmx.ClassNames{
										"border-r":    true,
										"bg-base-100": true,
										"bg-base-200": false,
									},
								},
								htmx.Nav(
									htmx.Merge(
										htmx.ClassNames{},
									),
									menus.Menu(
										menus.MenuProps{
											ClassNames: htmx.ClassNames{
												"w-full":      true,
												"bg-base-200": false,
											},
										},
										menus.MenuItem(
											menus.MenuItemProps{},
											menus.MenuLink(
												menus.MenuLinkProps{},
												htmx.Text("Home"),
											),
										),
										menus.MenuTitle(
											menus.MenuTitleProps{},
											htmx.Text("Components"),
										),
										menus.MenuItem(
											menus.MenuItemProps{},
											menus.MenuCollapsible(
												menus.MenuCollapsibleProps{},
												menus.MenuCollapsibleSummary(
													menus.MenuCollapsibleSummaryProps{},
													htmx.Text("Buttons"),
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
											menus.MenuCollapsible(
												menus.MenuCollapsibleProps{},
												menus.MenuCollapsibleSummary(
													menus.MenuCollapsibleSummaryProps{},
													htmx.Text("Forms"),
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
	app.Use(htmx.NewHtmxMessageHandler())

	app.Get("/", htmx.NewHxControllerHandler(func() htmx.Controller {
		return &exampleController{}
	}))

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")

		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			var i int
			for {
				i++
				msg := fmt.Sprintf("%d - the time is %v", i, time.Now())

				var buf bytes.Buffer
				err := htmx.Div(htmx.ID("message"), htmx.Text(msg)).Render(&buf)
				if err != nil {
					break
				}

				fmt.Fprintf(w, "event: %s\n", "message")
				fmt.Fprintf(w, "data: %s\n\n", buf.String())

				err = w.Flush()
				if err != nil {
					// Refreshing page in web browser will establish a new
					// SSE connection, but only (the last) one is alive, so
					// dead connections must be closed here.
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)

					break
				}
				time.Sleep(2 * time.Second)
			}
		}))

		return nil
	})

	app.Post("/error", htmx.NewHxControllerHandler(func() htmx.Controller {
		return &exampleController{}
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
