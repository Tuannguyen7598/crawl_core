package DatabaseAbtract

type Repository interface{
	Create(obj interface{}) (interface{}, error)
	Find(con interface{}) ([]interface{}, error)
	FindOne(con interface{}) (interface{}, error)
	Update(con interface{},obj interface{}) (error)
	Delete(con interface{}) (error)
}
