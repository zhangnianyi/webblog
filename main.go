package  main

import (
	"ginblog/modles"
	"ginblog/routes"
)

func main(){
	modles.InitDb()
	routes.InitRouter()
}
