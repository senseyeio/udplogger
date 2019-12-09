# the name of the command to build
ARG name=udplogger

####
# Creates a Docker image with the sources.
FROM golang:1.13.5-alpine3.10 AS build
ARG name

WORKDIR /tmp/src

# Copy the sources.
COPY . .

ENV CGO_ENABLED 0

RUN go build \
    -o /bin/healthcheck \
    -ldflags '-extldflags "-static"' \
    ./cmd/healthcheck

RUN go build \
    -o /bin/${name} \
    -ldflags '-extldflags "-static"' \
    ./cmd/${name}

####
# Creates a single layer image to run the service.
FROM scratch as run
ARG name

COPY --from=build /bin/healthcheck /bin/healthcheck
COPY --from=build /bin/${name} /bin/entrypoint
ENTRYPOINT ["/bin/entrypoint"]
