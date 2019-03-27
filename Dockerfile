ARG go_version=latest
ARG alpine_version=latest

FROM golang:$go_version as builder

WORKDIR /go/src/scrumtime

COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...


FROM alpine:$alpine_version

# install certificates
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/scrumtime /bin

CMD ["scrumtime", "-v"]
