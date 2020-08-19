FROM golang:1.15-alpine as builder
RUN apk add make
COPY . /go/src/autoAPI
WORKDIR /go/src/autoAPI
RUN make build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/autoAPI/autoAPI /
WORKDIR /
ENTRYPOINT ["/autoAPI"]
CMD ["--help"]
