svc1:
	@CGO_ENABLED=0 go build -o service1/bin/svc service1/*.go
	@service1/bin/svc -c service1/svc1_config.yaml

svc2:
	@CGO_ENABLED=0 go build -o service2/bin/svc service2/*.go
	@service2/bin/svc -c service2/svc2_config.yaml

svc3:
	@# didn't disable the CGO because the sqlite needs it
	@go build -o service3/bin/svc3 service3/*.go
	@service3/bin/svc3 -c service3/svc3_config.yaml

down:
	docker compose down

up:
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o service1/bin/svc service1/*.go
	CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o service2/bin/svc service2/*.go
	docker build -t data-collector-svc1:latest -f service1/Dockerfile ./service1
	docker build -t data-collector-svc2:latest -f service2/Dockerfile ./service2
	docker compose up $(ARGS)

test:
	@go test ./... -v

