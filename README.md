# Информационной система для POS-кредитования

## Описание задачи

Необходимо разработать веб сервис, со следующим функционалом:
1. Создание, просмотр, изменение, удаление информации о кредитных
агентах, магазинах, торговых точках и доступах в системе через
админку сайта.
1. Форма оформления кредита для агента на торговой точке.
1. Механизм контроля доступа по логину и паролю, сookies-сессии.

Администратор может выполнять действия по созданию, просмотру, изменению, удалению кредитных
агентов, магазинов, торговых точек и доступов в системе посредством админ панели. Имеет атрибуты –
ФИО, логин, пароль. В системе может быть только один аккаунт администратора, его данные задаются
в конфигурационных файлах или в форме начальной инициализации платформы.

Кредитный агент имеет атрибуты – ФИО, логин, пароль. Логин является уникальным. Кредитный агент
может получить доступ только к форме оформления кредита.

Магазин имеет атрибуты – название магазина и его символьный код, представляющие собой
произвольные строки. Символьный код является уникальным.

Торговая точка имеет атрибуты – символьный код магазина (из числа зарегистрированных магазинов) и
название торговой точки. Символьный код является уникальным.
## Реализация
- Дизайн REST API и "чистая архитектура".
- Работа с фреймворком [gin-gonic/gin](https://github.com/gin-gonic/gin).
- Работа с СУБД Postgres с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Запуск из Docker.

### Для запуска приложения:

```
make build && make run
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
make migrate
```
