DB_URL=mysql://root:password@tcp(192.168.1.110:3306)/goginapi


migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migrateforce:
	migrate -path db/migration -database "$(DB_URL)" force $(versao)

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

dev:
	air

.PHONY: migrate migrateup1 migratedown migratedown1 new_migration dev sqlc