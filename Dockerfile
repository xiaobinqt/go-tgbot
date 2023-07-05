FROM golang:1.18.3 as build

ARG ARCH=x86_64
#ARG CGO_TAGS="netgo osusergo"
COPY . ${GOPATH}/src/tgbot

RUN  go env -w GO111MODULE=auto && \
     go env -w GOPROXY=https://goproxy.cn,direct

RUN cd ${GOPATH}/src/tgbot && \
    GIT_COMMIT=$(git rev-parse --short HEAD) && \
    DATE=$(date) && \
    go build -ldflags="-X 'main.gitCommit=$GIT_COMMIT' -X 'main.buildAt=$DATE' -w -extldflags '-static'" -v -o tgbot

FROM debian:buster-slim

RUN apt update && \
    apt install -y ca-certificates

ENV TZ=Asia/Shanghai

COPY --from=build /go/src/tgbot/tgbot /usr/bin/

CMD ["/usr/bin/tgbot"]
