.PHONY: docker
docker:
	@rm webook || true
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook .
	@docker rmi -f cgy11/webook:v.0.0.1
	@docker build -t cgy11/webook:v.0.0.1 .
