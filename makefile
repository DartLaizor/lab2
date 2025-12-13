.PHONY: dev

dev:
	@echo "Запуск Docker"
	docker desktop start && docker compose up -d

	@echo "Ожидание готовности БД..."
	@until docker compose exec -T postgres pg_isready -U postgres 2>/dev/null; do \
		echo -n "."; \
		sleep 2; \
	done; \
	echo " БД готова!"

	@echo "Запуск Go"
	konsole -e bash -c 'cd ./back && CONFIG_PATH="./config/config.yaml" ./cmd/server; echo "Сервер завершил работу. Нажмите Enter..."; read' &

	@echo "Запуск vue"
	cd front/reviews-web && npm run dev

go:
	CONFIG_PATH=".back/config/config.yaml" go run .back/cmd/.
docker:
	docker desktop start;
	docker-compose up -d
web:
	cd front/reviews-web;
	npm run dev
