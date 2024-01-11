package postmodule

import (
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
	fmt.Println("Post controller route get")
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
