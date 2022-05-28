BUF_VERSION:=1.1.0

generate:
	@buf generate
	@wire

up:
	@docker-compose up -d

down:
	@docker-compose down

dev:
	@SERVE_HTTP=true air

test:
	@./tests.sh

wire:
	@wire ./...
