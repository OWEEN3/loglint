#!/bin/bash

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

if [[ "$1" == "--clean" ]]; then
    echo -e "${YELLOW}Удаление временных файлов...${NC}"
    rm -f .custom-gcl.yml .golangci.yml custom-gcl
    echo -e "${GREEN}Временные файлы удалены.${NC}"
    exit 0
fi

echo -e "${GREEN}===> 1. Проверка инструментов...${NC}"
if ! command -v golangci-lint &> /dev/null; then
    echo -e "${RED}Ошибка: golangci-lint не найден. Установите его перед запуском.${NC}"
    exit 1
fi

RAW_URL="https://raw.githubusercontent.com/OWEEN3/loglint/main"

echo -e "${GREEN}===> 2. Проверка конфигурационных файлов...${NC}"

if [[ ! -f ".custom-gcl.yml" ]] || [[ ! -f ".golangci.yml" ]]; then
    echo -e "${YELLOW}Файлы конфигурации не найдены. Скачиваю...${NC}"
    curl -sSfL "$RAW_URL/.custom-gcl.example.yml" -o ".custom-gcl.yml"
    curl -sSfL "$RAW_URL/.golangci.example.yml" -o ".golangci.yml"
    echo -e "${GREEN}Файлы .custom-gcl.yml и .golangci.yml созданы.${NC}"
else
    echo -e "${GREEN}Файлы конфигурации уже существуют. Использую существующие.${NC}"
fi

echo -e "${GREEN}===> 3. Проверка бинарника линтера...${NC}"

if [[ ! -f "custom-gcl" ]]; then
    echo -e "${YELLOW}Бинарник не найден. Сборка (это может занять минуту)...${NC}"
    if golangci-lint custom -v; then
        echo -e "${GREEN}Сборка завершена успешно!${NC}"
    else
        echo -e "${RED}Ошибка при сборке кастомного линтера.${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}Бинарник уже существует. Использую существующий.${NC}"
fi

echo -e "${GREEN}===> 4. Запуск линтера loglint...${NC}"
./custom-gcl run

echo -e "\n${YELLOW}Проверка завершена.${NC}"
echo -e "${YELLOW}Для удаления временных файлов выполните: $0 --clean${NC}"