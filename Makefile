help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'make dev: make dev for development work'
	@echo 'make build: make build container'
	@echo 'make production: docker production build'
	@echo 'clean: clean for all clear docker images'

dev:
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker-compose -f docker-compose.yml down
	docker-compose -f docker-compose.yml up

build:
	docker-compose -f docker-compose.yml down build

clean:
	docker-compose -f docker-compose.yml down -v
