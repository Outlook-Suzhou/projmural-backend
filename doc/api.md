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

## GET /api/user/{base64编码后的microsoftID}

respond:

```json
{
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  },
  "msg": "ok",
  "retc": 0
}
```

## POST /api/user

request:

```json
{
  "type": "update/insert",
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
  "msg": "ok",
  "retc": 0
}
```

错误码以及msg:

0: "ok"

-1: "jwt fail"

-2: "server error"