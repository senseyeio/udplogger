# udplogger

A command that echoes UDP payloads to the stdout.  It is specially
designed to be used as a debugging service in docker-compose.

A PORT environment variable is required, with the port to listen to.

Includes an internal healthcheck and can be gracefully stopped by a
SIGINT or a SIGTERM.

## Installation

```
bash; go get github.com/senseyeio/udplogger/cmd/udplogger
```

## Example of use

This is how to use udplogger to emulate a statsd service:
```
version: '2.1'

services:
  statsd:
    image: senseyeio/udplogger:0.0.1
    environment:
      - PORT=8125

  tester:
    image: subfuzion/netcat
    depends_on:
      statsd:
        condition: service_healthy
    entrypoint: "/bin/ash"
    command: "-c 'nc -u -w1 statsd 8125 < /etc/alpine-release'"
```

You can also run it as a single docker container:
```
bash; docker run --env PORT=12345 -p'0.0.0.0:12345:12345/udp' senseyeio/udplogger:0.0.1
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

Or as a command:
```
bash; PORT=12345 udplogger
2018/02/16 23:40:10 listening on 0.0.0.0:12345
```

## TODO

- Upload the image to Docker Hub.
- Set `SO_REUSEADDR`.
