<h1>  <a href="https://t.me/my_cheer_up_bot">Телеграм бот </a> поднимающий настроение </h1>

Бот отправляет комплименты или мотивационные сообщения в определенные промежутки времени (по умолчанию с 11 до 21:00 каждые 2 часа, задается в конфиге)
Так же отправляет при отправке любого сообщения в него

```
    docker-compose up -d postgres
```

Запуски миграций
```
    goose -dir db/migrations postgres "postgresql://ainur:secret@127.0.0.1:5432/mood_lifter?sslmode=disable" up
```

Пример простого запуска

```
    docker build -t mood-lifter:v1 .
    docker run -it --rm --network=host -e DB_HOST="127.0.0.1" -e DB_USERNAME=ainur -e BOT_TOKEN=6781544832:  mood-lifter:v1
```