## Генерация proto-файла

```
~/notify-platform/services/notification-command-service$

protoc \
--go_out=paths=source_relative:. \
--go-grpc_out=paths=source_relative:. \
proto/*.proto
```

## В проекте используется podman-compose
Запуск:
```
~/notify-platform/deployments/docker$ podman-compose up -d

podman ps

```

Создание топика в kafka:
```
podman exec -it notify_kafka \
/opt/kafka/bin/kafka-topics.sh \
--create \
--topic notifications.created \
--bootstrap-server localhost:9092 \
--partitions 3 \
--replication-factor 1
```

Проверка что топик появился:
```
$ podman exec -it notify_kafka \
/opt/kafka/bin/kafka-topics.sh \
--list \
--bootstrap-server localhost:9092
notifications.created
```