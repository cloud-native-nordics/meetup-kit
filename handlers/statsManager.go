package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"io/ioutil"

	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/golang/glog"
	"github.com/hashicorp/go-memdb"
)

//StatsManager is responsible for managing the stats.json file generated in the meetups repository
type StatsManager struct {
	filepath string
	URL      string
	db       *memdb.MemDB
}

type unmarshalledData struct {
	meetupGroups   []models.MeetupGroup
	organizers     []models.Organizer
	companies      []models.Company
	meetups        []models.Meetup
	sponsors       []models.Sponsor
	members        []models.Member
	meetupSponsors []models.MeetupSponsor
	// others          []models.Other
	// venues          []models.Venue
	speakers        []models.Speaker
	presentations   []models.Presentation
	countries       []models.Country
	entityToCountry []models.EntityToCountry
}

type jsonStructure struct {
	MeetupGroups []models.MeetupGroup `json:"meetupGroups"`
	Sponsors     []models.Sponsor     `json:"sponsors"`
	Members      []models.Member      `json:"members"`
}

//NewStatsManager fetches the stats.json file, marshals to structs, creates in-mem db
//and then returns a reference to the in-mem db.
func NewStatsManager(URL string) *memdb.MemDB {
	sm := &StatsManager{
		URL:      URL,
		filepath: "./data/stats.json",
	}
	schema := createDatabaseSchema()

	err := sm.fetchStats()

	if err != nil {
		glog.Fatalf("Fatal could not fetch stats.json: %s", err)
	}

	data, err := sm.unmarshallData()

	if err != nil {
		glog.Fatalf("Fatal could not unmarshall stats.json: %s", err)
	}

	db, err := sm.populateDatabase(schema, data)

	if err != nil {
		glog.Fatalf("Fatal could not populate database: %s", err)
	}

	return db
}

