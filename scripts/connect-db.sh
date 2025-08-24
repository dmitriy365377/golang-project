#!/bin/bash

# Скрипт для подключения к PostgreSQL базе данных
# Использование: ./scripts/connect-db.sh

echo "🔌 Подключение к PostgreSQL базе данных..."
echo "📊 База данных: chatdb"
echo "👤 Пользователь: chatuser"
echo "🔑 Пароль: chatpass"
echo "🌐 Хост: localhost:5432"
echo ""

# Проверяем, запущен ли контейнер
if ! docker ps | grep -q "golang-chat-postgres"; then
    echo "❌ Контейнер PostgreSQL не запущен!"
    echo "Запустите: make docker-up"
    exit 1
fi

echo "✅ Контейнер PostgreSQL запущен"
echo ""

# Подключаемся к базе данных
echo "🚀 Подключение к базе данных..."
docker exec -it golang-chat-postgres psql -U chatuser -d chatdb
