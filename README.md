# (ENG) TUI Project "ISU"

An information system of the university that allows creating groups, distributing students among groups, and transferring.

The main goal of the project is to work with the Bubble Tea framework.

## Available Functionality
The following features are implemented:</br>
- Enroll a student in a group
- Create a new group
- Transfer a student to another group

# (RU) TUI проект "ISU"
Информационная система университета позволяющая создавать группы, распределять студентов по группам и переводить

Главная цель проекта поработать с фреймворком Bubble Tea 

## Доступный функционал
Реализована возможность:</br>
- Зачислять студента в группу
- Создавать новую группу
- Переводить студента в другую группу

# Project Launch//Запуск проекта
1. Clone the repository// Клонируйте репозиторий
 ```bash
 git clone https://github.com/Vatset/ISU.git
```
3. Prepare the database for operation // Подготовьте бд к работе<br>
  *Download and run the Docker application beforehand* // *Предварительно скачайте и запустите приложение docker*<br>
Obtain the latest version of PostgreSQL // Получение последней версии postgres
```bash   
docker pull postgres
```
Run a Docker container named "isu" using the previously downloaded PostgreSQL  // Запуск Docker контейнера с именем "isu", используя ранее скачанный образ PostgreSQL. 
```bash
docker run --name=isu -e POSTGRES_PASSWORD="mewmew" -p 5436:5432 -d --rm postgres
```
Execute database migrations//Выполнение миграций базы данных
```bash 
migrate -path ./schema  -database 'postgres://postgres:mewmew@localhost:5436/postgres?sslmode=disable' up
```
4.Launch the project//Запускаем проект
```bash   
go run cmd/main.go
```
