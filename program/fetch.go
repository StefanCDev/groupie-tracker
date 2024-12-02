package tracker
import (
	"encoding/json"
	"log"
	"net/http"
)
func FetchApi(url string, object interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	//close the body after received the response
	defer resp.Body.Close()
	// Unmarshal the response body into a JsonArtists struct
	return json.NewDecoder(resp.Body).Decode(&object)
}