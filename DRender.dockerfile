# Базовий образ для збірки
FROM golang:1.23.0-alpine AS builder

# Встановлюємо робочу директорію
WORKDIR /app

# Копіюємо файли модулів та завантажуємо залежності
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо весь код
COPY . .

# Збираємо додаток
RUN go build -o app ./cmd/app/main.go

# Фінальний образ
FROM alpine:latest

# Встановлюємо робочу директорію
WORKDIR /app

# Копіюємо зібраний додаток
COPY --from=builder /app/app .

# Копіюємо необхідні файли (якщо потрібно)
COPY templates/ ./templates/
COPY config/ ./config/

# Встановлюємо змінні середовища
ENV CONFIG_PATH=/app/config/prod.yaml

# Відкриваємо порт
EXPOSE 8080

# Команда для запуску додатка
CMD ["./app"]
