package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	 
	portString := os.Getenv("PORT")
	fmt.Println("port env is",portString)

}