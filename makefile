test_setup:
	go install gotest.tools/gotestsum@v1.9.0

test_all: test_setup
	gotestsum ./... -timeout 400s