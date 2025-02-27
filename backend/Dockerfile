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


# Используем stage 2: минимальный контейнер
FROM alpine:latest
WORKDIR /root/

# Копируем бинарники, миграции и конфиги из builder-образа
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/SERVER .
COPY --from=builder /app/MIGRATOR .
COPY --from=builder /app/PARSER .
COPY --from=builder /app/SEEDER .

# Запускаем приложение
#CMD sh -c './MIGRATOR -typeTask up -dsn $DSN && ./SERVER'
#CMD sh -c './SERVER'
ENTRYPOINT ["sh", "-c", "./MIGRATOR -typeTask up -dsn $DSN && exec ./SERVER"]