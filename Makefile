.PHONY: dev build

dev:
	templ generate ./internals/templates
	npm run build:css &
	go run main.go

build:
	templ generate ./internals/templates
	npm run build:css
	go build -o ./builds/app main.go