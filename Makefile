build-svc:
	@go build -o service/bin/svc service/*.go

build-svc3:
	@go build -o collector/bin/svc collector/*.go

svc1: build-svc
	@service/bin/svc -c svc1_config.yaml

svc2: build-svc
	@service/bin/svc -c svc2_config.yaml

svc3: build-svc3
	@collector/bin/svc -c svc3_config.yaml

