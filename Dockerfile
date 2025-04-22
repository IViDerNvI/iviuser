# 第一阶段：构建应用
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux

RUN go build -o ./iviuser ./cmd/apiserver && chmod +x ./iviuser

FROM alpine:latest  

RUN apk add --no-cache ca-certificates

WORKDIR /root

RUN mkdir -p conf/cert

COPY --from=builder /app/iviuser ./iviuser
COPY --from=builder /app/conf/cert/cert.crt ./conf/cert/cert.crt
COPY --from=builder /app/conf/cert/key.crt ./conf/cert/key.crt

ENV IVIUSER_MYSQL_HOSTNAME="localhost"
ENV IVIUSER_MINIO_ENDPOINT="localhost:9000"
ENV IVIUSER_JUDGE_RPC_ENDPOINT="localhost:50052"
ENV IVIUSER_REDIS_HOSTNAME="localhost:6379"
ENV IVIUSER_JUDGE_RPC_ENDPOINT="localhost:50051"

EXPOSE 8080
EXPOSE 8443

CMD ["./iviuser", "apiserver"]