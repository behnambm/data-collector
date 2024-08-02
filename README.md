# Overview

According to task description, there should be 3 services. 
But the project consists of two different components. (See branch `v2` for separate service implementation)

1. **Service** (2 instances): A server that listens on a port and handles RPC calls.
2. **Collector** (1 instance): A service that makes concurrent calls to the RPC services and gathers data from them.



# How to run 

### Requirements
- [Go 1.22](https://go.dev/doc/install) 
- Docker(with compose)
  - [Install on Ubuntu](https://docs.docker.com/desktop/install/ubuntu/)
  - [Install on Windows](https://docs.docker.com/desktop/install/windows-install/)


Clone the repo:
```shell
git clone https://github.com/behnambm/data-collector.git
```

```shell
cd data-collector
```

### Makefile 

You'll need 3 separate terminal windows or tabs to run the services.

Run service 1:
```shell
make svc1
```
Run service 2:
```shell
make svc2
```
Run service 3:
```shell
make svc3
```

After running Service 3, a `database.sqlite` file will be created to store the results.


### Docker

You'll need 2 separate terminal windows or tabs to run the services.

```shell
make up
```
```shell
make svc3
```

### Notes
1. Both approaches use the default configurations found in the `configs` directory.
You can modify these configurations as needed.
2. There are two docker files, the reason for that is the network instability in Iran. One builds the binary during build time, the other one just copies the binary from host. 
3. The docker files and compose are only written for service 1 & 2.

# Run tests

```shell
make test
```


