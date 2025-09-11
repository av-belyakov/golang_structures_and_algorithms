# Как создать нового пользователя.

#### 1. Создаём новый alias:

```bash
mc alias set miniouserroot http://localhost:9000 admin password
```

(где miniotuserroot - псевдоним пользователя root). Создание выполняется от пользователя root, то есть под его учётными данными.

Удаление выполняется командой:

```bash
mc alias remove miniouserroot
```

#### 2. Проверить наличие alias:

```bash
mc ls или mc ls miniotuserroot
```

#### 3. После того как подключились к серверу MinIO, создаём пользователя:

```bash
mc admin user add miniotuserroot testuser 123qwe123
```

#### 4. Создаём политику доступа:

```bash
mc admin policy attach minioadmin readwrite --user gcm
```

#### 5. Теперь можно посмотреть список всех пользователей:

```bash
mc admin user list miniotuserroot
```

#### 6. Смотрим информацию о пользователе:

```bash
mc admin user info miniotuserroot gcm
```

#### 7. Удаление пользователя:

```bash
mc adminuser remove miniotuserroot gcm
```

#### 8. Изменение пароля пользователя:

```bash
mc admin user disable miniotuserroot gcm
```

сначала отключаем

```bash
mc admin user remove miniotuserroot gcm
```

удаляем

```bash
mc admin user add miniouserroot <пользователь> <пароль>
```

создаём нового пользователя

Включение/отключение пользователя:

```bash
mc admin user disable miniouserroot gcm
mc admin user enable miniouserroot gcm
```
