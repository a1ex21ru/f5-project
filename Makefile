migrateUp:
	goose -dir db/migrations postgres "postgresql://user:admin@127.0.0.1:5432/practice?sslmode=disable" up

removeAll:
	docker stop $(docker ps -aq) 2>/dev/null
	docker rm -f $(docker ps -aq) 2>/dev/null
	docker rmi -f $(docker images -aq) 2>/dev/null
	docker volume prune -f