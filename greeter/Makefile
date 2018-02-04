help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

test:  ## Run all the tests
	@go test

watch: ## Live reload tests
	@reflex -g '*.go' go test

cov:   ## Generate HTML coverage report
	@go test -coverprofile=cov.out && go tool cover -html=cov.out

bench: ## Run your benchmarks
	@go test --run xxx --bench Greet --benchmem --benchtime=3s