# Golang_project_site: Веб-сервис для вычисления арифметических выражений

## Описание
Данный проект представляет собой веб-сервис, который вычисляет арифметическое выражение, переданное пользователем через HTTP-запрос

## Установка и запуск
1. Для начала следует скопировать репозиторий:
   ```git init mod https://github.com/DRAGONBORN1999/Golang_project_site```
2. Затем можно запустить сервер командой:
   ```go run ./cmd/main.go```
3. После этого можно посылать запросы

При этом у вас должен быть установлен Golang: https://go.dev/doc/install

## Примеры запроса и возможные ошибки
1. Пример верного запроса: ```curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"10+2\"}" http://localhost:8080```
При этом будет получен ответ: ```{"result":12}```
2. Пример неверного запроса: ```curl -X POST -H "Content-Type: application/json" -d "{\"тфьу\": \"10+2\"}" http://localhost:8080```
При этом будет получен ответ: ```{"error":"Bad request"}```
3. Пример неверного выражения: ```curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"ssssss\"}" http://localhost:8080```
При этом будет получен ответ: ```{"error":"Expression is not valid"}```, то есть, если выражение некорректно (присутствуют сторонние символы помимо цифр, символов операций и скобок), возвращается ошибка
