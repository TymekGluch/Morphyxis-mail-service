FROM golang:1.26.3-alpine AS builder

WORKDIR /src
COPY . .

RUN apk add --no-cache make bash
RUN make build
RUN chmod +x bin/morphixis-mail-service

CMD ["./bin/morphixis-mail-service"]
