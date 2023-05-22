build:
	@go build -o dist/pokedexcli .
run: build
	./dist/pokedexcli