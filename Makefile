build:
	docker build -t apigwsample/auth ./auth-service
	docker build -t apigwsample/product ./product-service

run:
	docker compose up -d
	docker compose ps
	docker compose top