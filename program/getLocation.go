package tracker
import (
	"encoding/json"
	"log"
	"net/http"
)
func GetArtistLocation(url string) []string {
	//instance of json locations
	var locations JsonLocations
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		log.Fatal(err)
	}
	return locations.Locations
}
