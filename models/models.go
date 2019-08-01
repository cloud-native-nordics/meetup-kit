package models

type MeetupGroup struct {
	MeetupID   string       `json:"meetupID"`
	Name       *string      `json:"name"`
	City       *string      `json:"city"`
	Country    *string      `json:"country"`
	Organizers []*Organizer `json:"organizers"`
	Meetups    []*Meetup    `json:"meetups"`
}

type Organizer struct {
	ID             string    `json:"id"`
	Name           *string   `json:"name"`
	Title          *string   `json:"title"`
	Email          *string   `json:"email"`
	Company        *Company  `json:"company"`
	Countries      []*string `json:"countries"`
	Github         *string   `json:"github"`
	Twitter        *string   `json:"twitter"`
	SpeakersBureau *string   `json:"speakersBureau"`
	MeetupGroupID  *string   `json:"-"`
}

type Company struct {
	ID          string    `json:"id"`
	Name        *string   `json:"name"`
	WebsiteURL  *string   `json:"websiteURL"`
	LogoURL     *string   `json:"logoURL"`
	Countries   []*string `json:"countries"`
	OrganizerID *string   `json:"-"`
	SpeakerID   *string   `json:"-"`
}

type Meetup struct {
	ID            int             `json:"id"`
	Name          *string         `json:"name"`
	Date          *string         `json:"date"`
	Duration      *string         `json:"duration"`
	Attendees     int             `json:"attendees"`
	Address       *string         `json:"address"`
	Sponsors      *MeetupSponsor  `json:"sponsors"`
	Presentations []*Presentation `json:"presentations"`
	MeetupGroupID *string         `json:"-"`
}

type Sponsor struct {
	ID         string    `json:"id"`
	Name       *string   `json:"name"`
	WebsiteURL *string   `json:"websiteURL"`
	LogoURL    *string   `json:"logoURL"`
	Countries  []*string `json:"countries"`
}

type Member struct {
	ID         string    `json:"id"`
	Name       *string   `json:"name"`
	WebsiteURL *string   `json:"websiteURL"`
	LogoURL    *string   `json:"logoURL"`
	Countries  []*string `json:"countries"`
}

type MeetupSponsor struct {
	ID       string
	Venue    *Sponsor   `json:"venue"`
	Other    []*Sponsor `json:"other"`
	MeetupID int        `json:"-"`
}

type Presentation struct {
	ID       string     `json:"id"`
	Duration *string    `json:"duration"`
	Title    *string    `json:"title"`
	Slides   *string    `json:"slides"`
	Speakers []*Speaker `json:"speakers"`
	MeetupID int        `json:"-"`
}

type Speaker struct {
	ID             string    `json:"id"`
	Name           *string   `json:"name"`
	Title          *string   `json:"title"`
	Email          *string   `json:"email"`
	Company        *Company  `json:"company"`
	Countries      []*string `json:"countries"`
	Github         *string   `json:"github"`
	SpeakersBureau *string   `json:"speakersBureau"`
	PresentationID *string   `json:"-"`
}

type Country struct {
	ID   *string `json:"-"`
	Name *string `json:"-"`
}

type EntityToCountry struct {
	CountryID  *string `json:"-"`
	ID         *string `json:"-"`
	EntityType string  `json:"-"`
}
