# Используем официальный образ Go
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Используем минимальный образ для финального контейнера
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из builder
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8000

# Запускаем приложение
CMD ["./main"] 