# KVANTAKT_PlanNyam

Бэкенд приложения меню\
Описание endpoint'ов:

1 Регистрация:\
**POST** -> **localhost:8000/auth/sign-up** -> **{"name": "admin","username": "admin","password": "admin"}**\
При успешном запросе получаем -> **{"status": "ok"}** либо сервис вернет ошибку

2 Авторизация:\
**POST** -> **localhost:8000/auth/sign-in** -> **{"username": "admin","password": "admin"}**\
При успешной авторизации получаем 'Bearer Token' авторизации который действителен 12 часов -> **{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTc"}**\
Данный Bearer Token подставляется в Headers всех следующих запросов для отправки их от лица авторизованного поьзователя

3 Создание рецепта:\
**POST** -> **localhost:8000/api/lists** -> **{"title": "Название рецепта","description": "Описание рецепта"}**\
При успешном создании рецепта получаем его id -> **{"id": 1}**

4 Получение всех рецептов пользователя:\
**GET** -> **localhost:8000/api/lists** - > С пустым телом запроса\
Получаем список всех рецептов пользователя -> **{"data": [{"id": 1,"title": "Борщ","description": "Ингридиенты для приготовления борща!!!"}]}**

5 Создание ингридиентов рецепта:\
**POST** -> **localhost:8000/api/lists/1/items** -> **{"title": "Картошка", "quantity": "3", "price": "150", "description": "По требованию"}**\
Единица (1) в ендпоинте это id рцепта, куда будет вставляться ингридиент\
При успешном выполнении запроса вернется id данного ингридиента -> **{"id_add_items": 13}**

6 Получение всех ингридиентов рецепта\
**GET** -> **localhost:8000/api/lists/1/items/ -> Пустое тело запроса\**
Получаем массив всех ингридиентов даннаго рецепта 1 -> **[{"id": 1,"title": "Свекла","quantity": "6","price": "500","description": "Нужный ингридиент","done": false},{"id": 13,"title": "Картошка","quantity": "3","price": "150","description": "По требованию","done": false}]**

7 Получение одного ингридиента по id:\
**GET** -> **localhost:8000/api/items/13** -> Пустое тело запроса\
Получаем данные ингридиента под id=13 -> **{"id": 13,"title": "Картошка","quantity": "3","price": "150","description": "По требованию","done": false}**

8 Редактирование ингридиента:\
**PUT** -> **localhost:8000/api/items/13** -> **{"title": "Картошка","quantity": "9","price": "450","description": "Необходимо"}**\
При успешном выполнении запроса у ингридиента с id=13 возвращается -> **{"status": "ok"}**, либо ошибка\
Можно редактировать как все поля сразу, так и каждое по отдельности

9 Удаление ингридиентов из рецептов:\
**DELETE** -> **localhost:8000/api/items/2** -> Пустое тело запроса\
Возвращается ответ после удаленного элемента с id=2 -> **{"status": "ok"}**

10 Удаление рецептов:\
**DELETE** -> **localhost:8000/api/lists/2** -> Пустое тело запроса\
Возвращается ответ после удаленного рецепта с id=2 -> **{"status": "ok"}**

11 Изменение название или описания рецепта:\
**PUT** -> **localhost:8000/api/lists/2 -> {"title": "Борщ99уровня","description": "Божественые ингридиенты"}**\
Возвращает ответ -> **{"status": "ok"}**\
Можно редактировать как оба поля сразу, так и каждый по отдельности\

12 Добавление игридиентов из чека:\
**POST** -> **localhost:8000/api/lists/1/jsonParse** -> В теле запроса идет json из отсканированного чека\
Тело запроса можете посмотреть в папке files/test.json
Сам успешно отработавший запрос возвращает id добавленных товаров -> **{"id_add_items": [16,17]}**