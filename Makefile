INPUT = cmd/main.go
OUTPUT = bin/app
GC = go

deps:
	$(GC) mod tidy
	$(GC) mod download
build:
	$(GC) build -o $(OUTPUT) $(INPUT)
run:
	$(MAKE) build
	$(OUTPUT)
