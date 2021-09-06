FROM rust:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache -U musl-dev
RUN cargo build --release

FROM alpine
COPY --from=builder /app/target/release/auto-api /auto-api
WORKDIR /
ENTRYPOINT ["/auto-api"]
CMD [""]
