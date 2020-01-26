package models

type CompanyIn struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	WebsiteURL string `json:"websiteURL"`
	LogoURL    string `json:"logoURL"`
	WhiteLogo  bool   `json:"whiteLogo"`
}

type SpeakerIn struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Title          *string `json:"title"`
	Email          string  `json:"email"`
	Company        string  `json:"company"`
	Github         string  `json:"github"`
	Twitter        *string `json:"twitter"`
	SpeakersBureau string  `json:"speakersBureau"`
}

type MeetupGroupIn struct {
	Photo            *string              `json:"photo"`
	Name             *string              `json:"name"`
	City             *string              `json:"city"`
	Country          *string              `json:"country"`
	Description      *string              `json:"description"`
	SponsorTiers     map[string]string    `json:"sponsorTiers"`
	MeetupID         string               `json:"meetupID"`
	Organizers       []string             `json:"organizers"`
	CfpLink          string               `json:"cfpLink"`
	Latitude         float64              `json:"latitude"`
	Longitude        float64              `json:"longitude"`
	EcosystemMembers []string             `json:"ecosystemMembers"`
	Meetups          map[string]*MeetupIn `json:"meetups"`
}

type MeetupIn struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Date          string            `json:"date"`
	Duration      string            `json:"duration"`
	Attendees     int               `json:"attendees"`
	Address       string            `json:"address"`
	Photo         string            `json:"photo"`
	Recording     string            `json:"recording"`
	Sponsors      []*SponsorIn      `json:"sponsors"`
	Presentations []*PresentationIn `json:"presentations"`
}

type SponsorIn struct {
	Company string `json:"company"`
	Role    string `json:"role"`
}

type PresentationIn struct {
	Duration string    `json:"duration"`
	Title    string    `json:"title"`
	Slides   string    `json:"slides"`
	Speakers []*string `json:"speakers"`
}

type MeetupGroup struct {
	Photo       string
	Name        string
	City        string
	Country     string
	Description string
	MeetupID    string
	CfpLink     string
	Longitude   float64
	Latitude    float64
}

type Meetup struct {
	ID        int
	Name      string
	Date      string
	Duration  string
	Attendees int
	Address   string
	Photo     string
	Recording string
}

type Company struct {
	ID         string
	Name       string
	WebsiteURL string
	LogoURL    string
	WhiteLogo  bool
}

type Sponsor struct {
	ID   string
	Role string
}

type SponsorTier struct {
	ID   string
	Tier string
}

type Presentation struct {
	ID       string
	Duration string
	Title    string
	Slides   string
}

type Speaker struct {
	ID             string
	Name           string
	Title          *string
	Email          string
	Github         string
	Twitter        *string
	SpeakersBureau string
}

//Mapping Tables
type SpeakerToCompany struct {
	ID        string
	SpeakerID string
	CompanyID string
}

type SponsorTierToMeetupGroup struct {
	ID            string
	MeetupGroupID string
	SponsorTierID string
}

type SponsorTierToCompany struct {
	ID            string
	SponsorTierID string
	CompanyID     string
}

type MeetupGroupToOrganizer struct {
	ID            string
	MeetupGroupID string
	OrganizerID   string
}

type MeetupGroupToEcosystemMember struct {
	ID            string
	MeetupGroupID string
	CompanyID     string
}

type MeetupGroupToMeetup struct {
	ID            string
	MeetupGroupID string
	MeetupID      int
}

type MeetupToSponsor struct {
	ID        string
	MeetupID  int
	SponsorID string
}

type SponsorToCompany struct {
	ID        string
	SponsorID string
	CompanyID string
}

type MeetupToPresentation struct {
	ID             string
	MeetupID       int
	PresentationID string
}

type PresentationToSpeaker struct {
	ID             string
	PresentationID string
	SpeakerID      string
}
