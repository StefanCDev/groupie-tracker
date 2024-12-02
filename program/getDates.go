package tracker
import (
	"encoding/json"
	"log"
	"net/http"
)
func GetArtistDates(url string) []string {
	var dates JsonDates
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		log.Fatal(err)
	}
	return dates.Dates
}
