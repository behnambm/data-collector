FROM golang:1.22 AS builder

COPY ../go.mod ./go.sum /
RUN go mod download

COPY ./service/ /
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o /svc

# ----

FROM alpine:3.18

COPY ./service/config.yaml /config.yaml

COPY --from=builder /svc /svc

CMD ["/svc", "-c", "config.yaml"]