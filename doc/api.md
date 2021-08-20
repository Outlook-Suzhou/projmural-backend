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

jwt 放在 header 的 Authorization 里

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

## GET /api/currentUser

jwt 放在 header 的 Authorization 里
返回jwt对应的user信息

respond:

```json
{
  "msg": "ok",
  "retc": 0,
  //也可能是user not exist 这时没有data
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  }
}
```

错误码以及msg:

const (
RESP_OK_WITH_DATA = 1
RESP_OK = 0
RESP_JWT_FAIL = -1
RESP_SERVER_ERROR = -2
RESP_ACCESS_TOKEN_FAIL = -3
RESP_USER_NOT_EXIST = -4
RESP_INVALID_OPERATION = -5
RESP_INVALID_JSON_FORMAT = -6
)

var respMsg = map[int]string {
RESP_OK: "ok",
RESP_JWT_FAIL: "jwt fail",
RESP_SERVER_ERROR: "server error",
RESP_ACCESS_TOKEN_FAIL: "access token fail",
RESP_USER_NOT_EXIST: "user not exist",
RESP_INVALID_OPERATION: "invalid operation",
RESP_INVALID_JSON_FORMAT: "invalid json format",
}