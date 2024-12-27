# •• pingpongerbot

**pingpongerbot** — это Telegram-бот, написанный на Go, который отвечает на команды и выполняет простые действия. Это учебный проект, созданный для изучения работы с Telegram Bot API и библиотекой Telebot.

<img src="https://github.com/user-attachments/assets/957ebb07-10ad-4ea2-82f4-b1118a855556">

## Возможности

**Основные команды**:
- **/start** - приветствие
- **/ping** - бот скажет «pong!» в ответ
- **/echo** - напиши /echo <любой текст здесь>, бот продублирует твое сообщение

**Обработка некорректного ввода**:  
Любой неизвестный текст или медиа сообщение перенаправляет пользователя на приветственное сообщение.

**Логирование**:   
- Логирует взаимодействия с пользователями (ID чата, имя пользователя, текст сообщения).
- Логирует ошибки для их дальнейшего анализа.

**Обработка ошибок**:   
Показывает пользователю дружелюбное сообщение об ошибке и сохраняет её детали в логах.

---

## Технологии

- **ЯП**: Go (Golang)
- **Telegram Bot Framework**: [telebot.v4](https://github.com/tucnak/telebot)
- **Управление конфигурацией**: [joho/godotenv](https://github.com/joho/godotenv) для работы с переменными окружения.
- **Логирование**: [rs/zerolog](https://github.com/rs/zerolog) для структурированного логирования.

---

## Структура проекта

```
pingpongerbot/
├── cmd/
│   └── pingpongerbot/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── handlers/
│   │   └── handlers.go
├── middleware/
│   ├── error.go
│   └── logger.go
├── .gitignore
├── go.mod
└── go.sum
```

---

## Детали реализации

**main.go**
- Инициализирует приложение с помощью InitApp.
- Добавляет middleware для обработки ошибок и логирования.
- Регистрирует команды.
-	Обрабатывает неизвестные сообщения через группу unrecognized.

**app.go**
- Конфигурирует бота с использованием telebot.NewBot.
- Устанавливает токен, настройки поллера и обработку ошибок.
- Опция оффлайн-режима для тестирования.

**handlers.go**
- Start: Приветственное сообщение, описание и список доступных команд.
- Ping: Отправляет “pong!”.
- Echo: Повторяет сообщение пользователя.

**logger.go:**   
- Логирует обновления (ID чата, имя пользователя, текст сообщения).
- Логирует ошибки с дополнительной информацией.

**error.go:**   
Отправляет пользователю сообщение об ошибке и записывает её детали.

## Лицензия

[MIT](https://mit-license.org/)
