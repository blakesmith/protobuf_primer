package main

import (
	"code.google.com/p/goprotobuf/proto"
	"log"
	"./messages"
	"net/http"
)

func userShow(w http.ResponseWriter, r *http.Request) {
	bob := &messages.User{
		FirstName: proto.String("Bob"),
		LastName: proto.String("Dole"),
		PostCount: proto.Int64(10),
	}
	joe := &messages.User{
		FirstName: proto.String("Joe"),
		LastName: proto.String("Developer"),
		PostCount: proto.Int64(4),
		Friends: []*messages.User{bob},
	}

	b, err := proto.Marshal(joe)
	if err != nil {
		http.Error(w, "Failed to Marshal joe!", 500)

		return
	}
	w.Header().Set("Content-type", "application/x-protobuf")
	w.Write(b)
}

func main() {
	http.HandleFunc("/users/joe", userShow)

	log.Printf("Listening for requests...")
	http.ListenAndServe(":5555", nil)
}