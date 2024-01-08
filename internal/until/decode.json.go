package until

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func DecodeJSON(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func ConvertEmptyToNil(ptr interface{}) {
	// Kiểm tra xem ptr có phải là con trỏ không
	val := reflect.ValueOf(ptr)
	if val.Kind() != reflect.Ptr {
		fmt.Println("Not a pointer.")
		return
	}

	// Lấy giá trị thực của con trỏ
	val = val.Elem()

	// Kiểm tra từng trường của giá trị thực
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			// Nếu là trường chuỗi và giá trị là rỗng, gán giá trị nil
			if field.String() == "" {
				field.Set(reflect.Zero(field.Type()))
			}
			// Các kiểu dữ liệu khác có thể được xử lý tương tự nếu cần
		}
	}
}
