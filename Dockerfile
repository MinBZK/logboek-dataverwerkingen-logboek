FROM golang:1.22.3-alpine AS build

WORKDIR /app/server

RUN apk add --no-cache gcc musl-dev

COPY libs/logboek-go /app/libs/logboek-go
COPY server/go.mod server/go.sum .

RUN go mod download -x

COPY server/cmd cmd
COPY server/pkg pkg

RUN go install -ldflags "-s -w" -tags sqlite_omit_load_extension -v ./cmd/...


FROM alpine:3.19.1

WORKDIR /var/lib/logboek

COPY --from=build /go/bin/ /usr/local/bin

ENV LISTEN_ADDRESS="0.0.0.0:9000"

EXPOSE 9000

ENTRYPOINT ["/usr/local/bin/logboek"]
