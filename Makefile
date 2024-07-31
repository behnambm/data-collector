build-svc:
	@go build -o bin/svc service/*.go

build-svc3:
	@go build -o bin/svc3 collector/*.go

svc1: build-svc
	@bin/svc -c svc1_config.yaml

svc2: build-svc
	@bin/svc -c svc2_config.yaml

svc3: build-svc3
	@bin/svc3 -c svc3_config.yaml

