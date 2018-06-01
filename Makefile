flags := -p test

.PHONY: test
test:
	docker-compose $(flags) build
	docker-compose $(flags) up --force-recreate --remove-orphans -d
	docker wait test_client-example_1 > test-return-value
	docker-compose $(flags) stop
	docker-compose $(flags) logs -f
	docker-compose $(flags) rm -v --force
	exit `cat test-return-value`

.PHONY: kill
kill:
	docker-compose $(flags) kill
	docker-compose $(flags) rm -v --force

.PHONY: clean
clean:
	- rm -f test-return-value
