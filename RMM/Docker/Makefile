# Docker compose

MIGRATIONS_DIR=$(GOPATH)/src/github.com/ContinuumLLC/rmm-dynamicgroup-service/cql/migrations
MIGRATE="./Cassandra/migrate"
DB_URL="cassandra://localhost:9042/dynamic_groups_space?consistency=ONE"

.PHONY: clear
clear:
	@echo ===== Stop and remove containers
	@docker-compose stop
	@docker system prune -f

.PHONY: start
start:
	@echo ===== Start docker compose
	@docker-compose build && docker-compose up -d

.PHONY: up
up: clear start migrate show

.PHONY: show
show:
	@docker container ls --format "{{.Names}} [{{.Ports}}]"

.PHONY: migrate
migrate:
	@echo ===== Wait until Cassandra runs
	@until cqlsh --cqlversion 3.4.4 -e "show version" 2> /dev/null ; do sleep 1 ; done
	@echo ===== Make migrations
	@cat ./Cassandra/cql/schema.cql | cqlsh --cqlversion 3.4.4
	@$(MIGRATE) -verbose -database $(DB_URL) -path $(MIGRATIONS_DIR) up
	@echo ===== Done