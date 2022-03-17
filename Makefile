build:
	docker run --name=broker-db -e POSTGRES_PASSWORD='qwerty' -p 5442:5432 -d --rm postgres && sleep 2
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5442/postgres?sslmode=disable' up && sleep 1
	(nats-streaming-server -cid prod -store file -dir store && sleep 1) &
	(go run cmd/main.go && sleep 2) &
	go run pkg/publisher/publisher.go model.json && sleep 1
	open http://127.0.0.1:8007/orders/

stop:
	docker stop broker-db
	nats-streaming-server -sl stop
	pkill main
