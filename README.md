## В чем цель калькулятора


Цель калькулятора- сделать калькулятор более быстрым и более удобным для обычных людей путем добавления парралелелизма и веб интерфейса на html и css по сравнению с предыдущей версией.



для достижения парралелизма используется go и горутины, каналы, а также инструменты синхронизации(например мьютексы или вэйт групп)


Несмотря на значительные улучшения по сравнению с версией калькулятора из 1 спринта он также умеет выполнять лишь  базовые операции  (```
/,*,-,+```
)



## Как работает:

пользователь вводит свое выражение через веб-интерфейс или же с помощью инструментов(curl,postman)

сервер возращает код 201 и id выражения в случае если пользователь ввел выражение без ошибок или же вернуть код 400 в ином случае
сервер преобразовывет пользовательское выражение из инфиниксного в обратную польскую нотацию 
сервер делает из rpn дерево ast


сервер делает из ast слайс задач для выполнения

сервер отправляет слайс к адресу /internal

адрес в свою очередь отправляет лишь 1 задачу на адрес  /internal/task и убирает эту задачу из слайса

в агенте каждая горутина обращается к адресу /internal/task и получает задачу, выполняет и отправляет результат в канал

агент вычисляет результаты горутин и отправляет результат на сервер

сервер в свою очередь отправляет результат на адрес /expressions 

пользователь может посмотреть все выполненные задачи зайдя на адрес /expressions, или же посмотреть результат нужной ему задачи по адресу /expressions/id  (вместо id должен быть реальный id задачи)



## Запросы Curl:

посмотреть, создаст ли сервер задачу без ошибок:

```curl --location  'http://localhost:8080/api/v1/calculate' --header 'Content-Type:application/json' --data '{"expression": "5+5 - 2/2"}'```

Посмотреть результат конкретно этой задачи:

```curl --location  'http://localhost:8080/api/v1/expressions/id'```

*Важно! обязательно замените id на id реальной задачи!*


Посмотреть, возращает ли сервер ошибку при неправильном выражении:




```curl --location  'http://localhost:8080/api/v1/calculate' --header 'Content-Type:application/json' --data '{"expression": "5+5 -RRRR 2/2"}'```




Посмотреть все готовые выражениия:
```curl --location  'http://localhost:8080/api/v1/expressions'```

*Важно! Если у вас не unix подобная  операционая система вы должны скачать git bash!*

а если у вас все же unix подобная система то вы можете всего лишь открыть терминал  начав делать запросы предварительно скачав git если он у вас еще не установлен командой  ```sudo apt install git```



## Как запустить проект

для копирования репозитория к вам используйте команду
```git clone https://github.com/bust6k/calc-yandex-lms-version-2.0.git```

или если у вас не получается используйте команду

```git clone git@github.com:bust6k/calc-yandex-lms-version-2.0.git```


для того чтобы запустить сам сервер введите команду ```go run .```

для запуска тестов введите команду ```go test```


для запуска теста правильности преобразования в rpn введите команду ```go test -run=TestInfixToRPN```

для запуска теста правильности преобразования rpn в ast введите команду ```go test -run=TestBuildAST```

для запуска теста правильности парсинга ast в tasks  введите команду ```go test -run=TestSplitAST```


для запуска теста правильности вычислений агента введите команду   ```go test -run=TestCreateRootExpressionHandlerGoodCase```

*перед запуском этого теста запустите сервер чтобы не было паники*


для запуска теста правильности  того что сервер возращает ошибку в случае неккоректного выражения введите команду ```go test -run=TestCreateRootExpressionHandlerBadCase```



##  Информация о пакетах

*Agent*- содержит код агента а также код API для взаимодействия агента между сервером

*calc*- содержит код калькулятора из 0 спринта

*important*- является очень важным пакетом. Тут хранится основная логика программы


*structures* - содержит структуры которые могут понадобиться другим пакетам

*Tests* - содержит тесты для проверки правильности программы

*ui/html*- содержит страницы (если точнее то только 1 страницу) для frontend

*variables* - содержит переменные которые могут понадобиться другим пакетам
















