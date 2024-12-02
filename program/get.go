package tracker
import (
	"net/http"
)
func Get(w http.ResponseWriter, r *http.Request) {
	//check the url
	if r.URL.Path != "/artists/" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	//parse the form and get the value of the input text
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	userInput := r.Form.Get("search")
	//vaildate user input
	if isValid(userInput) {
		artist := new(Artist)                                            //define *instance of artist to be displayed
		newArtist, err := artist.fromJsonArtists(TempArtists, userInput) //receives json artists as argument, returns []newArtist
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		tmpl.ExecuteTemplate(w, "artists.html", newArtist)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
}