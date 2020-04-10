FROM golang:1.14 as builder
COPY . /go/src/github.com/radityaarya/learn-kubernetes/1_kubernetes_ready_microservice

WORKDIR /go/src/github.com/radityaarya/learn-kubernetes/1_kubernetes_ready_microservice
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.9
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/radityaarya/learn-kubernetes/1_kubernetes_ready_microservice/ .
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

EXPOSE 3000
CMD /wait && ./app
