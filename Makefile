

build-svc:
	@go build -o bin/svc service/*.go

svc1: build-svc
	@bin/svc -c svc1_config.yaml

svc2: build-svc
	@bin/svc -c svc2_config.yaml
