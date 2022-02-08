# Простой API server для программы "Факультетские часы Хогвартса"

Написан на Go:

- сервер [Gin](https://github.com/gin-gonic/gin)
- база данных SQLite
- ORM [Gorm](https://github.com/go-gorm/gorm)

## Схема эндпоинтов API

### Формат ответа на `GET` запрос по адресу `localhost:8080/faculties`

```json
{
    "data": [
        {
            "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
            "id": 1,
            "name": "Гриффиндор",
            "score": 50,
            "color": "0xFFF44336"
        },
        {
            "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
            "id": 2,
            "name": "Пуффендуй",
            "score": 30,
            "color": "0xFFFF9800"
        },
        {
            "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
            "id": 3,
            "name": "Когтевран",
            "score": 20,
            "color": "0xFF2196F3"
        },
        {
            "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
            "id": 4,
            "name": "Слизерин",
            "score": 10,
            "color": "0xFF4CAF50"
        }
    ]
}
```

Где `color` это string шестнадцатеричное значение [Color класса](https://api.flutter.dev/flutter/dart-ui/Color-class.html)  

### Формат тела `PUT` запроса по адресу `localhost:8080/faculties` для внесения изменений в базу данных

```json
[
    {
        "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
        "id": 1,
        "name": "Гриффиндор",
        "score": 50,
        "color" : "0xFFF44336"
    },
    {
        "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
        "id": 2,
        "name": "Пуффендуй",
        "score": 30,
        "color": "0xFFFF9800"
    },
    {
        "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
        "id": 3,
        "name": "Когтевран",
        "score": 20,
        "color": "0xFF2196F3"
    },
    {
        "branchID": "f9afda82-8774-11ec-a8a3-0242ac120002",
        "id": 4,
        "name": "Слизерин",
        "score": 10,
        "color": "0xFF4CAF50"
    }
]

```

### Запуск сервера

```bash
./house_points_server
```
