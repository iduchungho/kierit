# Sử dụng image golang:1.18
FROM golang:1.18

# Thiết lập biến môi trường
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Sao chép mã nguồn ứng dụng vào image
WORKDIR /app
COPY . .

# Build ứng dụng
RUN go mod download
RUN go build cmd/main.go

# Chạy ứng dụng
CMD ["./main"]