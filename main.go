package main

import (
	"Test-backend/banner"
	"fmt"
	"net/http"
)

func main() {
	bannerService := banner.NewBanner()

	http.HandleFunc("/user_banner", func(w http.ResponseWriter, r *http.Request) {
		banner.UserBannerHandler(w, r, bannerService)
	})
	http.HandleFunc("/banner", func(w http.ResponseWriter, r *http.Request) {
		banner.BannerHandler(w, r, bannerService)
	})

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
