# udplogger

A command that echoes UDP payloads to the stdout.
It is specially designed to be used as a docker-compose service
to debug services that send UDP messages.
You will find Docker images for this project in [DockerHub](https://hub.docker.com/r/senseyeio/udplogger).

A PORT environment variable is required, with the port to listen to.
The process can be gracefully stopped with SIGTERM or SIGINT.

This repository includes an additional `healthcheck` command that can be used
to monitor the health of `udplogger` in docker or docker-compose.

## Installation

```
bash; go get github.com/senseyeio/udplogger/cmd/udplogger
```

## Example of use

```
bash; PORT=12345 udplogger
2018-06-01T14:17:42.102Z udplogger started
2018-06-01T14:17:42.102Z listening on 0.0.0.0:12345
```

## Example of use (docker)
```
bash; docker run --env PORT=12345 -p'0.0.0.0:12345:12345/udp' senseyeio/udplogger:0.0.2
2018-06-01T14:17:44.023Z udplogger started
2018-06-01T14:17:44.023Z listening on 0.0.0.0:12345
```

## Example of use (docker-compose)

```
bash; cat docker-compose.yml
version: '2.1'

services:
  udplogger:
    build: .
    environment:
      - PORT=12345
    healthcheck:
      test: ["CMD", "/bin/healthcheck"]
      interval: 250ms
      timeout: 250ms
      retries: 7

  client-example:
    image: subfuzion/netcat
    depends_on:
      udplogger:
        condition: service_healthy
    entrypoint: "/bin/ash"
    command: "-c 'echo -n hi! | nc -u -w1 udplogger 12345'"
bash;
bash;
bash; docker-compose build && docker-compose up
[...]
Successfully tagged udplogger_udplogger:latest
client-example uses an image, skipping
Starting udplogger_udplogger_1 ... 
Starting udplogger_udplogger_1 ... done
Recreating udplogger_client-example_1 ... 
Recreating udplogger_client-example_1 ... done
Attaching to udplogger_udplogger_1, udplogger_client-example_1
udplogger_1         | 2018-06-01T14:17:46.408Z udplogger started
udplogger_1         | 2018-06-01T14:17:46.408Z listening on 0.0.0.0:12345
udplogger_1         | 2018-06-01T14:17:47.307Z received: "hi!"
udplogger_client-example_1 exited with code 0
^CGracefully stopping... (press Ctrl+C again to force)
Stopping udplogger_udplogger_1 ... done
```

