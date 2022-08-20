package order

import (
	"L0/pkg/storage"
)

type App struct {
	st *storage.Storage
}

func NewApp(st *storage.Storage) *App {
	return &App{st: st}
}

/*func (a *App) ByID(w http.ResponseWriter, r *http.Request, ps *httprouter.Params) {
	orderNum := ps.ByName("orderNum")
	if orderNum == "" {
		fmt.Println("Order not found")
	}
	order, err := a.st.Order(orderNum)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, string(order))
}*/
