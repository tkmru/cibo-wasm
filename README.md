# cibo-wasm
Demo web page to run [cibo](https://github.com/tkmru/cibo) as WebAssembly on browser.
`cibo` is cpu emulator written by golang.

## Usage
Three files in `./demo/` need to be served from a web server. 
For example, with `goexec`:

```
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```

## How to build
`GOOS=js` and `GOARCH=wasm` environment variables to compile for WebAssembly.

```
$ GOOS=js GOARCH=wasm go build -o demo/cibo.wasm cibo-wasm.go
```
