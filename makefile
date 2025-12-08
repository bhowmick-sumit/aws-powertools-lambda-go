test:
	go test -cover -v ./... 

view-coverage:
	go tool cover -html=cover.out
