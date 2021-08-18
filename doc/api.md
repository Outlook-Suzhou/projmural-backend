# API 文档

## POST /api/login

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
  // 0 是正常 其他是错误
  "retc": 0
}
```

## POST /api/user

jwt 放在 header 的 authorization 里

request:

```json
{
  "type": "update/insert/query",
  "data": {
    // query的时候只需要填microsoft_id
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  }
}
```

respond:

```json
{
  "msg": "ok",
  "retc": 0,
  // data只有query的时候有
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  }
}
```

错误码以及msg:

0: "ok"

-1: "jwt fail"

-2: "server error"