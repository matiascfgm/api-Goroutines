package main

import (
	"fmt"
	"github.com/matiascfgm/api-Goroutines/controller"
	router "github.com/matiascfgm/api-Goroutines/http"
	"github.com/matiascfgm/api-Goroutines/service"
	"net/http"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)

	httpRouter.SERVE(port)
}
