# remote-desktop
## Naming
* host - устройство к которому осуществляется доступ
* client - устройство с которого производится доступ к хосту
* server - соединяет клиент и хост 

## Actions
Описание видов действий, и формата данных в сообщениях. Для простоты опускается упоминание полей action и data.
Если не описывается структура поля data то это строка. Сгрупированы по получателю\обработчику.

**Server:**
* HOST_REGISTER - Регистрация хоста на сервере. Хост отправляет свой логин.
* GET_HOSTS - Запрос клиентом списка доступных хостов у сервера.

**Host:**
* MOUSE_MOVE - Перемещение мыши.
* MOUSE_RCLICK - Правый клик мыши.
* MOUSE_LCLICK - Левый клик мыши.

**Client:**
* IMG_FRAME - Кадр с хоста.
