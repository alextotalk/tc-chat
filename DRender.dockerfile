# Етап 1: Будуємо Go-додаток
FROM golang:1.20-alpine AS builder

# Встановлюємо робочу директорію всередині контейнера
WORKDIR /usr/local/src

# Встановлюємо необхідні пакети
RUN apk --no-cache add bash git make gcc musl-dev

# Копіюємо файли go.mod та go.sum та завантажуємо залежності
COPY go.mod go.sum ./
RUN go mod download

# Копіюємо решту коду додатку
COPY . .

# Будуємо Go-додаток
RUN go build -o ./bin/app ./cmd/app/main.go

# Етап 2: Створюємо фінальний образ
FROM alpine:latest

# Встановлюємо змінні середовища для PostgreSQL та додатку
ENV CONFIG_PATH=/usr/local/src/local.yaml \
    POSTGRES_DB=postgresql \
    POSTGRES_USER=alex \
    POSTGRES_PASSWORD=secret \
    PGDATA=/var/lib/postgresql/data

# Виставляємо порти
EXPOSE 8080 5432

# Встановлюємо необхідні пакети
RUN apk --no-cache add bash postgresql postgresql-contrib

# Створюємо необхідні директорії та встановлюємо права доступу
RUN mkdir -p /run/postgresql && chown -R postgres:postgres /run/postgresql /var/lib/postgresql

# Встановлюємо робочу директорію
WORKDIR /usr/local/src

# Копіюємо скомпільований Go-додаток з попереднього етапу
COPY --from=builder /usr/local/src/bin/app /app

# Копіюємо конфігураційні файли та шаблони
COPY config/ config/
COPY templates/ templates/

# Копіюємо скрипт start.sh та надаємо права на виконання
COPY start.sh /start.sh
RUN chmod +x /start.sh

# Вказуємо команду запуску
CMD ["/start.sh"]
