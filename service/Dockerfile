FROM alpine:3.18

# load default config
COPY ./config.yaml /config.yaml

# move previously built binary to the image
COPY ./bin/svc /svc

# use the default config
# during run time if there are any environment variables avaialble, they will be overriden
CMD ["/svc", "-c", "config.yaml"]
