# 抽象工廠模式 

`DAO = Data Access Object`

最上層有一個抽象工廠 DAOFactory的介面
介面需要實作 CreateOrderMainDAO() 和 CreateOrderDetailDAO()

建立OrderMainDAO 和 OrderDetailDAO 的介面
並分別建立 RDB 和 XML 類型的 DAO 來實作

最後實驗的時候
可以指派是 RDBDAOFactory 或 XMLDAOFactory 其中一種
再呼叫 CreateOrderMainDAO() 的 SaveOrderMain()
和 CreateOrderDetailDAO() 的 SaveOrderDetail()

提示:
抽象工廠模式用於生成做產品的工廠
所生成的對象是有關聯的
如果沒有關聯的話則退化回工廠方法模式
比如 RDB 和 XML 的例子，
抽象工廠分別能生成相關的主訂單訊息和訂單詳細訊息
如果業務邏輯需要替換的時候只要改動工廠函數相關的類別
就可以替換使用不同的存儲方式
