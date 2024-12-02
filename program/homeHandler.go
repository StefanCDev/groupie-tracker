package tracker
import (
	"html/template"
	"net/http"
)
var tmpl *template.Template
func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// Fetching the artist api and returning the json artists.
	err := FetchApi(Artist_api, &TempArtists)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
