build-svc:
	@CGO_ENABLED=0 go build -o service/bin/svc service/*.go

svc1: build-svc
	@service/bin/svc -c configs/svc1_config.yaml

svc2: build-svc
	@service/bin/svc -c configs/svc2_config.yaml

svc3:
	@# didn't disable the CGO because the sqlite needs it
	@go build -o collector/bin/svc3 collector/*.go
	@collector/bin/svc3 -c configs/svc3_config.yaml

down:
	docker compose down

up:
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o service/bin/svc service/*.go
	docker build -t data-collector:latest -f Dockerfile .
	docker compose up $(ARGS)

test:
	@go test ./... -v

