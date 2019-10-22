build:
	dep ensure
	go build -o metroalert

run: 
	./metroalert
