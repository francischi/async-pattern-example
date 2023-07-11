# **Async Pattern Example - APP**

## **Introduction**
此為主要後端APP，有會員與訂單相關功能，詳細請參考 api-document.md
此後端購物下單功能僅為了展示下單成功時發email如何使用message broker流程，並無下單時扣庫存動作相關實作。

### connection pool
message broker連線開銷較大，因此app初始化時，就建立對message broker之連線，
並建立連線池，讓連線能夠復用，不用一直創造新連線。