FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN \
    apt-get -qq update  &&\
    apt-get -yqq install upx

WORKDIR /src
COPY . .

RUN \
  go build \
    -ldflags "-s -w -extldflags '-static'" \
    -o /bin/action .  &&\
  strip /bin/action  &&\
  upx -q -9 /bin/action




FROM scratch

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/passwd
COPY --from=builder --chown=65534:0 /bin/action /action

USER nobody
ENTRYPOINT ["/action"]
