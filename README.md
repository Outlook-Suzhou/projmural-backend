# projmural-backend

## Configuration
position: storage\env\{environment variables}\
when running in docker, change it to storage\env\docker\
### app.yml
```yaml
env: "local" #when running in docker, change it to "docker" 
name: "projmural"
port: "8081"
```

### mongodb.yml

```yaml
timeout_second: 5
connect_url: "mongodb://localhost" #when running in docker, change it to "mongodb://my_mongo" 
database_name: "projmural"
```

### jwt.yml
```yaml
secret: "your secret"
expired_seconds: 10800 # 3h
issuer: "projmural"
graph_me_endpoint: "https://graph.microsoft.com/v1.0/me" #https://docs.microsoft.com/en-us/graph/overview
admin_key: "your admin_key"
```

## Database
based on mongodb
``` go
type UserRequest struct {
	Type string   `json:"type"`// operator type update/insert/query
	Data dao.User `json:"data"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"-"`// id in mongodb
	MicrosoftId string             `bson:"microsoft_id,omitempty" json:"microsoft_id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Mail        string             `bson:"mail,omitempty" json:"mail"`
	Canvas      []CanvaInfo           `bson:"canvas,omitempty" json:"canvas"`
}

type CanvaInfo struct {
	ID string `bson:"id" json:"id"`// more information is in sharedb, you could refer projmural-frontend
	Name string `bson:"name" json:"name"`
	RecentOpen int32 `bson:"recent_open", json:"recent_open"`
}
```


## API

### POST /api/login

request

```json
{
  "access_token": "a string"
}
```

respond

```json
{
  "data": {
    "jwt": "a string"
  },
  "msg": "ok",
  // 0 true else false
  "retc": 0
}
```

### POST /api/user

Authorization：Bearer Token

JWT is the Token

request:
```json
{
  "type": "update/insert/query",
  "data": {
    //when in query, only microsoft_id is needed
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
  // only in query, the data is returned
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  }
}
```

### GET /api/currentUser
Authorization：Bearer Token

JWT is the Token

return the information of microsoft_id in JWT

respond:

```json
{
  "msg": "ok",
  "retc": 0,
  //when user is not exist, there is no data
  "data": {
    "microsoft_id": "a string",
    "name": "a string",
    "canvas": ["canvaId1", "canvaId2"]
  }
}
```

error code and msg:

```go
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
```

## run in docker
important：you should change the configuration files before running in docker.
eg:
related files position should in storage\env\docker\
env in app.yml should be "docker"
connect_url in mongodb.yml should be "mongodb://my_mongo"

```
docker pull mongo
docker run -p 27017:27017 --name my_mongo -itd mongo
docker build -t backend:v1 .
docker run -p 8081:8081 --name my_backend --link my_mongo -itd backend:v1 --env=docker
```