FROM golang:1.13-alpine as builder
RUN apk add git
COPY . /go/src/shuhelperConfigStorage
ENV GO111MODULE on
WORKDIR /go/src/shuhelperConfigStorage
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/shuhelperConfigStorage/shuhelperConfigStorage /
WORKDIR /
CMD ./shuhelperConfigStorage
ENV PORT 8000
EXPOSE 8000