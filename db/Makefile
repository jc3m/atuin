# TODO: Share DB connection variables between Makefiles

GOOSE_DBSTRING = "user=atuin password=password dbname=atuin host=localhost port=5200 sslmode=disable"

.PHONY: goose-status
goose-status:
	cd migrations; \
	goose postgres $(GOOSE_DBSTRING) status

.PHONY: goose-up
goose-up:
	cd migrations; \
	goose postgres $(GOOSE_DBSTRING) up


.PHONY: generate
generate:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate


.PHONY: dump
dump:
	pg_dump --host localhost --port 5200 --username atuin --format plain --file dump.sql --table tubul.events atuin
