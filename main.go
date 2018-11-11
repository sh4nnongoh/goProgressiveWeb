package main

import (
	//"net/http"
	"log"
	"net/http"

	"github.com/xeoncross/secureserver"
	//"github.com/gorilla/mux"
)

func main() {
	// Here we are instantiating the gorilla/mux router
	//r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	//r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Our application will run on port 3000. Here we declare the port and pass in our router.
	//http.ListenAndServe(":3000", r)

	domain := "127.0.0.1"
	secureserver.RunHTTPRedirectServer()
	s := secureserver.GetHTTPSServer(domain)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server on " + domain + ".\n"))
	})

	s.Handler = mux

	log.Fatal(s.ListenAndServeTLS("", ""))
}
