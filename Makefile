# can pass migration file name as name=file_name
migrate_init:
	@goose -dir db/schema postgres ${DATABASE_URL} create $(name) sql

migrate_up:
	@goose -dir db/schema postgres ${DATABASE_URL} up

migrate_down:
	@goose -dir db/schema postgres ${DATABASE_URL} down