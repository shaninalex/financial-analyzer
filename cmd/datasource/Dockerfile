FROM golang:1.22.0
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o datasource ./cmd/datasource

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /build/datasource ./
CMD ["./datasource"]
