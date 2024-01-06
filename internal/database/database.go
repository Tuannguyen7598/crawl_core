package database

import (
	mongodb "bodyplate.com/internal/database/mongo"
	abtractRepository "bodyplate.com/internal/database/type"
)




type DataBase struct {
	ResHttpRepository abtractRepository.Repository
}

var DataBaselayer = DataBase{}
func (m *DataBase) Init() {
	mongoDb := mongodb.MongoDB{}
	mongoDb.Init()
	m.ResHttpRepository = mongoDb.ResHttpRepo
}

func (m *DataBase) GetRepository() abtractRepository.Repository {
	return m.ResHttpRepository
}