//fetchStats downloads the stats.json from the supplied URL to a local file.
func (sm *StatsManager) fetchStats() error {
	glog.V(5).Infof("Fetching stats.json from: %s", sm.URL)
	// Get the data
	resp, err := http.Get(sm.URL)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	defer resp.Body.Close()

	// Create the dir if not present
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "data")
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		glog.V(5).Infof("Creating data directory in %s", filesDir)
		err = os.Mkdir(filesDir, os.ModePerm)
		if err != nil {
			glog.V(3).Infof("Couldn't create a new directory. %s", err.Error())
			return err
		}
	}

	// Create the file
	out, err := os.Create(sm.filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	glog.V(5).Infof("Writing stats.json to: %s", sm.filepath)
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

//Create models from json file
func (sm *StatsManager) unmarshallData() (*unmarshalledData, error) {
	var data jsonStructure
	output := &unmarshalledData{}

	file, err := ioutil.ReadFile(sm.filepath)

	if err != nil {
		return nil, err
	}

	glog.V(5).Info("Unmarhsalling stats.json")
	err = json.Unmarshal(file, &data)

	if err != nil {
		return nil, err
	}

	sm.generateCountries(output)
	sm.generateSponsors(output, data.Sponsors)
	sm.generateMembers(output, data.Members)
	sm.generateMeetupGroups(output, data.MeetupGroups)

	return output, nil
}

func (sm *StatsManager) generateCountries(output *unmarshalledData) {
	var denmark = "denmark"
	var finland = "finland"
	var sweden = "sweden"

	countries := []models.Country{
		models.Country{&denmark, &denmark},
		models.Country{&finland, &finland},
		models.Country{&sweden, &sweden},
	}

	output.countries = countries
}

func (sm *StatsManager) generateCountryRelation(output *unmarshalledData, countries []*string, id string, entityType string) {
	for _, country := range countries {
		output.entityToCountry = append(output.entityToCountry,
			models.EntityToCountry{
				ID:         &id,
				CountryID:  country,
				EntityType: entityType,
			})
	}

}

func (sm *StatsManager) generateCompany(output *unmarshalledData, company *models.Company) {
	if company != nil {
		output.companies = append(output.companies, *company)
		sm.generateCountryRelation(output, company.Countries, company.ID, "company")
	}
}

func (sm *StatsManager) generateOrganizers(output *unmarshalledData, organizers []*models.Organizer, meetupID string) {
	for _, organizer := range organizers {
		organizerToBe := organizer
		organizerToBe.MeetupGroupID = &meetupID
		output.organizers = append(output.organizers, *organizerToBe)

		sm.generateCountryRelation(output, organizer.Countries, organizer.ID, "organizer")
		sm.generateCompany(output, organizer.Company)
	}
}

func (sm *StatsManager) generateSponsors(output *unmarshalledData, sponsors []models.Sponsor) {
	if sponsors != nil {
		for _, sponsor := range sponsors {
			sm.generateCountryRelation(output, sponsor.Countries, sponsor.ID, "sponsor")
			output.sponsors = append(output.sponsors, sponsor)

		}
	}
}

func (sm *StatsManager) generateMembers(output *unmarshalledData, members []models.Member) {
	if members != nil {
		for _, member := range members {
			sm.generateCountryRelation(output, member.Countries, member.ID, "member")
			output.members = append(output.members, member)

		}
	}
}

func (sm *StatsManager) generateMeetupSponsors(output *unmarshalledData, sponsors *models.MeetupSponsor, meetupID int) {
	if sponsors != nil {
		sponsorsToBe := sponsors
		sponsorsToBe.ID = strconv.Itoa(rand.Int())
		sponsorsToBe.MeetupID = meetupID

		//Add venue sponsor
		// if sponsorsToBe.Venue != nil {
		// sponsorsToBe.Venue.SponsorID = sponsorsToBe.ID
		// output.venues = append(output.venues, *sponsorsToBe.Venue)

		// sm.generateCountryRelation(output, sponsorsToBe.Venue.Countries, sponsorsToBe.Venue.ID, "venue")

		// }

		//Add other sponsors
		// for _, other := range sponsorsToBe.Other {
		// other.SponsorID = other.ID
		// output.others = append(output.others, *other)

		// sm.generateCountryRelation(output, other.Countries, other.ID, "other")

		// }

		output.meetupSponsors = append(output.meetupSponsors, *sponsorsToBe)
	}
}

func (sm *StatsManager) generatePresentations(output *unmarshalledData, presentations []*models.Presentation, meetupID int) {
	for _, presentation := range presentations {
		presentationToBe := presentation
		presentationToBe.MeetupID = meetupID
		presentationToBe.ID = strconv.Itoa(rand.Int())

		output.presentations = append(output.presentations, *presentationToBe)

		for _, speaker := range presentation.Speakers {
			speakerToBe := speaker
			speakerToBe.PresentationID = &presentation.ID

			output.speakers = append(output.speakers, *speakerToBe)

			sm.generateCountryRelation(output, speaker.Countries, speaker.ID, "speaker")

			if speaker.Company != nil {
				output.companies = append(output.companies, *speaker.Company)

				sm.generateCountryRelation(output, speaker.Company.Countries, speaker.Company.ID, "company")

			}
		}
	}
}

func (sm *StatsManager) generateMeetups(output *unmarshalledData, meetups []*models.Meetup, meetupGroupID string) {
	for _, meetup := range meetups {
		meetupToBe := meetup
		meetupToBe.MeetupGroupID = &meetupGroupID
		output.meetups = append(output.meetups, *meetupToBe)

		//Add sponsors
		sm.generateMeetupSponsors(output, meetup.Sponsors, meetupToBe.ID)

		//Add presentations
		sm.generatePresentations(output, meetup.Presentations, meetup.ID)
	}
}

func (sm *StatsManager) generateMeetupGroups(output *unmarshalledData, meetupGroups []models.MeetupGroup) {
	output.meetupGroups = meetupGroups
	for _, meetupGroup := range meetupGroups {
		//Add organizers
		sm.generateOrganizers(output, meetupGroup.Organizers, meetupGroup.MeetupID)

		//Add meetups
		sm.generateMeetups(output, meetupGroup.Meetups, meetupGroup.MeetupID)
	}
}

//Create and populate database
func (sm *StatsManager) populateDatabase(schema *memdb.DBSchema, data *unmarshalledData) (*memdb.MemDB, error) {
	// Create a new data base
	glog.V(5).Info("Creating database")
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert Countries
	glog.V(5).Infof("Inserting %d Countries", len(data.countries))
	for _, mg := range data.countries {
		if err := txn.Insert("country", mg); err != nil {
			return nil, err
		}
	}

	// Insert MeetupGroups
	glog.V(5).Infof("Inserting %d Meetup Groups", len(data.meetupGroups))
	for _, mg := range data.meetupGroups {
		if err := txn.Insert("meetupGroup", mg); err != nil {
			return nil, err
		}
	}

	// Insert Companies
	glog.V(5).Infof("Inserting %d Companies", len(data.companies))
	for _, company := range data.companies {
		if err := txn.Insert("company", company); err != nil {
			return nil, err
		}
	}

	// Insert Organizers
	glog.V(5).Infof("Inserting %d Organizers", len(data.organizers))
	for _, organizer := range data.organizers {
		if err := txn.Insert("organizer", organizer); err != nil {
			return nil, err
		}
	}

	// Insert Meetups
	glog.V(5).Infof("Inserting %d Meetups", len(data.meetups))
	for _, meetup := range data.meetups {
		if err := txn.Insert("meetup", meetup); err != nil {
			return nil, err
		}
	}

	// Insert MeetupSponsors
	glog.V(5).Infof("Inserting %d Meetup Sponsors", len(data.meetupSponsors))
	for _, sponsor := range data.meetupSponsors {
		if err := txn.Insert("meetupSponsor", sponsor); err != nil {
			return nil, err
		}
	}

	// Insert Sponsors
	glog.V(5).Infof("Inserting %d Sponsors", len(data.sponsors))
	for _, sponsor := range data.sponsors {
		if err := txn.Insert("sponsor", sponsor); err != nil {
			return nil, err
		}
	}

	// Insert Members
	glog.V(5).Infof("Inserting %d Members", len(data.members))
	for _, member := range data.members {
		if err := txn.Insert("member", member); err != nil {
			return nil, err
		}
	}

	// Insert Presentations
	glog.V(5).Infof("Inserting %d Presentations", len(data.presentations))
	for _, presentation := range data.presentations {
		if err := txn.Insert("presentation", presentation); err != nil {
			return nil, err
		}
	}

	// Insert Speakers
	glog.V(5).Infof("Inserting %d Speakers", len(data.speakers))
	for _, speaker := range data.speakers {
		if err := txn.Insert("speaker", speaker); err != nil {
			return nil, err
		}
	}

	// Insert EntityToCountry
	glog.V(5).Infof("Inserting %d EntityCountryRelations", len(data.entityToCountry))
	for _, entityToCountry := range data.entityToCountry {
		if err := txn.Insert("entityToCountry", entityToCountry); err != nil {
			return nil, err
		}
	}

	// Commit the transaction
	txn.Commit()

	return db, nil
}
