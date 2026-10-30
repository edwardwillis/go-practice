.PHONY: test test-verbose cover fmt vet lint new

test:
	go test ./...

test-verbose:
	go test -v ./...

cover:
	go test -cover ./...

fmt:
	gofmt -l -w .

vet:
	go vet ./...

lint: fmt vet

# Usage: make new CATEGORY=arrays NAME=three_sum
new:
	@if [ -z "$(CATEGORY)" ] || [ -z "$(NAME)" ]; then \
		echo "usage: make new CATEGORY=<dir> NAME=<problem_name>"; exit 1; \
	fi
	@mkdir -p problems/$(CATEGORY)
	@sed "s/template/$(CATEGORY)/g; s/Solve/$(NAME)/g" template/TEMPLATE.go > problems/$(CATEGORY)/$(NAME).go
	@sed "s/template/$(CATEGORY)/g; s/Solve/$(NAME)/g" template/TEMPLATE_test.go > problems/$(CATEGORY)/$(NAME)_test.go
	@echo "created problems/$(CATEGORY)/$(NAME).go and $(NAME)_test.go"
