build:
	docker compose --progress plain build
up:
	docker compose up -d
down:
	docker compose down -v
start:
	docker compose start
stop:
	docker compose stop
logs:
	docker compose logs -f google-cloud-run-webserver