package postmodule

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dto struct {
	Name int `json:"name"`
	Age  string
}

type PostController struct {
}

func (b PostController) HandleGetRoute(w http.ResponseWriter, r *http.Request) {
	var data Dto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("done", data)
}

func (b PostController) HandlePostRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post controller route post")
}
func (b PostController) HandlePutRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post controller route put")
}
func (b PostController) HandleDeleteRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post controller route delete")
}

func GetPostByIdPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("id")
}
