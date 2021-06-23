package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/michaelahli/Simple-Image-Transfer/server/config/database"
	"github.com/michaelahli/Simple-Image-Transfer/server/config/utils"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/controllers"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/repositories"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/router"
	"github.com/michaelahli/Simple-Image-Transfer/server/src/usecases"
)

/*
	@load 	.env,
			mongodb
*/
func init() {
	godotenv.Load()
	database.LoadMongo()
}

func main() {
	mongodb := database.LoadDatabase()
	repository := repositories.SetupRepository(mongodb)
	usecase := usecases.SetupUsecase(repository)
	ctrl := controllers.SetupController(usecase)
	route := router.SetupRouter(ctrl)
	routes := route.RouterGroup()
	log.Printf("[SERVER] starting at port : %v", os.Getenv("SERVER_PORT"))
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			utils.AllowCORS(routes),
		),
	)
}
