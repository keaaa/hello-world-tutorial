FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o main .

FROM scratch as runner
COPY --from=builder /build/main /

ENTRYPOINT [ "/main" ]