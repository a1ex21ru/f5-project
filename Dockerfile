
# Используем официальный образ Go с конкретной версией для стабильности сборки
FROM golang:1.25-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем инструменты и сертификаты в builder
RUN apk add --no-cache git ca-certificates

# Копируем файлы go.mod и go.sum отдельно для кеширования зависимостей
COPY go.mod go.sum ./

RUN go mod download

# Копируем весь исходный код приложения
COPY . .



# ⚠️ ВАЖНО: Собираем статически слинкованный бинарник для Alpine
 #CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
RUN go build -ldflags="-s -w" -o f5-project ./cmd/app

# ======= Этап 2: Минимальный образ для запуска приложения =======

# Используем минимальный стабильный образ Alpine Linux для финального контейнера
FROM alpine:3.21.3

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

RUN apk add --no-cache ca-certificates

# Создаём непривилегированного пользователя и группу с фиксированными UID/GID
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Копируем готовый бинарник из стадии сборки
COPY --from=builder /app/f5-project .
COPY --from=builder /app/static ./static

# Меняем владельца файлов на созданного пользователя
RUN chown -R appuser:appgroup /app

# Переключаемся на непривилегированного пользователя
USER appuser

# Устанавливаем команду по умолчанию для запуска контейнера
CMD ["./f5-project"]