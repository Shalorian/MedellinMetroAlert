build:
	dep ensure
	go build -o metroalert

run: build
	./metroalert