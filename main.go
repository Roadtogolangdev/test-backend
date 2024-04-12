package main

import (
	"Test-backend/banner"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	bannerService := banner.NewBannerService()

	router.GET("/user_banner", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		banner.UserBannerHandler(w, r, bannerService)
	})

	router.GET("/banner", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		banner.BannerHandler(w, r, bannerService)
	})

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
