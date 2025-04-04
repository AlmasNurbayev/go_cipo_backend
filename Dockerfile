# Используем официальный образ Go
FROM golang:1.23-alpine AS builder
# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify && go mod tidy

# Копируем файлы проекта
COPY . .

# Собираем приложение
RUN go build -o SERVER ./cmd/server/main.go
RUN go build -o MIGRATOR ./cmd/migrator/main.go
RUN go build -o PARSER ./cmd/parser/main.go
RUN go build -o SEEDER ./cmd/seeder/main.go
RUN go build -o CLEARDB ./cmd/cleardb/main.go


# Используем stage 2: минимальный контейнер
FROM alpine:3.21.3 AS final
WORKDIR /app/

# Добавляем необходимые зависимости
# tzdata - для установки временной зоны
# curl - для проверки доступности сервиса
# dumb-init - для корректного запуска и завершения работы приложения в prefork-режиме
RUN apk add --no-cache tzdata curl dumb-init

# Копируем бинарники, миграции и конфиги из builder-образа
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/seeds ./seeds
COPY --from=builder /app/SERVER .
COPY --from=builder /app/MIGRATOR .
COPY --from=builder /app/PARSER .
COPY --from=builder /app/SEEDER .
COPY --from=builder /app/CLEARDB .

# dumb-init нужен для нормального запуска в prefork-режиме
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "./MIGRATOR -typeTask up -dsn $DSN && ./SEEDER -typeTask up -dsn $DSN && exec ./SERVER"]