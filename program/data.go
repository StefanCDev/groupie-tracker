package tracker
//data from json
type JsonArtists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}
type JsonLocations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
	type JsonDates struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}
	//artist to be displayed
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string `json:"locations"`
	ConcertDates []string `json:"concertDates"`
}
//instances of data
var (
	TempArtists JsonArtists
)
//api to be fetched
const (
	Artist_api = "https://groupietrackers.herokuapp.com/api/artists"
)