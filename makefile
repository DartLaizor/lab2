.PHONY: dev

dev:
	@echo "Запуск Docker..."
	docker desktop start && docker compose up -d

	@echo "Запуск Go backend в фоне..."
	konsole -e bash -c 'cd ./back && CONFIG_PATH="./config/config.yaml" go run ./cmd/; echo "Сервер завершил работу. Нажмите Enter..."; read'&

	@echo "Запуск фронтенда..."
	cd front/reviews-web && npm run dev

go:
	CONFIG_PATH=".back/config/config.yaml" go run .back/cmd/. 
docker:
	docker desktop start;
	docker-compose up -d
web:
	cd front/reviews-web;
	npm run dev


