//go:build generate
// +build generate

package htmx

//go:generate npx esbuild --bundle --minify --sourcemap --outfile=dist/out.js src/main.ts
