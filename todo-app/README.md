# Todo API

Bu proje, Go ve Gin framework kullanılarak geliştirilmiş bir TODO liste yönetim API'sidir.

## Özellikler

- JWT tabanlı kimlik doğrulama
- İki farklı kullanıcı tipi (Normal ve Admin)
- Todo listeleri ve todo itemları için CRUD işlemleri
- Soft delete özelliği
- Todo tamamlanma yüzdesi takibi

## Kurulum

1. Projeyi klonlayın:

    ```bash
    git clone https://github.com/yourusername/todo-app.git
    cd todo-app
    ```

2. Bağımlılıkları yükleyin:

    ```bash
    go mod tidy
    ```

3. Uygulamayı çalıştırın:

    ```bash
    go run main.go
    ```

   Uygulama varsayılan olarak 8080 portunda çalışacaktır.

### Test Kullanıcıları

- **Normal Kullanıcı**:
    - Username: user1
    - Password: password1

- **Admin Kullanıcı**:
    - Username: admin
    - Password: admin123

## API Endpoint'leri

### Authentication

- **POST /login** - Giriş yapma ve token alma

### Todo İşlemleri (Protected Routes)

- **GET /api/todos** - Tüm todoları listele
- **POST /api/todos** - Yeni todo oluştur
- **PUT /api/todos/:id** - Todo güncelle
- **DELETE /api/todos/:id** - Todo sil
- **POST /api/todos/:id/items** - Todo'ya item ekle
- **PUT /api/todos/:id/items/:itemId** - Todo item güncelle
- **DELETE /api/todos/:id/items/:itemId** - Todo item sil

## API Kullanımı

### cURL ile Test

1. **Login ve Token Alma**:

    ```bash
    curl -X POST http://localhost:8080/login \
    -H "Content-Type: application/json" \
    -d '{
        "username": "user1",
        "password": "password1"
    }'
    ```

2. **Yeni Todo Oluşturma**:

    ```bash
    curl -X POST http://localhost:8080/api/todos \
    -H "Authorization: Bearer <your_token>" \
    -H "Content-Type: application/json" \
    -d '{
        "name": "My First Todo List"
    }'
    ```

3. **Todoları Listeleme**:

    ```bash
    curl -X GET http://localhost:8080/api/todos \
    -H "Authorization: Bearer <your_token>"
    ```

4. **Todo'ya Item Ekleme**:

    ```bash
    curl -X POST http://localhost:8080/api/todos/1/items \
    -H "Authorization: Bearer <your_token>" \
    -H "Content-Type: application/json" \
    -d '{
        "content": "Buy groceries",
        "is_completed": false
    }'
    ```

5. **Todo Item Güncelleme**:

    ```bash
    curl -X PUT http://localhost:8080/api/todos/1/items/1 \
    -H "Authorization: Bearer <your_token>" \
    -H "Content-Type: application/json" \
    -d '{
        "content": "Buy groceries",
        "is_completed": true
    }'
    ```

### Postman ile Test

**Postman koleksiyonunu import etmek için**:

1. Postman'i açın
2. Collections > Import düğmesine tıklayın
3. Aşağıdaki JSON'ı yapıştırın:

```json
{
    "info": {
        "name": "Todo API",
        "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
    },
    "item": [
        {
            "name": "Login",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "url": "http://localhost:8080/login",
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"username\": \"user1\",\n    \"password\": \"password1\"\n}"
                }
            }
        },
        {
            "name": "Create Todo",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "url": "http://localhost:8080/api/todos",
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"name\": \"My First Todo List\"\n}"
                }
            }
        },
        {
            "name": "Get All Todos",
            "request": {
                "method": "GET",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{token}}"
                    }
                ],
                "url": "http://localhost:8080/api/todos"
            }
        },
        {
            "name": "Add Todo Item",
            "request": {
                "method": "POST",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "url": "http://localhost:8080/api/todos/1/items",
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"content\": \"Buy groceries\",\n    \"is_completed\": false\n}"
                }
            }
        },
        {
            "name": "Update Todo Item",
            "request": {
                "method": "PUT",
                "header": [
                    {
                        "key": "Authorization",
                        "value": "Bearer {{token}}"
                    },
                    {
                        "key": "Content-Type",
                        "value": "application/json"
                    }
                ],
                "url": "http://localhost:8080/api/todos/1/items/1",
                "body": {
                    "mode": "raw",
                    "raw": "{\n    \"content\": \"Buy groceries\",\n    \"is_completed\": true\n}"
                }
            }
        }
    ]
}
```

### Postman Kullanım Adımları

1. Önce "Login" isteğini gönderin ve dönen response'dan token'ı kopyalayın.
2. Diğer isteklerin Authorization header'ında `Bearer <token>` formatında token'ı kullanın.
3. İstekleri sırayla test edin:
    - Todo oluşturma
    - Todoları listeleme
    - Todo'ya item ekleme
    - Item'ı güncelleme

### Notlar

- Token'lar 24 saat geçerlidir.
- Normal kullanıcılar sadece kendi todo'larını görebilir ve değiştirebilir.
- Admin kullanıcıları tüm todo'lara erişebilir.
- Silme işlemleri soft delete şeklinde gerçekleşir.
- Todo completion yüzdesi, tamamlanan item'ların toplam item sayısına oranıdır.
