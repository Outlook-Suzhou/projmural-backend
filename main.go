package main

import (
	"projmural-backend/dao"
)

func main() {
	mongoDao := dao.NewMongoDao()
	defer mongoDao.Close()
}
