# 工廠方法模式 

OperatorBase 只用來定義 a 和 b 和 SetA 和 SetB
但 Operator 的 interface 應該要包含 SetA / SetB / Result
此時建立 PlusOperator (裡面包著OperatorBase，即用來擴充OperatorBase)
以及建立 MinusOperator (裡面包著OperatorBase，即用來擴充OperatorBase)
並各自實作 Result()，此時 PlusOperator 和 MinusOperator 滿足 Operator 介面

PlusOperatorFactory 和 MinusOperatorFactory 都實作 Create()
並傳出 PlusOperator 和 MinusOperator 的 struct

提示:
Golang中沒有繼承，所以用匿名組合來實現
使用子類別的方式延遲生成對象到子類別中實現
