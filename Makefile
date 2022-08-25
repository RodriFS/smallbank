DB:="$$(docker ps -f name=db -q)"
rundb:
	docker-compose up -d

createdb:
	docker exec -it $(DB) createdb --username=admin --owner=admin small_bank

dropdb:
	docker exec -it $(DB) dropdb -U admin small_bank

migratedbup:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/small_bank?sslmode=disable" --verbose up

migratedbdown:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/small_bank?sslmode=disable" --verbose down

accessdb:
	docker exec -it $(DB) psql -U admin smal_bank

logdb:
	docker logs $(DB)

compile:
	CompileDaemon -command="./main"

.PHONY: rundb createdb dropdb migratedbup migratedbdown accessdb logdb compile
