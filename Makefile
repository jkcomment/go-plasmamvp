build:
	go build cmd/plasma/*.go

install:
	go install cmd/plasma

clean:
	rm -f cmd/plasma/plasma
