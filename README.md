<h1>Инструкция по запкуску</h1>
<h2>Локально</h2>
1.Поднять postgres у себя
2.Составать конфиг в формате json, в следующем вид
```
{
  "address": ":8080", 
  "timeoutDB": 10, //Сколько программа будет ждать ответа от базы данных
  "env": "local", //[dev, local, prod] В зависмости от выбранного параметра, меняются настройки логера, local - логи будут выводится в консоль, dev - логи будут выводится в файл
  "driverName": "postgres",
  "connStr": "postgres://postgres:qwerty@172.17.0.2:5432/commands?sslmode=disable", //строка подключения к базе даных
  "migrationsDir": "./migrations" //путь до фалйов миграции
}
```
3. Скомпилировать и запустить приложение указав -path={путь до файлов конфигурации}
<h2>Через docker-compose</h2>
1. В файле конфигурации docker-compose, используются следующие переменные окружения
   POSTGRES_USER
   POSTGRES_PASSWORD
   POSTGRES_DB
   
   Можно задать их через .env файл
2.Составать конфиг в формате json, в следующем виде
```
{
  "address": ":8080", 
  "timeoutDB": 10, //Сколько программа будет ждать ответа от базы данных
  "env": "local", //[dev, local, prod] В зависмости от выбранного параметра, меняются настройки логера, local - логи будут выводится в консоль, dev - логи будут выводится в файл
  "driverName": "postgres",
  "connStr": "postgres://POSTGRES_USER:POSTGRES_PASSWORD@db:5432/POSTGRES_DB?sslmode=disable", //строка подключения к базе даных, использовать POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB - которые указали в .env файле
  "migrationsDir": "./migrations" //путь до фалйов миграции
}
```
3. Собрать с помощью комманды docker-compose build
4. Запустить с помощью комманды docker-compose up
