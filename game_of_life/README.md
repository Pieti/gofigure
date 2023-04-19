# Game of Life on a hexgrid

Because hexagons are bestagons!


## Requirements

* [Go](https://go.dev/)
* [Ebiten](https://ebitengine.org/en/documents/install.html)


## Run the game

Compile into WebAssembly

```
go install github.com/hajimehoshi/wasmserve@latest
wasmserve .
```

Browse to http://localhost:8080/ and you should see a game!


Alternatively it can run natively
```
go run .
```
