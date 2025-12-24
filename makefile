test:
	go test --count=1 -cover -v ./... 

coverage-report:
	rm -f coverage.out
	go test -coverprofile=coverage.out ./...

view-report: coverage-report
	go tool cover -html=coverage.out
