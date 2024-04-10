//go:build generate
// +build generate

package htmx

//go:generate npx tailwindcss -i ./src/input.css -o ./static/output.css
//go:generate npx esbuild --bundle --minify --sourcemap --outfile=static/output.js src/main.ts
