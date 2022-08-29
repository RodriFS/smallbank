DB:="$$(docker ps -f name=db -q)"
rundb:
	docker-compose up -d

killdb:
	docker-compose down

createdb:
	docker exec -it $(DB) createdb --username=admin --owner=admin small_bank

createtestdb:
	docker exec -it $(DB) createdb --username=admin --owner=admin testing

dropdb:
	docker exec -it $(DB) dropdb -U admin small_bank

droptestdb:
	docker exec -it $(DB) dropdb -U admin testing

accessdb:
	docker exec -it $(DB) psql -U admin smal_bank

logdb:
	docker logs $(DB)

compile:
	CompileDaemon -command="./server"

test:
	go test -v ./...

.PHONY: rundb killdb createdb createtestdb dropdb droptestdb accessdb logdb compile test
