package tracker
import (
	"errors"
	"strings"
)
func (a *Artist) fromJsonArtists(ja JsonArtists, name string) (Artist, error) {
	//define an instance of the artist to be displayed
	var artist Artist
	var err = errors.New("Artist with name '" + name + "' not found")
	//loop over the json artists struct and create a new artist object
	for _, eachArtist := range ja {
		if strings.EqualFold(eachArtist.Name, name) { //EqualFold compares two strings in a case-insensitive manner.
			//new artist with assigned values from json artists
			artist = Artist{
				ID:           eachArtist.ID,
				Image:        eachArtist.Image,
				Name:         eachArtist.Name,
				Members:      eachArtist.Members,
				CreationDate: eachArtist.CreationDate,
				FirstAlbum:   eachArtist.FirstAlbum,
				Locations:    GetArtistLocation(eachArtist.Locations),
				ConcertDates: GetArtistDates(eachArtist.ConcertDates),
			}
			return artist, nil
		}
	}
	return artist, err
}