DOCKER_POSTGRES_CONTAINER_NAME = atuin-postgres-container

DB_PASSWORD = password
DB_PORT = 5200
DB_DATA = /var/lib/postgresql/data
DB_USERNAME = atuin

# https://github.com/docker-library/docs/blob/master/postgres/README.md
.PHONY: run-postgresql
run-postgresql:
	docker run --name $(DOCKER_POSTGRES_CONTAINER_NAME) \
		-v atuin-pgdata:$(DB_DATA) \
		-e PGDATA=$(DB_DATA) \
		-e POSTGRES_USER=$(DB_USERNAME) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-p $(DB_PORT):5432 -d \
		postgres:15.2-alpine3.17

.PHONY: rm-postgresql
rm-postgresql:
	docker rm $(DOCKER_POSTGRES_CONTAINER_NAME)

.PHONY: start-postgresql
start-postgresql:
	docker start $(DOCKER_POSTGRES_CONTAINER_NAME)

.PHONY: stop-postgresql
stop-postgresql:
	docker stop $(DOCKER_POSTGRES_CONTAINER_NAME)

.PHONY: pgcli
pgcli:
	pgcli postgres://$(DB_USERNAME):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_USERNAME)
