# Етап 1: Будуємо Go-додаток
FROM golang:1.23-alpine AS builder

# Встановлюємо робочу директорію всередині контейнера
WORKDIR /usr/local/src
# Копіюємо решту коду додатку
COPY . .
# Встановлюємо необхідні пакети
RUN apk --no-cache add bash git make gcc musl-dev curl \
    && curl -fsSL https://get.docker.com -o get-docker.sh \
    && sh get-docker.sh \
    && rm get-docker.sh \
    && apk --no-cache add docker-compose

Run docker-compose build

Run	docker-compose up -d


# Виставляємо порти
EXPOSE 8080

