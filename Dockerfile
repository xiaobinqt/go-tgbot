FROM golang:1.18.3 as build

ARG ARCH=x86_64
#ARG CGO_TAGS="netgo osusergo"
COPY . ${GOPATH}/src/tgbot

RUN  go env -w GO111MODULE=auto && \
     go env -w GOPROXY=https://goproxy.cn,direct

RUN cd ${GOPATH}/src/tgbot && \
    GIT_COMMIT=$(git rev-parse --short HEAD) && \
    DATE=$(date) && \
    go build -ldflags="-X 'main.gitCommit=11111' -X 'main.buildAt=2222' -w -extldflags '-static'" -v -o tgbot
##    upx tgbot

FROM busybox:stable
COPY --from=build /go/src/tgbot/tgbot /

CMD ["/tgbot"]
