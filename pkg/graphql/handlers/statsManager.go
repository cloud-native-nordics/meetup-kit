package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"io/ioutil"

	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

//StatsManager is responsible for managing the stats.json file generated in the meetups repository
type StatsManager struct {
	filepath string
	URL      string
	db       *memdb.MemDB
}

type unmarshalledData struct {
	companies                    []models.Company
	speakers                     []models.Speaker
	speakerToCompany             []models.SpeakerToCompany
	meetupGroups                 []models.MeetupGroup
	sponsorTiers                 []models.SponsorTier
	sponsorTierToMeetupGroup     []models.SponsorTierToMeetupGroup
	sponsorTierToCompany         []models.SponsorTierToCompany
	meetupGroupToOrganizer       []models.MeetupGroupToOrganizer
	meetupGroupToEcosystemMember []models.MeetupGroupToEcosystemMember
	meetups                      []models.Meetup
	meetupGroupToMeetup          []models.MeetupGroupToMeetup
	sponsors                     []models.Sponsor
	meetupToSponsor              []models.MeetupToSponsor
	sponsorToCompany             []models.SponsorToCompany
	presentations                []models.Presentation
	meetupToPresentation         []models.MeetupToPresentation
	presentationToSpeaker        []models.PresentationToSpeaker
}

type jsonStructure struct {
	Companies    []models.CompanyIn     `json:"companies"`
	Speakers     []models.SpeakerIn     `json:"speakers"`
	MeetupGroups []models.MeetupGroupIn `json:"meetupGroups"`
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

	sm.generateCompanies(output, data.Companies)
	sm.generateSpeakers(output, data.Speakers)

	sm.generateMeetupGroups(output, data.MeetupGroups)

	return output, nil
}

func (sm *StatsManager) generateCompanies(output *unmarshalledData, companies []models.CompanyIn) {
	for _, company := range companies {
		newCompany := &models.Company{
			ID:         company.ID,
			Name:       company.Name,
			WebsiteURL: company.WebsiteURL,
			LogoURL:    company.LogoURL,
			WhiteLogo:  company.WhiteLogo,
		}
		output.companies = append(output.companies, *newCompany)
	}
}

func (sm *StatsManager) generateSpeakers(output *unmarshalledData, speakers []models.SpeakerIn) {
	for _, speaker := range speakers {
		newSpeaker := &models.Speaker{
			ID:             speaker.ID,
			Name:           speaker.Name,
			Title:          speaker.Title,
			Email:          speaker.Email,
			Github:         speaker.Github,
			Twitter:        speaker.Twitter,
			SpeakersBureau: speaker.SpeakersBureau,
		}
		output.speakers = append(output.speakers, *newSpeaker)
		if speaker.Company != "" {
			speakerToCompany := &models.SpeakerToCompany{ID: uuid.New().String(), SpeakerID: newSpeaker.ID, CompanyID: speaker.Company}
			output.speakerToCompany = append(output.speakerToCompany, *speakerToCompany)
		}
	}
}

func (sm *StatsManager) generateMeetupGroups(output *unmarshalledData, meetupGroups []models.MeetupGroupIn) {
	for _, group := range meetupGroups {
		newMeetupGroup := &models.MeetupGroup{
			Photo:       *group.Photo,
			Name:        *group.Name,
			City:        *group.City,
			Country:     *group.Country,
			Description: *group.Description,
			MeetupID:    group.MeetupID,
			CfpLink:     group.CfpLink,
			Latitude:    group.Latitude,
			Longitude:   group.Longitude,
		}
		output.meetupGroups = append(output.meetupGroups, *newMeetupGroup)
		sm.generateSponsorTiers(output, group.SponsorTiers, group.MeetupID)
		sm.generateOrganizers(output, group.Organizers, group.MeetupID)
		sm.generateEcosystemMembers(output, group.EcosystemMembers, group.MeetupID)
		sm.generateMeetups(output, group.Meetups, group.MeetupID)
	}
}

