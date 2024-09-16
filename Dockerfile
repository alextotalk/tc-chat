FROM golang:1.23.0-alpine AS builder

WORKDIR /usr/local/src

# Установка необходимых пакетов
RUN apk --no-cache add bash git make gcc gettext musl-dev

# Копируем зависимости Go и загружаем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходные коды проекта

# Копіюємо файл конфігурації до очікуваного шляху
COPY config/local.yaml /usr/local/src/local.yaml
COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/
COPY templates/ templates/

# Сборка Go-приложения
RUN go build -o ./bin/app ./cmd/app/main.go

# Финальный этап для создания минимального образа
FROM alpine AS runner

# Копируем бинарный файл из фазы сборки
COPY --from=builder /usr/local/src/bin/app /app

# Копируем конфигурацию и шаблоны
COPY config/local.yaml /usr/local/src/
COPY templates/ /usr/local/src/templates/

# Указываем рабочую директорию
WORKDIR /usr/local/src

# Указываем команду для запуска контейнера
CMD ["/app"]
