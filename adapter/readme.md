# 適配器模式

適配器模式是轉換一種介面來適配另一種介面

提示:
Target 是外層的物件
Adaptee 是被包裝起來的裡層物件
把 adaptee 的實例丟到 target 裡面
Target裡面匿名組合了 Adaptee 的介面
預期是呼叫 Target 物件以後 Request() 會轉接 Adaptee 的 SpecificRequest
