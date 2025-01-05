# Ssh сервер для храниения секретов(пароли, карты, заметки)

## Общая концепция

Сервер реализует `TUI` и самое главное преимущество: отсутствие клиента приложения как такового.
При обновлении функционала сервера нет необходимости заново устанавливать клиент всем желающим.

Также одним из плюсов является то что ssh соединение - изначально защищенное.

Для хранения содержимого секретов был выбран [minio](https://min.io/), шифрование в котором уже есть из коробки.

## Настройка

Настройка осуществляется через файл конфигурации в формате `yaml`

```yaml
environment: development # среда исполнения для прода / разработки
server_address: 127.0.0.1:4040 # адрес ssh сервера
database_dsn: postgres://postgres:postgres@127.0.0.1:5432/praktikum?sslmode=disable # строка подключения к базе
server_cert_path: /home/mikhail/Learning/Go/certs/server # путь к сертификатам сервера
minio_bucket_name: secrets # имя бакета minio
minio_endpoint: localhost:9000 # адрес / endpoint minio
minio_access_key: minioadmin # access key minio
minio_secret_key: minioadmin # secret key minio
minio_use_ssl: false # ssl соединение minio
```

Путь к файлу конфига задается при запуске сервера флагом `-c`

## Запуск проекта

Поднять контейнеры postgres и minio

```shell
$> docker compose up
```

Добавить бакет minio (например через webui)

```text
http://127.0.0.1:9001/buckets
```

Создать и заполнить файл конфигурации из примера `config-server.example.yaml`

Скомпилировать go проект

```shell
$> go build -o ./cmd/server/keeper ./cmd/server/*.go
```

Запустить сервер

```shell
$> ./cmd/server/keeper -c ПУТЬ_К_ФАЙЛУ_КОНФИГУРАЦИИ
```

Profit!

## Подключение к TUI по ssh

```shell
$> ssh АДРЕС_СЕРВЕРА -p ПОРТ_СЕРВЕРА
```

Profit!

## Планы на будущее

1. Поддержка файлов в секретах (через [wish.scp](https://github.com/charmbracelet/wish/tree/main/scp))