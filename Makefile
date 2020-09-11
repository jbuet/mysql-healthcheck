.PHONY: build
build:
	@go build

.PHONY: docker-up
docker-up:
	@docker-compose up -d

.PHONY: docker-down
docker-down:
	@docker-compose down

.PHONY: run
run:
	@sudo MYSQL_HEALTHCHECK_PATH=".mysql_healthcheck.conf" ./mysql-healthcheck
 
	


