package web

import (
	"fmt"
	"net/http"

	"github.com/shaninalex/financial-analyzer/pkg/router"
)

func Websocket(port int) {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("X-User")
		if userID == "" {
			http.Error(w, "user id is empty", http.StatusUnauthorized)
			return
		}
		router.ServeWebsocket(userID, w, r)
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