func (sm *StatsManager) generateSponsorTiers(output *unmarshalledData, sponsorTiers map[string]string, meetupGroupID string) {
	for company, tier := range sponsorTiers {
		newSponsorTier := &models.SponsorTier{
			ID:   uuid.New().String(),
			Tier: tier,
		}
		output.sponsorTiers = append(output.sponsorTiers, *newSponsorTier)
		newSponsorTierToMeetupGroup := &models.SponsorTierToMeetupGroup{
			ID:            uuid.New().String(),
			MeetupGroupID: meetupGroupID,
			SponsorTierID: newSponsorTier.ID,
		}
		output.sponsorTierToMeetupGroup = append(output.sponsorTierToMeetupGroup, *newSponsorTierToMeetupGroup)
		newSponsorTierToCompany := &models.SponsorTierToCompany{
			ID:            uuid.New().String(),
			SponsorTierID: newSponsorTier.ID,
			CompanyID:     company,
		}
		output.sponsorTierToCompany = append(output.sponsorTierToCompany, *newSponsorTierToCompany)
	}
}

func (sm *StatsManager) generateOrganizers(output *unmarshalledData, organizers []string, meetupGroupID string) {
	for _, organizer := range organizers {
		newMeetupGroupToOrganizer := &models.MeetupGroupToOrganizer{
			ID:            uuid.New().String(),
			MeetupGroupID: meetupGroupID,
			OrganizerID:   organizer,
		}
		output.meetupGroupToOrganizer = append(output.meetupGroupToOrganizer, *newMeetupGroupToOrganizer)
	}
}

func (sm *StatsManager) generateEcosystemMembers(output *unmarshalledData, members []string, meetupGroupID string) {
	for _, member := range members {
		newMeetupGroupToEcosystemMember := &models.MeetupGroupToEcosystemMember{
			ID:            uuid.New().String(),
			MeetupGroupID: meetupGroupID,
			CompanyID:     member,
		}
		output.meetupGroupToEcosystemMember = append(output.meetupGroupToEcosystemMember, *newMeetupGroupToEcosystemMember)
	}
}

func (sm *StatsManager) generateMeetups(output *unmarshalledData, meetups map[string]*models.MeetupIn, meetupGroupID string) {
	for _, meetup := range meetups {
		newMeetup := &models.Meetup{
			ID:        meetup.ID,
			Name:      meetup.Name,
			Date:      meetup.Date,
			Duration:  meetup.Duration,
			Photo:     meetup.Photo,
			Attendees: meetup.Attendees,
			Address:   meetup.Address,
			Recording: meetup.Recording,
		}
		output.meetups = append(output.meetups, *newMeetup)
		newMeetupGroupToMeetup := &models.MeetupGroupToMeetup{
			ID:            uuid.New().String(),
			MeetupGroupID: meetupGroupID,
			MeetupID:      newMeetup.ID,
		}
		output.meetupGroupToMeetup = append(output.meetupGroupToMeetup, *newMeetupGroupToMeetup)

		sm.generateSponsors(output, meetup.Sponsors, meetup.ID)
		sm.generatePresentations(output, meetup.Presentations, meetup.ID)
	}
}

func (sm *StatsManager) generateSponsors(output *unmarshalledData, sponsors []*models.SponsorIn, meetupID int) {
	for _, sponsor := range sponsors {
		newSponsor := &models.Sponsor{
			ID:   uuid.New().String(),
			Role: sponsor.Role,
		}
		output.sponsors = append(output.sponsors, *newSponsor)
		newMeetupToSponsor := &models.MeetupToSponsor{
			ID:        uuid.New().String(),
			SponsorID: newSponsor.ID,
			MeetupID:  meetupID,
		}
		output.meetupToSponsor = append(output.meetupToSponsor, *newMeetupToSponsor)
		newSponsorToCompany := &models.SponsorToCompany{
			ID:        uuid.New().String(),
			SponsorID: newSponsor.ID,
			CompanyID: sponsor.Company,
		}
		output.sponsorToCompany = append(output.sponsorToCompany, *newSponsorToCompany)
	}
}

