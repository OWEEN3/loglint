#!/bin/bash

# Цвета для красоты
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${GREEN}===> 1. Проверка инструментов...${NC}"
if ! command -v golangci-lint &> /dev/null; then
    echo -e "${RED}Ошибка: golangci-lint не найден. Установите его перед запуском.${NC}"
    exit 1
fi

RAW_URL="https://raw.githubusercontent.com/OWEEN3/loglint/main"

echo -e "${GREEN}===> 2. Скачивание конфигурационных файлов...${NC}"
curl -sSfL "$RAW_URL/.custom-gcl.example.yml" -o ".custom-gcl.yml"
curl -sSfL "$RAW_URL/.golangci.example.yml" -o ".golangci.yml"
echo -e "${YELLOW}Файлы .custom-gcl.yml и .golangci.yml созданы.${NC}"

echo -e "${GREEN}===> 3. Сборка кастомного бинарника (это может занять минуту)...${NC}"
rm -f custom-gcl
if golangci-lint custom -v; then
    echo -e "${GREEN}Сборка завершена успешно!${NC}"
else
    echo -e "${RED}Ошибка при сборке кастомного линтера.${NC}"
    exit 1
fi

echo -e "${GREEN}===> 4. Запуск линтера loglint...${NC}"
./custom-gcl run

echo -e "\n${YELLOW}Проверка завершена.${NC}"
read -p "Хотите удалить временные файлы (.custom-gcl.yml, .golangci.yml, custom-gcl)? [y/N] " confirm
if [[ "$confirm" =~ ^[Yy]$ ]]; then
    rm -f .custom-gcl.yml .golangci.yml custom-gcl
    echo -e "${GREEN}Временные файлы удалены.${NC}"
fi
