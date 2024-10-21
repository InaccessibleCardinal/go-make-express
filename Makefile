
tst:
	go test ./... -coverprofile cover.out ./exclude-coverage.sh && go tool cover -html=cover.out