func (sm *StatsManager) generatePresentations(output *unmarshalledData, presentations []*models.PresentationIn, meetupID int) {
	for _, presentation := range presentations {
		newPresentation := &models.Presentation{
			ID:       uuid.New().String(),
			Duration: presentation.Duration,
			Title:    presentation.Title,
			Slides:   presentation.Slides,
		}
		output.presentations = append(output.presentations, *newPresentation)
		newMeetupToPresentation := &models.MeetupToPresentation{
			ID:             uuid.New().String(),
			PresentationID: newPresentation.ID,
			MeetupID:       meetupID,
		}
		output.meetupToPresentation = append(output.meetupToPresentation, *newMeetupToPresentation)

		for _, speaker := range presentation.Speakers {
			newPresentationToSpeaker := &models.PresentationToSpeaker{
				ID:             uuid.New().String(),
				PresentationID: newPresentation.ID,
				SpeakerID:      *speaker,
			}
			output.presentationToSpeaker = append(output.presentationToSpeaker, *newPresentationToSpeaker)
		}
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

	// Insert Companies
	glog.V(5).Infof("Inserting %d Companies", len(data.companies))
	for _, company := range data.companies {
		if err := txn.Insert("company", company); err != nil {
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

	// Insert SpeakerToCompany
	glog.V(5).Infof("Inserting %d SpeakerToCompany Relations", len(data.speakerToCompany))
	for _, relation := range data.speakerToCompany {
		if err := txn.Insert("speakerToCompany", relation); err != nil {
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

	// Insert SponsorTiers
	glog.V(5).Infof("Inserting %d Sponsor Tiers", len(data.sponsorTiers))
	for _, sponsorTier := range data.sponsorTiers {
		if err := txn.Insert("sponsorTier", sponsorTier); err != nil {
			return nil, err
		}
	}

	// Insert SponsorTierToMeetupGroup
	glog.V(5).Infof("Inserting %d SponsorTierToMeetupGroup Relations", len(data.sponsorTierToMeetupGroup))
	for _, relation := range data.sponsorTierToMeetupGroup {
		if err := txn.Insert("sponsorTierToMeetupGroup", relation); err != nil {
			return nil, err
		}
	}

	// Insert SponsorTierToCompany
	glog.V(5).Infof("Inserting %d SponsorTierToCompany Relations", len(data.sponsorTierToCompany))
	for _, relation := range data.sponsorTierToCompany {
		if err := txn.Insert("sponsorTierToCompany", relation); err != nil {
			return nil, err
		}
	}

	// Insert SponsorToCompany
	glog.V(5).Infof("Inserting %d SponsorToCompany Relations", len(data.sponsorToCompany))
	for _, relation := range data.sponsorToCompany {
		if err := txn.Insert("sponsorToCompany", relation); err != nil {
			return nil, err
		}
	}

	// Insert MeetupGroupToOrganizer
	glog.V(5).Infof("Inserting %d MeetupGroupToOrganizer Relations", len(data.meetupGroupToOrganizer))
	for _, relation := range data.meetupGroupToOrganizer {
		if err := txn.Insert("meetupGroupToOrganizer", relation); err != nil {
			return nil, err
		}
	}

	// Insert MeetupGroupToEcosystemMember
	glog.V(5).Infof("Inserting %d MeetupGroupToEcosystemMember Relations", len(data.meetupGroupToEcosystemMember))
	for _, relation := range data.meetupGroupToEcosystemMember {
		if err := txn.Insert("meetupGroupToEcosystemMember", relation); err != nil {
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

	// Insert MeetupGroupToMeetup
	glog.V(5).Infof("Inserting %d MeetupGroupToMeetup Relations", len(data.meetupGroupToMeetup))
	for _, relation := range data.meetupGroupToMeetup {
		if err := txn.Insert("meetupGroupToMeetup", relation); err != nil {
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

	// Insert MeetupToSponsor
	glog.V(5).Infof("Inserting %d MeetupToSponsor Relations", len(data.meetupToSponsor))
	for _, relation := range data.meetupToSponsor {
		if err := txn.Insert("meetupToSponsor", relation); err != nil {
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

	// Insert MeetupToPresentation
	glog.V(5).Infof("Inserting %d MeetupToPresentation Relations", len(data.meetupToPresentation))
	for _, relation := range data.meetupToPresentation {
		if err := txn.Insert("meetupToPresentation", relation); err != nil {
			return nil, err
		}
	}

	// Insert PresentationToSpeaker
	glog.V(5).Infof("Inserting %d PresentationToSpeaker Relations", len(data.presentationToSpeaker))
	for _, relation := range data.presentationToSpeaker {
		if err := txn.Insert("presentationToSpeaker", relation); err != nil {
			return nil, err
		}
	}
	// Commit the transaction
	txn.Commit()

	return db, nil
}
