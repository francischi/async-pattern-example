# **Async Pattern Example - Consumer**

## **Introduction**
此為message broker 中的consumer，負責對rabbitmq進行連線並發出信件，
對於流量較大時，可啟動多個consumer服務，以分散流量。

*後續再補上queue內部錯誤處理之部分