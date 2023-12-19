createmigration:
	docker run -v .:/migrations --network host migrate/migrate create -ext sql -dir /migrations/sql/migrations -seq init
migrate:
	docker run -v .:/migrations --network host migrate/migrate -database "mysql://root:root@tcp(localhost:3306)/orders" -path /migrations/sql/migrations -verbose up

migradedown:
	docker run -v .:/migrations --network host migrate/migrate -database "mysql://root:root@tcp(localhost:3306)/orders" -path /migrations/sql/migrations -verbose down

.PHONY: migrate migradedown createmigration