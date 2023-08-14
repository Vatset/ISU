# TUI проект "ISU"
Информационная система университета позволяющая создавать группы, распределять студентов по группам и переводить

Главная цель проекта поработать с фреймворком Bubble Tea 
## Запуск проекта
1. Клонируйте репозиторий
 ```bash
 git clone https://github.com/Vatset/ISU.git
```
3. Подготовьте бд к работе<br>
   *Предварительно скачайте и запустите приложение docker*<br>
Получение последней версии postgres
```bash   
docker pull postgres
```
Запуск Docker контейнера с именем "isu", используя ранее скачанный образ PostgreSQL. 
```bash
docker run --name=isu -e POSTGRES_PASSWORD="mewmew" -p 5436:5432 -d --rm postgres
```
Выполнение миграций базы данных
```bash 
migrate -path ./schema  -database 'postgres://postgres:mewmew@localhost:5436/postgres?sslmode=disable' up
```
4.Запускаем проект
```bash   
go run cmd/main.go
```
## Доступный функционал
Реализована возможность :
*Регистрировать студента
*Создавать новую группу
*Переводить студента в другую группу
