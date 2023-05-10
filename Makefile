start-mongo:
	docker run --name mongo_db -d --rm -e ME_CONFIG_MONGODB_SERVER=some-mongo -p 27017:27017 mongo:latest


stop-mongo:
	docker stop mongo_db


run:start-mongo
	go run cmd/server/server.go
