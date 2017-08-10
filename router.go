package reuse

import "github.com/gorilla/mux"

type Router struct {
	responsible *mux.Router
}

func NewRouter() *Router {
	r := mux.NewRouter()
	return &Router{	r}
}

func (r *Router) LoadModule(app Application) bool {
	app.Container.SetInstance(`router`, r)
	return true
}