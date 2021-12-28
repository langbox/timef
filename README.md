# timef

## 使用说明

### 引用包

```
go get github.com/langbox/timef
```

### 时间戳转字符串

```
var timestamp int64 = 1640682392
str := timef.GetNormalStringByTimestamp(timestamp)
fmt.Println(str)
```