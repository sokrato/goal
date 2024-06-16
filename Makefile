.PHONY: test

TEST ?= "./..."
# TEST=./leetcode/dp make test
# go test -timeout 30s -run ^Test_diffWaysToCompute$ github.com/sokrato/goal/leetcode/dp
test:
	go test -v $(TEST)

