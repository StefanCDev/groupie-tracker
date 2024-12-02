package main
import (
	"fmt"
	"log"
	"net/http"
	groupie "tracker/program"
)
func main() {
	//handle functions
	http.HandleFunc("/", groupie.HomeHandler)
	http.HandleFunc("/artists/", groupie.Get)
	//file servers to render things like css,custom font,svg's etc..
	fileserver := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))
	//web server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("temp artists: ", groupie.TempArtists)
}