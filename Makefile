test:
	go test -race ./... -coverprofile=profile.out -covermode=atomic

testv:
	go test -race ./... -v -coverprofile=profile.out -covermode=atomic

coverage:
	go tool cover -html=profile.out

run:
	go build && ./ayanami

help:
	go build && ./ayanami -h
