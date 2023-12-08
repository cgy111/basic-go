.PHONY: docker
docker:
	@rm webook || true
	#@GOOS=linux GOARCH=arm go build -o webook .      Mac上用该命令

	#Windows上用以下命令
	@$env:GOOS="linux"
	@$env:GOARCH="arm"
	@go build -o webook .
	@docker rmi -f cgy11/webook:v.0.0.1
	@docker build -t cgy11/webook:v.0.0.1 .


	#查看日志的命令   kubectl logs webook-5f6854dcb-cbkkn   和  kubectl logs -f webook-5f6854dcb-cbkkn