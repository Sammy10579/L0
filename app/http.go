package app

import (
	"log"
	"net/http"
)

func runHttpServer(listenedAddr string, service OrderService) {
	http.HandleFunc("/orders", func(writer http.ResponseWriter, request *http.Request) {
		uids := request.URL.Query()["uid"]
		if len(uids) != 1 || uids[0] == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("invalid GET argument uid"))
			return
		}

		orderPayload, err := service.PayloadByUid(request.Context(), uids[0])
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("can't get order from storage"))
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(orderPayload)
	})

	if err := http.ListenAndServe(listenedAddr, nil); err != nil {
		log.Fatal(err)
	}
}
