# README
## Описание
Проект: API-booking. Реализован с использованием паттерна чистой архитектуры, с использованием 

* HTTP API:
- [http://localhost:8080]

## Сборка и запуск приложения
```bash
make start
```
## End-points
- POST: /room/create - создание комнаты
```sh
curl --location 'http://localhost:8080/room/create' \
--header 'Content-Type: application/json' \
--data '{
    "hotel_id": "redisson",
    "room_id": "k25",
    "category": "luxury",
    "description": "luxuryies",
    "status": "new"
}'
```
- POST: /order/create - создание брони
```sh
curl --location 'http://localhost:8080/order/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hotel_id": "redisson",
    "room_id": "k22",
    "email": "example@email.com",
    "from": "2024-01-11T00:00:00Z",
    "to": "2024-01-15T00:00:00Z"
}
'
```
- GET /available/day?date={date}: Получение списка доступных комнат на указанную дату
```sh
curl --location 'http://localhost:8080/available/day?date=2024-01-15'
```
- GET /available/interval?from_date={from_date}&to_date={to_date}: Получение списка доступных комнат в указанном диапазоне дат
```sh
curl --location 'http://localhost:8080/available/interval?from_date=2024-04-20&to_date=2024-04-25' \
--header 'Content-Type: application/json'
```
Для запуска нужно создать `TOOLS_ENV` и запустить сборку через make start
```
DEPLOY_CONF=deploy/config.yml
CONFIG_SOURCE=local
```