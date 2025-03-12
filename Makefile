
# Sample tests

export TEST1='[{"operation":"buy", "unit-cost":10.00, "quantity": 100},{"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]'
export TEST2='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000},{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]'
export TEST12='[{"operation":"buy", "unit-cost":10.00, "quantity": 100},{"operation":"sell", "unit-cost":15.00, "quantity": 50},{"operation":"sell", "unit-cost":15.00, "quantity": 50}]\n[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":20.00, "quantity": 5000},{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]'
export TEST3='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":5.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 3000}]'
export TEST4='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"buy", "unit-cost":25.00, "quantity": 5000},{"operation":"sell", "unit-cost":15.00, "quantity": 10000}]'
export TEST5='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"buy", "unit-cost":25.00, "quantity": 5000},{"operation":"sell", "unit-cost":15.00, "quantity": 10000},{"operation":"sell", "unit-cost":25.00, "quantity": 5000}]'
export TEST6='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":2.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":25.00, "quantity": 1000}]'
export TEST7='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":2.00, "quantity": 5000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":20.00, "quantity": 2000},{"operation":"sell", "unit-cost":25.00, "quantity": 1000},{"operation":"buy", "unit-cost":20.00, "quantity": 10000},{"operation":"sell", "unit-cost":15.00, "quantity": 5000},{"operation":"sell", "unit-cost":30.00, "quantity": 4350},{"operation":"sell", "unit-cost":30.00, "quantity": 650}]'
export TEST8='[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},{"operation":"sell", "unit-cost":50.00, "quantity": 10000},{"operation":"buy", "unit-cost":20.00, "quantity": 10000},{"operation":"sell", "unit-cost":50.00, "quantity": 10000}]'

# Building

.PHONY: setup
setup:
	@go mod tidy

.PHONY: build
build:
	@go build -o capital-gains ./cmd/main.go

# Code quality

.PHONY: test-coverage
test-coverage:
	@echo "==> Running tests with coverage"
	@go test -race -failfast -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

.PHONY: test
test:
	@echo "==> Running tests"
	@go clean -testcache
	@go test -race -failfast ./...

# Project-wide

.PHONY: run.sample
run.sample:
	@echo "==> Running sample test 1"
	@echo $(TEST1) | go run cmd/main.go
	@echo "==> Running sample test 2"
	@echo $(TEST2) | go run cmd/main.go
	@echo "==> Running sample test 1+2"
	@echo $(TEST12) | go run cmd/main.go
	@echo "==> Running sample test 3"
	@echo $(TEST3) | go run cmd/main.go
	@echo "==> Running sample test 4"
	@echo $(TEST4) | go run cmd/main.go
	@echo "==> Running sample test 5"
	@echo $(TEST5) | go run cmd/main.go
	@echo "==> Running sample test 6"
	@echo $(TEST6) | go run cmd/main.go
	@echo "==> Running sample test 7"
	@echo $(TEST7) | go run cmd/main.go
	@echo "==> Running sample test 8"
	@echo $(TEST8) | go run cmd/main.go