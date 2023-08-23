# nats-server-test

## Launching
```shell
dokcer compose up --build
```

### requests
```shell
#subscribe to time ticker
nats sub time.tick -s nats://localhost:4222

#send request
nats req request.echo "Hello, NATS" -s nats://localhost:4222
```
