.PHONY: build test run clean

build:
	go build -o angelica .

test:
	go test ./...

run:
	go run .

clean:
	rm -f angelica
