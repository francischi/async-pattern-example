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