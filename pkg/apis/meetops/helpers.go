package meetops

import (
	"fmt"
	"strings"
)

var (
	ValidSponsorRoles = map[SponsorRole]struct{}{
		SponsorRoleVenue:    {},
		SponsorRoleLongterm: {},
		SponsorRoleCloud:    {},
		SponsorRoleFood:     {},
		SponsorRoleOther:    {},
	}
)

func (s Speaker) String() string {
	str := s.Name
	if len(s.Github) != 0 {
		str += fmt.Sprintf(" [@%s](https://github.com/%s)", s.Github, s.Github)
	}
	if len(s.Title) != 0 {
		str += fmt.Sprintf(", %s", s.Title)
	}
	/*
		TODO: Enable again
		if s.Company.Company != nil {
			str += fmt.Sprintf(", [%s](%s)", s.Company.Name, s.Company.WebsiteURL)
		}
	*/
	if len(s.SpeakersBureau) != 0 {
		str += fmt.Sprintf(", [Contact](https://www.cncf.io/speaker/%s)", s.SpeakersBureau)
	}
	return str
}

// CityLowercase gets the lowercase variant of the city
func (mg *MeetupGroupStatus) CityLowercase() string {
	return strings.ToLower(mg.City)
}

/*func (mg *MeetupGroup) SetMeetupList() {
	marr := []Meetup{}
	for _, m := range mg.Meetups {
		marr = append(marr, m)
	}
	mg.MeetupList = MeetupList(marr)
	sort.Sort(mg.MeetupList)
}*/

func (m *MeetupStatus) DateTime() string {
	d := m.Date.UTC()
	year, month, day := d.Date()
	hour, min, _ := d.Clock()
	hour2, min2, _ := d.Add(m.Duration.Duration).Clock()
	return fmt.Sprintf("%d %s, %d at %d:%02d - %d:%02d", day, month, year, hour, min, hour2, min2)
}

/*func (p *Presentation) StartTime() string {
	return fmt.Sprintf("%d:%02d", p.Start.UTC().Hour(), p.Start.UTC().Minute())
}

func (p *Presentation) EndTime() string {
	return fmt.Sprintf("%d:%02d", p.End.UTC().Hour(), p.End.UTC().Minute())
}*/

/*func (mg *MeetupGroup) ApplyGeneratedData() {
	for key := range mg.AutoMeetups {
		m, ok := mg.Meetups[key]
		if !ok {
			found := false
			for _, date := range mg.IgnoreMeetupDates {
				if key == date {
					found = true
				}
			}
			if !found {
				log.Warnf("Didn't find information about meetup at %s on date %q\n", mg.Name, key)
			}
			continue
		}
		autoMeetup := mg.AutoMeetups[key]
		m.AutogenMeetup = &autoMeetup
		mg.Meetups[key] = m
	}
}*/
