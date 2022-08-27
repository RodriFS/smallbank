DB:="$$(docker ps -f name=db -q)"
rundb:
	docker-compose up -d

killdb:
	docker-compose down

createdb:
	docker exec -it $(DB) createdb --username=admin --owner=admin small_bank

dropdb:
	docker exec -it $(DB) dropdb -U admin small_bank

accessdb:
	docker exec -it $(DB) psql -U admin smal_bank

logdb:
	docker logs $(DB)

compile:
	CompileDaemon -command="./server"

.PHONY: rundb killdb createdb dropdb accessdb logdb compile
