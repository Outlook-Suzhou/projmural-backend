# API 文档

## /api/login

request

```json
{
  "data": {
    "microsoft_id": "a string"
  }
}
```

respond

```json
{
  "data": {
    "jwt": "a string"
  },
  "msg": "ok",
  "retc": 0
}
```

## /api/user

request:

```json
{
  "type": "insert/update/query",
  // 如果type是query的话只传microsoft_id否则三个字段都要传
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  },
  "jwt": "a string"
}
```

respond:

```json
{
  // 只有查询有data字段
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  },
  "msg": "ok",
  // 0 是正常 其他是错误
  "retc": 0
}
```

错误码以及msg:

0: "ok"

-1: "jwt fail"