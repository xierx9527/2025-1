FROM golang:1.24-alpine AS builder

WORKDIR /app

# 设置国内Go模块代理
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o book-management .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/book-management .
EXPOSE 8080
ENV MYSQL_DSN="root:20041017@tcp(mysql:3306)/Book_Management?charset=utf8mb4&parseTime=True&loc=Local"
CMD ["./book-management"]