# **Async Pattern Example**
## **Introduction**
使用訂單系統為範例，以符合clean architecture 之方式撰寫，示範如何使用rabbitmq進行非同步任務
之處理，避免不必要網路IO等待，並實作了message broker的connection pool，
以此方式建構分散式服務能夠提高整體效能，且服務之間的鬆耦合有助於後續維護、再開發與水平擴展。

## **技術點**
* Clean Architecture
* Message Broker ( RabbitMQ )
* Message Broker Connection Pool
* Docker

## **架構**
```
producer ---> message queue ---> consumer
```

## **使用場景**
通常應用在不關心結果之事務(非同步)，或不用即時處理之事務。
* 機台接收數據
```
機台數據較多可以使用message queue當buffer，再啟動多個consumer消耗這些數據。
```
* 發信件
```
先將要發之信件內容與收件者存入message queue，再慢慢消耗。
```
* 寫log
```
將要寫的log內容放入message queue，再慢慢消耗。
```
* 第三方進行足跡追蹤
```
若串接FB 或 google行銷之事件，需要頻繁傳送事件到fb or google，
就可以使用message queue，才不會降低原API之效率。
```

## 使用方式
* rabbitmq
```
> docker pull rabbitmq:management
> docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management

透過localhost:15672查看rabbitmq是否啟動
```

* app
```
> cd ./App
> go run main.go
```

* consumer
```
> cd ../MailConsumer
> go run main.go
```

<br>

```
PS 此專案架構較適合中大型專案開發使用
```