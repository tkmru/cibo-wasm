# cibo-wasm
Demo web page to run [cibo](https://github.com/tkmru/cibo) as WebAssembly on browser.
`cibo` is cpu emulator written by golang.

```
$ goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
```
