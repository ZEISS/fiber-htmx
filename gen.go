//go:build generate
// +build generate

package htmx

//go:generate npx esbuild --bundle --minify --sourcemap --outfile=dist/fiber-htmx.min.js src/main.ts
//go:generate npx esbuild --bundle --platform=neutral --packages=external --outfile=dist/fiber-htmx.esm.js src/main.ts
