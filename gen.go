//go:build generate
// +build generate

package htmx

//go:generate npx tailwindcss -i ./src/input.css -o ./out/out.css
