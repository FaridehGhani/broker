# broker

broker is a test project.

It's using the ZeroMQ messaging library for transferring random messages.
Random messages are published from source, listener receives every published message and stores it,
ensuring no messages are lost and destination subscribes on them.
The publish and subscribe pattern is used from zeroMQ.

## How to run 
make sure ZeroMQ is installed on your OS.

create .env file, copy from .env.sample and replace values with your environmental variables.

run 
```bash
go mod tidy 
```
```bash
docker-compose up -d
```
```bash
go run main.go
```

