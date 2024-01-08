package Database

import (
	DatabaseMongo "bodyplate.com/internal/repo/mongo"
	DatabaseAbtract "bodyplate.com/internal/repo/type"
)




type DataBaselayerStruct struct {
	ResHttpRepository DatabaseAbtract.Repository
	DomRepository DatabaseAbtract.Repository
}

var DataBaselayer = DataBaselayerStruct{}
func (m *DataBaselayerStruct) Init() {
	mongoDb := DatabaseMongo.MongoDB{}
	mongoDb.Init()
	m.ResHttpRepository = mongoDb.ResHttpRepo
	m.DomRepository = mongoDb.DomRepo
}



