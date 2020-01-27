package meetops

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/weaveworks/gitops-toolkit/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindSpeaker     runtime.Kind = "Speaker"
	KindCompany     runtime.Kind = "Company"
	KindMeetup      runtime.Kind = "Meetup"
	KindMeetupGroup runtime.Kind = "MeetupGroup"
)

type CompanyID string
type SpeakerID string

/*type StatsFile struct {
	MeetupGroups uint64                 `json:"meetupGroups"`
	AllMeetups   MeetupStats            `json:"allMeetups"`
	PerMeetup    map[string]MeetupStats `json:"perMeetup"`
}

type MeetupStats struct {
	Sponsors      uint64                 `json:"sponsors"`
	SponsorByTier map[SponsorTier]uint64 `json:"sponsorByTier,omitempty"`
	Speakers      uint64                 `json:"speakers"`
	Meetups       uint64                 `json:"meetups"`
	Members       uint64                 `json:"members"`
	TotalRSVPs    uint64                 `json:"totalRSVPs"`
	AverageRSVPs  uint64                 `json:"averageRSVPs"`
	UniqueRSVPs   uint64                 `json:"uniqueRSVPs"`
}*/

type SponsorRole string

var (
	SponsorRoleVenue    SponsorRole = "Venue"
	SponsorRoleLongterm SponsorRole = "Longterm"
	SponsorRoleCloud    SponsorRole = "Cloud"
	SponsorRoleFood     SponsorRole = "Food"
	SponsorRoleOther    SponsorRole = "Other"
)

type SponsorTier string

var (
	SponsorTierLongterm        SponsorTier = "Longterm"
	SponsorTierMeetup          SponsorTier = "Meetup"
	SponsorTierSpeakerProvider SponsorTier = "SpeakerProvider"
	SponsorTierEcosystemMember SponsorTier = "EcosystemMember"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Company struct {
	// TypeMeta defines what kind of type this object is
	runtime.TypeMeta `json:",inline"`
	// ObjectMeta defines metadata like a human-readable name, and the machine-readable ID
	runtime.ObjectMeta `json:"metadata"`

	ID         CompanyID `json:"id"`
	Name       string    `json:"name"`
	WebsiteURL string    `json:"websiteURL"`
	LogoURL    string    `json:"logoURL"`
	WhiteLogo  bool      `json:"whiteLogo,omitempty"`
}

type CompanyRef struct {
	Ref *Company
	ID  CompanyID
}

func (c CompanyRef) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.ID + `"`), nil
}

func (c *CompanyRef) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		*c = CompanyRef{}
		return nil
	}
	cid := CompanyID("")
	if err := json.Unmarshal(b, &cid); err != nil {
		return fmt.Errorf("couldn't unmarshal company %q: %v", string(b), err)
	}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Speaker struct {
	// TypeMeta defines what kind of type this object is
	runtime.TypeMeta `json:",inline"`
	// ObjectMeta defines metadata like a human-readable name, and the machine-readable ID
	runtime.ObjectMeta `json:"metadata"`

	ID             SpeakerID  `json:"id"`
	Name           string     `json:"name"`
	Title          string     `json:"title,omitempty"`
	Email          string     `json:"email"`
	Company        CompanyRef `json:"company"`
	Github         string     `json:"github"`
	Twitter        string     `json:"twitter,omitempty"`
	SpeakersBureau string     `json:"speakersBureau"`
}

type SpeakerRef struct {
	Ref *Speaker
	ID  SpeakerID
}

func (s SpeakerRef) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.ID + `"`), nil
}

func (s *SpeakerRef) UnmarshalJSON(b []byte) error {
	if string(b) == "null" || string(b) == `""` {
		*s = SpeakerRef{}
		return nil
	}
	sid := SpeakerID("")
	if err := json.Unmarshal(b, &sid); err != nil {
		return fmt.Errorf("couldn't unmarshal speaker %q: %v", string(b), err)
	}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type MeetupGroup struct {
	// TypeMeta defines what kind of type this object is
	runtime.TypeMeta `json:",inline"`
	// ObjectMeta defines metadata like a human-readable name, and the machine-readable ID
	runtime.ObjectMeta `json:"metadata"`

	Spec   MeetupGroupSpec   `json:"spec"`
	Status MeetupGroupStatus `json:"status"`
}

type MeetupGroupSpec struct {
	MeetupID          string            `json:"meetupID"`
	Organizers        []SpeakerRef      `json:"organizers"`
	IgnoreMeetupDates []string          `json:"ignoreMeetupDates,omitempty"`
	CFP               string            `json:"cfpLink"`
	Latitude          float64           `json:"latitude"`
	Longitude         float64           `json:"longitude"`
	EcosystemMembers  []CompanyRef      `json:"ecosystemMembers"`
	Meetups           map[string]Meetup `json:"meetups"`
	//MeetupList MeetupList `json:"-"`
}

type MeetupGroupStatus struct {
	Photo        string                    `json:"photo,omitempty"`
	Name         string                    `json:"name"`
	City         string                    `json:"city"`
	Country      string                    `json:"country"`
	Description  string                    `json:"description"`
	SponsorTiers map[CompanyID]SponsorTier `json:"sponsorTiers"`
	//AutoMeetups  map[string]AutogenMeetup  `json:"-"`

	//Members uint64 `json:"-"`
}

// MeetupList is a slice of meetups implementing sort.Interface
type MeetupList []Meetup

var _ sort.Interface = MeetupList{}

func (ml MeetupList) Len() int {
	return len(ml)
}

func (ml MeetupList) Less(i, j int) bool {
	return ml[i].Status.Date.Time.After(ml[j].Status.Date.Time)
}

func (ml MeetupList) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Meetup struct {
	// TypeMeta defines what kind of type this object is
	runtime.TypeMeta `json:",inline"`
	// ObjectMeta defines metadata like a human-readable name, and the machine-readable ID
	runtime.ObjectMeta `json:"metadata"`

	Spec   MeetupSpec   `json:"spec"`
	Status MeetupStatus `json:"status"`
}

type MeetupSpec struct {
	Recording     string          `json:"recording"`
	Sponsors      []MeetupSponsor `json:"sponsors"`
	Presentations []Presentation  `json:"presentations"`
}

type MeetupStatus struct {
	ID        uint64          `json:"id"`
	Photo     string          `json:"photo,omitempty"`
	Name      string          `json:"name"`
	Date      metav1.Time     `json:"date,omitempty"`
	Duration  metav1.Duration `json:"duration,omitempty"`
	Attendees uint64          `json:"attendees,omitempty"`
	Address   string          `json:"address"`

	// RSVPs map the user ID to how many rsvp's they used at this event (themselves + guests)
	//RSVPs map[uint64]uint64 `json:"-"`
}

type MeetupSponsor struct {
	Role    SponsorRole `json:"role"`
	Company CompanyRef  `json:"company"`
}

type Presentation struct {
	Duration  metav1.Duration  `json:"duration"`
	Delay     *metav1.Duration `json:"delay,omitempty"`
	Title     string           `json:"title"`
	Slides    string           `json:"slides"`
	Recording string           `json:"recording,omitempty"`
	Speakers  []SpeakerRef     `json:"speakers"`

	//Start metav1.Time `json:"-"`
	//End   metav1.Time `json:"-"`
}
