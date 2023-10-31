## payment 启动命令：
DB_DRIVER=mysql	\
DATA_SOURCE_URL="abc:abc123456@tcp(192.168.85.131:3306)/payment?parseTime=true" \
APPLICATION_PORT=3001 \
ENV=development \
go run cmd/main.go

## order 启动命令：
DB_DRIVER=mysql	\
DATA_SOURCE_URL="abc:abc123456@tcp(192.168.85.131:3306)/order?parseTime=true" \
APPLICATION_PORT=3000 \
ENV=development \
PAYMENT_SERVICE_URL=localhost:3001 \
go run cmd/main.go