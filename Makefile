# @author Futa Nakayama & Riku Nunokawa

SERVER_REPOSITORY_NAME:=teamF/scenepicks
SERVER_CONTAINER_NAME:=scenepicks

DBNAME:=nexus_db
DOCKER_DNS:=db
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=password

HOST_APP_BASE:=$(shell pwd)
DOCKER_APP_BASE:=/go/src/github.com/shortintern2020-C-cryptograph/TeamF/server

# @author Futa Nakayama

local/run/server: # コンテナでdbだけ、サーバはローカルで起動（フロントはなにもしない）
	$(make) docker/run/db
	cd server && make run
	@echo 'connect server port :3000 !!!'

docker/run: # 全部コンテナで起動
	docker-compose -f ./docker-compose.yml build --no-cache
	docker-compose -f ./docker-compose.yml up -d
#	$(MAKE) docker/run/server
#	$(MAKE) docker/run/db

docker/run/server: # server, dbのみをコンテナで起動
	#docker run -d --name $(SERVER_CONTAINER_NAME) -p 1323:1323 -v $(HOST_APP_BASE):$(DOCKER_APP_BASE) $(SERVER_REPOSITORY_NAME):latest
	docker-compose -f ./docker-compose.yml build --no-cache server db
	docker-compose -f ./docker-compose.yml up -d server db
	@echo 'connect server port :8080 !!!'

docker/run/db: # dbのみをコンテナで起動
	docker-compose -f ./docker-compose.yml up -d db


docker/stop: # コンテナ全部落として、削除する
	docker-compose down
	#docker container rm $(docker container ls -a -q)
	#docker image rm $(docker image ls -q)
	#$(MAKE) docker/stop/server
	#docker container rm $(SERVER_CONTAINER_NAME)

docker/stop/server: # コンテナ全部落として、削除する
	docker-compose down

local/run/frontend: # フロントエンドを起動！
	$(make) docker/run/server
	cd app && yarn && yarn run dev
	@echo 'frontend served at port 3000 !'

local/stop: # ローカルで動かしてたやつ以外のコンテナを削除（only db or db + server）
	docker-compose down

# @author Riku Nunokawa

DB_SERVICE:=db
mysql/client:
	docker-compose exec $(DB_SERVICE) mysql -uroot -hlocalhost -ppassword $(DBNAME)

mysql/init:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -ppassword \
		-e "create database \`$(DBNAME)\`"

__mysql/drop:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -ppassword \
		-e "drop database \`$(DBNAME)\`"


MIGRATION_SERVICE:=migration
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

flyway/validate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) validate

flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline

flyway/clean:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) clean

