package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cloud-native-nordics/meetup-kit/pkg/graphql/models"
)

// GetMeetupInfoFromAPI fetches all information it can about the given meetup group
// from the meetup.com API, and returns the autogenerated type
func GetMeetupInfoFromAPI(meetupGroup models.MeetupGroup) (*meetupGroupAPI, error) {
	mg := &meetupGroupAPI{}
	if err := fetchMeetupGroup(meetupGroup.MeetupID, mg); err != nil {
		return nil, err
	}

	return mg, nil
}

func fetchMeetupGroup(meetupGroupID string, mg *meetupGroupAPI) error {
	url := fmt.Sprintf("https://api.meetup.com/%s", meetupGroupID)
	return getJSON(url, mg)
}

type meetupGroupAPI struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	City        string `json:"untranslated_city"`
	Country     string `json:"localized_country_name"`
	Members     int    `json:"members"`
	Photo       struct {
		Link string `json:"highres_link"`
	} `json:"key_photo"`
}

func getJSON(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("GetJSON failed for url %s with error %v", url, err)
	}
	return nil
}
