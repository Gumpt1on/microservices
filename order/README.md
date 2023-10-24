sudo docker run -p 3306:3306	\
	-e MYSQL_ROOT_PASSWORD=abc123456 \
	-e MYSQL_DATABASE=order mysql

DATA_SOURCE_URL="root:abc123456@tcp(127.0.0.1:3306)/order?parseTime=true" \
APPLICATION_PORT=3000 \
ENV=development \
go run cmd/main.go

grpcurl -d \
'{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' \
-plaintext localhost:3000 \
Order/Create

grpcurl -d \
'{"order_id": 1}' \
-plaintext localhost:3000 \
Order/Get