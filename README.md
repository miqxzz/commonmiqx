# Common JWT Module

Модуль для работы с JWT-аутентификацией на Go.

## Установка

```bash

go get github.com/miqxzz/commonmiqx
```

## Использование

```go
import "github.com/miqxzz/commonmiqx"

// Создание нового JWT утилиты
jwtUtil := common.NewJWTUtil("ваш_секретный_ключ")

// Генерация токена
token, err := jwtUtil.GenerateToken(userID, role)

// Валидация токена
claims, err := jwtUtil.ValidateToken(tokenString)

// Получение ID пользователя из токена
userID, err := jwtUtil.GetUserIDFromToken(tokenString)

// Получение роли пользователя из токена
role, err := jwtUtil.GetRoleFromToken(tokenString)
```

## Функциональность

- Генерация JWT токенов
- Валидация токенов
- Извлечение данных пользователя из токена
- Настраиваемое время жизни токена 