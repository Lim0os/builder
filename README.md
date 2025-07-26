# Project Builder

Генератор проектов на Go с архитектурой портов и адаптеров (Hexagonal Architecture).

## Описание

Project Builder - это CLI инструмент для быстрого создания новых Go проектов с предустановленной архитектурой портов и адаптеров. Генератор создает базовую структуру проекта с готовыми шаблонами файлов, следуя принципам чистой архитектуры.

## Архитектура

Проект использует архитектуру портов и адаптеров (Hexagonal Architecture):

```
src/
├── application/          # Слой приложения (бизнес-логика)
│   ├── application.go
│   └── commands/         # Команды приложения
├── domain/              # Доменный слой (сущности и интерфейсы)
├── ports_adapters/      # Порты и адаптеры
│   ├── primary/         # Первичные адаптеры (входящие)
│   │   └── cmd/         # CLI команды
│   └── secondary/       # Вторичные адаптеры (исходящие)
└── common/              # Общие утилиты
    ├── config/          # Конфигурация
    └── decorator/       # Декораторы
```

## Установка

```bash
# Клонирование репозитория
git clone https://github.com/Lim0os/builder.git
cd builder

# Сборка проекта
go build -o bin/builder cmd/builder/main.go

# Установка в систему (опционально)
sudo cp bin/builder /usr/local/bin/
```

## Использование

### Создание нового проекта

```bash
# Базовое использование
./bin/builder create my-project

# Или если установлен в систему
builder create my-project
```

### Структура создаваемого проекта

После выполнения команды `create` будет создана следующая структура:

```
my-project/
├── src/
│   ├── application/
│   │   └── app.go
│   ├── domain/
│   │   ├── entity.go
│   │   └── repository.go
│   └── ports_adapters/
│       └── primary/
│           └── http_server/
│               └── server.go
├── go.mod
├── main.go
├── Makefile
├── Dockerfile
├── README.md
└── .gitignore
```

## Генерируемые файлы

- **src/application/app.go** - Основное приложение
- **src/domain/entity.go** - Доменные сущности
- **src/domain/repository.go** - Интерфейсы репозиториев
- **src/ports_adapters/primary/http_server/server.go** - HTTP сервер
- **go.mod** - Модуль Go с базовыми зависимостями
- **main.go** - Точка входа в приложение
- **Makefile** - Команды для сборки и запуска
- **Dockerfile** - Контейнеризация приложения
- **README.md** - Документация проекта
- **.gitignore** - Исключения Git

## Требования

- Go 1.24.5 или выше
- Linux/macOS/Windows

## Разработка

### Структура проекта

```
project_builder/
├── cmd/
│   └── builder/         # Точка входа CLI
├── src/
│   ├── application/     # Слой приложения
│   ├── domain/          # Доменный слой
│   ├── ports_adapters/  # Порты и адаптеры
│   └── common/          # Общие утилиты
└── go.mod
```

### Добавление новых команд

1. Создайте новую команду в `src/application/commands/`
2. Добавьте обработчик в `cmd/builder/main.go`
3. Обновите роутер команд

### Добавление новых шаблонов

1. Добавьте константы шаблонов в `src/domain/templates.go`
2. Обновите `NewProjectTemplate()` в `create_project_command.go`

## Команды

- `create <project-name>` - Создает новый проект с указанным именем

## Логирование

Проект использует структурированное логирование с помощью `log/slog`. Уровень логирования по умолчанию: `debug`.

## Лицензия

MIT License

## Автор

Lim0os

## Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции
3. Внесите изменения
4. Создайте Pull Request

## Поддержка

Если у вас есть вопросы или предложения, создайте Issue в репозитории. 