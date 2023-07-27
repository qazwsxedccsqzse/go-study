### 初始化專案
> go mod init neilwang.com/go-study

### import 自己本地的 package
> import neilwang.com/go-study/demo

### 爲什麼我們需要框架
1. 讓部分基礎功能重複使用, 不用重複造輪
1. 定義好標準規範, 讓專案架構與權責分明
1. 好維護, 方便擴展, 方便協作開發..等  
1. [ref](https://zhuanlan.zhihu.com/p/31076196)

### golang 特性..?
1. golang body 讀完一次之後, 資料就會被清空，所以框架在處理 middleware 的時候都需要額外處理這個特性
2. 存取 request.GetBody() 的時候會發現，原生的 GetBody 會是 nil，需要各個框架自己去實作
3. request.URL 只有 Path 一定不為空值，其他不確定..等，故需要有框架處理這些問題

### 自幹簡單路由