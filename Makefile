build:
	go build -o bin/schedviz ./cmd/schedviz

build_example:
	go build -o bin/load ./examples/load

run:
	go run ./cmd/schedviz

run_example:
	go run ./examles/load



