package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//SlackRepository handles slack invites
type SlackRepository struct {
	slackToken     string
	slackURL       string
	slackCommunity string
}

//NewSlackRepository returns a new SlackRepository Singleton
func NewSlackRepository(slackToken string, slackURL string, slackCommunity string) *SlackRepository {
	return &SlackRepository{slackToken: slackToken, slackURL: slackURL, slackCommunity: slackCommunity}
}

//DoInvite sends an email invitation to the submitted email
func (sr *SlackRepository) DoInvite(email string) string {
	slackURL := fmt.Sprintf("%s/api/users.admin.invite", sr.slackURL)

	values := url.Values{
		"email":      {email},
		"token":      {sr.slackToken},
		"set_active": {"true"},
	}

	res, err := http.PostForm(slackURL, values)

	if err != nil {
		return "Something has gone wrong. Please contact a system administrator."
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return "Something has gone wrong. Please contact a system administrator."
	}

	if data["ok"] == false {
		if data["error"].(string) == "already_invited" || data["error"].(string) == "already_in_team" {
			return fmt.Sprintf("Success! You were already invited.<br>"+
				"Visit <a href='https://%s'>%s</a>", sr.slackURL, sr.slackCommunity)
		} else if data["error"].(string) == "invalid_email" {
			return "The email you entered is an invalid email."
		} else if data["error"].(string) == "invalid_auth" {
			return "Something has gone wrong. Please contact a system administrator."
		}
		return "Something has gone wrong. Please contact a system administrator."
	}

	return fmt.Sprintf("Success! Check “%s“ for an invite from Slack.", email)
}
