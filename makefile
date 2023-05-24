test_setup:
	go install gotest.tools/gotestsum@v1.9.0

test: test_setup
	gotestsum