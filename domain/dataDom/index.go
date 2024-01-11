package DataDomUseCase

import (
	Database "bodyplate.com/internal/repo"
	mongoDatabase "bodyplate.com/internal/repo/mongo"
)
type DataDomUseCase struct {
	
}

func (d *DataDomUseCase) SendEvent() {
	
}


func  (d *DataDomUseCase ) Create(payload mongoDatabase.DomSchema)(interface{},error) {
	result,err := Database.DataBaselayer.DomRepository.Create(payload)
	go func() {
		d.SendEvent()
	}()
	if err != nil {
		return nil, err
	}
	return result, nil
}