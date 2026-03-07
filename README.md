# loglint 🔍

**loglint** — линтер, совместимый с golangci-lint, который проверяет лог-записи вашего проекта.
Он помогает поддерживать единообразие логов во всем проекте и предотвращает утечку чувствительных данных.

## 🎯 Возможности

- **Проверка регистра первой буквы** — все сообщения логов должны начинаться со строчной буквы
- **Валидация символов** — запрет на использование недопустимых символов (кириллица, спецсимволы, эмодзи)
- **Поиск чувствительных данных** — обнаружение паролей, токенов, ключей API в логах
- **Поддержка популярных логгеров**:
  - `log/slog`
  - `go.uber.org/zap`
- **Интеграция с golangci-lint** — работает как плагин к нему

## 📋 Требования

- Go 1.21 или выше
- [golangci-lint](https://golangci-lint.run/) версии 1.64.5 или выше

## 🚀 Быстрый старт



### Запуск loglint

Скопируйте и выполните одну команду:

```bash
curl -sSfL https://raw.githubusercontent.com/OWEEN3/loglint/main/install_and_run.sh | bash
```

Эта команда:
1. Скачает конфигурационные файлы
2. Соберет loglint как плагин к golangci-lint
3. Запустит проверку вашего кода

**При первом запуске** файлы будут скачаны и собран бинарник.  
**При повторном запуске** (в той же папке) будут использоваться существующие файлы.

### Очистка временных файлов

```bash
curl -sSfL https://raw.githubusercontent.com/OWEEN3/loglint/main/install_and_run.sh | bash -s -- --clean
```

## 📝 Примеры

### Что проверяет loglint

```go
package main

import (
    "log/slog"
    "go.uber.org/zap"
)

func main() {
    // НЕПРАВИЛЬНО: первая буква заглавная
    slog.Info("Starting server on port 8080")  // Ошибка: first letter should be lowercase
    
    // НЕПРАВИЛЬНО: кириллица
    log.Error("ошибка подключения")  // Ошибка: invalid letters
    
    // НЕПРАВИЛЬНО: эмодзи и спецсимволы
    slog.Info("server started!🚀")  // Ошибка: invalid letters
    
    // НЕПРАВИЛЬНО: чувствительные данные
    password := "secret123"
    slog.Info("user password: " + password)  // Ошибка: sensitive data
    
    // ПРАВИЛЬНО: все требования соблюдены
    slog.Info("server started on port 8080")
    log.Error("failed to connect to database")
}
```

## 🔧 Ручная установка

Если вы хотите установить loglint вручную или лучше понять процесс, выполните следующие шаги:

### Шаг 1: Создание конфигурационных файлов

Создайте файл `.golangci.yml` в корне вашего проекта:

```yaml
# .golangci.yml
linters:
  enable:
    - loglint

linters-settings:
  loglint:
    # настройки будут добавлены позже
```

Создайте файл `.custom-gcl.yml` для сборки кастомного бинарника:

```yaml
# .custom-gcl.yml
plugins:
  - module: github.com/OWEEN3/loglint
    import: github.com/OWEEN3/loglint/plugin
    version: latest # или конкретная версия, например v1.0.0
```

Или просто скопируйте и переименуйте существующие из этого репозитория.

### Шаг 2: Сборка кастомного линтера

Выполните команду для сборки loglint как плагина к golangci-lint:

```bash
# Сборка кастомного бинарника
golangci-lint custom -v
```

Эта команда создаст исполняемый файл `custom-gcl` в текущей директории.

### Шаг 3: Запуск линтера

Запустите проверку вашего кода:

```bash
# Запуск линтера
./custom-gcl run
```

### Шаг 4: Очистка (опционально)

После завершения работы вы можете удалить временные файлы:

```bash
# Удаление конфигов и бинарника
rm -f .custom-gcl.yml .golangci.yml custom-gcl
```

## 💡 Совет:

Все эти шаги автоматически выполняет скрипт быстрого старта


## 🏗️ Структура проекта

```
loglint/
├── cmd/                    # точка входа
│   └── loglint/           
├── pkg/                    # основная логика
│   └── analyzer/           
│       ├── analyzer.go     # код анализатора
│       └── rules/          # правила проверки
├── plugin/                 # плагин для golangci-lint
│   └── loglint.go          
├── testdata/               # тестовые файлы
└── install_and_run.sh      # скрипт для быстрого запуска
```

