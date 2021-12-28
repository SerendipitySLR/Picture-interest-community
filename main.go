package main

import "Picture-interest-community/router"

func main(){
	r := router.InitRouter()
	r.Run(":9999")
}