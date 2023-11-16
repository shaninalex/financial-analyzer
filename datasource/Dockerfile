FROM golang:1.21.3
WORKDIR /build
COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o datasource .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
WORKDIR /root/demo_data
COPY --from=0 /build/demo_data ./demo_data
COPY --from=0 /build/datasource ./
CMD ["./datasource"]