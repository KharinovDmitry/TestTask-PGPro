<h1>Инструкция по запуску</h1>
<h2>Локально</h2>
<li>Поднять postgres у себя</li>
<li>Составать конфиг в формате json, в следующем виде</li>
<code>{
  "address": ":8080", 
  "timeoutDB": 10, //Сколько программа будет ждать ответа от базы данных
  "env": "local", //[dev, local, prod] В зависмости от выбранного параметра, меняются настройки логера, 
                  //local - все логи будут выводится в консоль,
                  //dev - все логи будут выводится в файл
                  //prod - логи уровня INFO и выше будут выводится в файл (логи уровня debug выводится не будут)
  "driverName": "postgres",
  "connStr": "postgres://postgres:qwerty@172.17.0.2:5432/commands?sslmode=disable", //строка подключения к базе даных
  "migrationsDir": "./migrations" //путь до фалйов миграции
}</code>
<li>Скомпилировать и запустить приложение указав -path={путь до файлов конфигурации}</li>
<h2>Через docker-compose</h2>
<li>В файле конфигурации docker-compose, используются следующие переменные окружения. Можно задать их через .env файл</li>
    <ul>POSTGRES_USER</ul>
    <ul>POSTGRES_PASSWORD</ul>
    <ul>POSTGRES_DB</ul>
<li>Составать конфиг в формате json, в следующем виде, сохранить его с названием dev_config.json и положить в internal/config</li>
<code>{
  "address": ":8080", 
  "timeoutDB": 10, //Сколько программа будет ждать ответа от базы данных
  "env": "local", //[dev, local, prod] В зависмости от выбранного параметра, меняются настройки логера, 
                  //local - все логи будут выводится в консоль,
                  //dev - все логи будут выводится в файл
                  //prod - логи уровня INFO и выше будут выводится в файл (логи уровня debug выводится не будут)
  "driverName": "postgres",
  "connStr": "postgres://POSTGRES_USER:POSTGRES_PASSWORD@db:5432/POSTGRES_DB?sslmode=disable", //строка подключения к базе даных, использовать POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB - которые указали в .env файле
  "migrationsDir": "./migrations" //путь до фалйов миграции
}</code>
<li>Собрать с помощью комманды docker-compose build</li>
<li>Запустить с помощью комманды docker-compose up</li>
<h1>Что дополнительно реализованно</h1>
<li>Поддержка долгих команд (сохранять вывод команды в БД по мере ее выполнения, отображать вывод при получении одной команды)</li>
<p>Не совсем понял, что имеется виду под отображать вывод при получении одной команды, было решено сделать две таблицы с коммандами и запусками, запустить команду можно по эндпоинту /run/{id комманды},
Вернется айди запуска, по которому можно узнать вывод комманды по эндпоинту /launch/{id запуска}, остановить по эндпоинту /stop/{id запуска}</p>
<li>Добавить в API метод для остановки команды</li>
<li>Сборка и деплой приложения</li>
<p>Написн dockerfile и docker-compose.yml для запуска приложения и базы данных в контейнерах</p>
<h1>Используемые технологии</h1>
<li>Локально, писал на linux/ubuntu, в контейнере запускается на alpine</li>
<p>Помимо стандратной библиотеки были использованы</p>
<li>github.com/jmoiron/sqlx - для работы с базами данных</li>
<li>github.com/jackc/pgx - в качестве драйвера для postgres</li>
<li>github.com/pressly/goose - для работы с миграциями</li>


