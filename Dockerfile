# 第一阶段：构建应用
FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux

RUN go build -o ./iviuser ./cmd/apiserver && chmod +x ./iviuser

RUN ls -l /app

FROM alpine:latest  

RUN apk add --no-cache ca-certificates

WORKDIR /root

RUN mkdir -p conf/cert

COPY --from=builder /app/iviuser ./iviuser
COPY --from=builder /app/conf/cert/cert.crt ./conf/cert/cert.crt
COPY --from=builder /app/conf/cert/key.crt ./conf/cert/key.crt

RUN ls -l /root

ENV IVIUSER_MYSQL_HOSTNAME="localhost"

EXPOSE 8080
EXPOSE 8443

CMD ["./iviuser", "apiserver"]