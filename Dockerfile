FROM golang:1.22.3-alpine AS build

WORKDIR /app/server

RUN apk add --no-cache gcc musl-dev

COPY server/go.mod server/go.sum .

# TEMP
COPY libs/logboek-go /app/libs/logboek-go
RUN go mod edit -replace github.com/MinBZK/logboek-dataverwerkingen-logboek/libs/logboek-go=/app/libs/logboek-go
# TEMP END

RUN go mod download -x

COPY server/cmd cmd
COPY server/pkg pkg

RUN go install -ldflags "-s -w" -tags sqlite_omit_load_extension -v ./cmd/...


FROM alpine:3.19.1

WORKDIR /var/lib/logboek

COPY --from=build /go/bin/ /usr/local/bin

ENTRYPOINT ["/usr/local/bin/logboek"]
