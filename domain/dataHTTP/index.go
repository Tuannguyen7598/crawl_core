package DataHTTPUseCase

import (
	Database "bodyplate.com/internal/repo"
	mongoDatabase "bodyplate.com/internal/repo/mongo"
)
type DataHttpUseCase struct {
	
}

func (d *DataHttpUseCase) SendEvent() {
	
}


func  (d *DataHttpUseCase ) Create(payload mongoDatabase.ResponseHttpSchema)(interface{},error) {
	result,err := Database.DataBaselayer.ResHttpRepository.Create(payload)
	go func() {
		d.SendEvent()
	}()
	if err != nil {
		return nil, err
	}
	return result, nil
}