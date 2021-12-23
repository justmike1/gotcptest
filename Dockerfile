FROM golang:1.17 AS builder

WORKDIR /app/
ADD go.* /app/
RUN go mod download

ADD main.go /app/
ENV CGO_ENABLED=0
RUN go build -o /gotcptest main.go

FROM scratch
COPY --from=builder /gotcptest /gotcptest
EXPOSE 1337
CMD ["/gotcptest"]

