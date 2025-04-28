# MAT Services Portal
- 建立單機版 MAT Services 
- 用來協助與 Frontend 開發討論之範例

## Running the Project
```bash
go run main.go
OR
go run -mod=mod main.go
```

## Tool 
- github.com/gin-gonic/gin
- go.uber.org/zap
- go.uber.org/zap/zapcore

## Func
- logger
- middleware
- swagger

## swagger 文件網頁
- http://localhost:8080/swagger/index.html

### param 參數
```json
{
  "OVC_START_DATE_TIME": "2025/01/01 00:00:00",
  "OVC_END_DATE_TIME": "2025/01/02 00:00:00",
  "PairType": "2N3N"
}
```

