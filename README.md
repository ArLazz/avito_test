# Сервис динамического сегментирования пользователей
Cервис, хранящий пользователя и сегменты, в которых он состоит (создание, изменение, удаление сегментов, добавление и удаление пользователей в сегмент, а также создание истории попадания/выбывания пользователей из сегментов и выдача истории в виде ссылки на CSV файл).

Используемые технологии:
- Язык разработки: Golang
- Хранилище данных: PostgreSQL
- Для выполнения HTTP-запросов: HTTPie
- Миграции БД: golang-migrate/migrate

К сожалению, не удалось использовать docker и docker-compose для поднятия и развертывания dev-среды.

## Examples
Некоторые примеры запросов
- [Создание сегмента](#create)
- [Удаление сегмента](#delete)
- [Добавление пользователя в сегмент](#add)
- [Получение активных сегментов пользователя](#usersegments)
- [Ссылка на CSV файл с иторией](#link)



### Создание сегмента <a name="create"></a>
Создание:
```curl
http POST localhost:8080/createsegment slug_name=AVITO_VOICE_MESSAGES
```
Где slug_name - slug (название) сегмента.

Пример ответа, если сегмента с таким названием ещё нет:
```json
{
    "status": "A segment with the name: AVITO_VOICE_MESSAGES has been created."
}
```
Пример ответа, если сегмент с таким названием уже есть:
```json
{
    "status": "There is already a segment with the name: AVITO_VOICE_MESSAGES."
}
```

### Удаление сегмента <a name="delete"></a>
Удаление:
```curl
http POST localhost:8080/deletesegment slug_name=AVITO_VOICE_MESSAGES
```
Где slug_name - slug (название) сегмента.

Пример ответа, если сегмент с таким названием есть:
```json
{
   "status": "The segment with the name: AVITO_VOICE_MESSAGES has been deleted."
}
```
Где slug_name - slug (название) сегмента.
Пример ответа, если сегмента с таким названием нет:
```json
{
    "status": "The segment with the name: AVITO_VOICE_MESSAGES does not exist."
}
```
### Добавление пользователя в сегмент <a name="add"></a>
Добавление:
```curl
http POST localhost:8080/adduser user_id=9090 add_slug=AVITO_PERFORMANCE_VAS,AVITO_DISCOUNT_30 delete_slug=AVITO_VOICE_MESSAGES
```
Где user_id - id пользователя, add_slug - список (через запятую) slug (названий) сегментов которые нужно добавить пользователю , delete_slug - список (через запятую) slug (названий) сегментов которые нужно удалить у пользователя. Полей add_slug или delete_slug может не быть.

Пример ответа:
```json
{
    "status": "To the user 9090 segment/s added: AVITO_PERFORMANCE_VAS,AVITO_DISCOUNT_30 and segment/s deleted: AVITO_VOICE_MESSAGES."
}
```

### Получение активных сегментов пользователя <a name="usersegments"></a>
Получение:
```curl
 http POST localhost:8080/usersegments user_id=9090     
```
Где user_id - id пользователя.

Пример ответа:
```json
{
    "status": "The user 9090 has segments: AVITO_PERFORMANCE_VAS; AVITO_DISCOUNT_30;"
}
```

### Ссылка на CSV файл с иторией <a name="link"></a>
Сервис формирует историю, затем загружает его в Google Drive и возвращает ссылку на файл с открытым доступом на чтение.

Получение:
```curl
http POST localhost:8080/usershistory year=2023 month=8
```
Где year - год, month - месяц.

Пример ответа:
```json
{
    "link": "https://drive.google.com/file/d/1lp2Y_w3QJfAj6cJh4Ug4oP-ns81TsmRe/view?usp=sharing"
}
```
# Decisions <a name="decisions"></a>
В ходе разработки был сомнения по тем или иным вопросам, которые были решены следующим образом:
1. Что делать если пользователя хотят добавить в сегмент, в котором он уже находиться?
> Решил не дублировать сегменты, а возвращать ответ с пояснением, что пользователь уже состоит в этом сегменте. Аналогично поступил с удалением.
2. Как составлять ссылку на историю?
> В задании указано, что нужно вернуть ссылку на отчёт. Решил использовать Google Drive API с сервисным аккаунтом и создавать ссылки там.

