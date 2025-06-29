# Решение задачи Level 0

Здесь представлено решение задачи описанной [тут](./TASK.md)

Решение состоит из окружения и двух приложений:
- [Container] Postgres (База инициализируется на страте)
- [Container] Zookeeper
- [Container] Kafka
- [Container] Kafka Init - создаёт топик
- [Container] Kafka UI
- [APP] Producer - Публикует сообщения в топик. Раз в 1 секунду
- [APP] Web And Consumer - HTTP сервер + перекладчик сообщений в Postgres



## Порядок запуска

> ⚠️ Все команды запускать из директории `level_0`


Если требуется поменять порты, то нужно:
1. Поправить [Docker Compose](./docker-compose.yml)
2. Поправить [ENV конфиги](./configs/)


### Запуск окружения
```shell
docker compose up -d
```

Kafka UI будет доступна по [ссылке](http://localhost:20002)


### Запуск Producer-а
```shell
set -o allexport
source ./configs/producer.env
set +o allexport
printenv | grep KAFKA
cd level_0/apps/cmd/producer
go run main.go
```


### Запуск Web и Consumer
```shell
set -o allexport
source ./configs/web_and_consumer.env
set +o allexport
cd level_0/apps/cmd/web_and_consumer
printenv | grep -E "KAFKA|POSTGRES"
go run main.go
```

HTTP сервер будет доступен по [ссылке](http://localhost:3000)


### Остановка окружения
```shell
docker compose up -d
```
