FROM golang:1.21-alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o skychart main.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/skychart /bin/skychart

ENTRYPOINT ["/bin/skychart"]
CMD ["cosmos/chain-registry", ":8090"]
