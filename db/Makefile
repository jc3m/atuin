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
