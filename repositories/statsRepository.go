package repositories

import (
	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/hashicorp/go-memdb"
)

//StatsRepository is used to query the in-mem database
type StatsRepository struct {
	db *memdb.MemDB
}

//NewStatsRepository returns a new Stats Repository to be used for fetching data from the in-mem database
func NewStatsRepository(db *memdb.MemDB) *StatsRepository {
	return &StatsRepository{
		db: db,
	}
}

func (sr *StatsRepository) GetAllMeetupGroups() ([]*models.MeetupGroup, error) {
	output := []*models.MeetupGroup{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all meetup groups
	it, err := txn.Get("meetupGroup", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.MeetupGroup)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetAllOrganizers() ([]*models.Organizer, error) {
	output := []*models.Organizer{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all organizers
	it, err := txn.Get("organizer", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Organizer)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetOrganizer(id string) (*models.Organizer, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get organizer by id
	it, err := txn.First("organizer", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Organizer)
	return &out, nil
}

func (sr *StatsRepository) GetAllCompanies() ([]*models.Company, error) {
	output := []*models.Company{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all companies
	it, err := txn.Get("company", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Company)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetCompany(id string) (*models.Company, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get company by id
	it, err := txn.First("company", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Company)
	return &out, nil
}

func (sr *StatsRepository) GetAllMeetups() ([]*models.Meetup, error) {
	output := []*models.Meetup{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all meetups
	it, err := txn.Get("meetup", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Meetup)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetup(id int) (*models.Meetup, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get meetup by id
	it, err := txn.First("meetup", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Meetup)
	return &out, nil
}

func (sr *StatsRepository) GetAllSponsors() ([]*models.Sponsor, error) {
	output := []*models.Sponsor{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all sponsors
	it, err := txn.Get("sponsor", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Sponsor)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetSponsor(id string) (*models.Sponsor, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get sponsor by id
	it, err := txn.First("sponsor", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Sponsor)
	return &out, nil
}

func (sr *StatsRepository) GetAllMeetupSponsors() ([]*models.MeetupSponsor, error) {
	output := []*models.MeetupSponsor{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all sponsors
	it, err := txn.Get("meetupSponsor", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.MeetupSponsor)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetupSponsor(id string) (*models.MeetupSponsor, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get sponsor by id
	it, err := txn.First("meetupSponsor", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.MeetupSponsor)
	return &out, nil
}

func (sr *StatsRepository) GetMember(id string) (*models.Member, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get sponsor by id
	it, err := txn.First("member", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Member)
	return &out, nil
}

func (sr *StatsRepository) GetAllMembers() ([]*models.Member, error) {
	output := []*models.Member{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all members
	it, err := txn.Get("member", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Member)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetAllPresentations() ([]*models.Presentation, error) {
	output := []*models.Presentation{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all presentations
	it, err := txn.Get("presentation", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Presentation)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetPresentation(id string) (*models.Presentation, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get presentation by id
	it, err := txn.First("presentation", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Presentation)
	return &out, nil
}

func (sr *StatsRepository) GetAllSpeakers() ([]*models.Speaker, error) {
	output := []*models.Speaker{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all speakers
	it, err := txn.Get("speaker", "id")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.Speaker)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetSpeaker(id string) (*models.Speaker, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get speaker by id
	it, err := txn.First("speaker", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.Speaker)
	return &out, nil
}

func (sr *StatsRepository) GetMeetupGroup(id string) (*models.MeetupGroup, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get meetup group by id
	it, err := txn.First("meetupGroup", "id", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.MeetupGroup)
	return &out, nil
}

func (sr *StatsRepository) GetCountriesForEntity(entityID string, entityType string) ([]*models.Country, error) {
	output := []*models.Country{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all country relations
	it, err := txn.Get("entityToCountry", "id", entityID)
	if err != nil {
		return nil, err
	}

	entityToCountry := []*models.EntityToCountry{}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.EntityToCountry)
		entityToCountry = append(entityToCountry, &p)
	}

	for _, entityToCountry := range entityToCountry {
		it, err := txn.Get("country", "id", *entityToCountry.CountryID)
		if err != nil {
			return nil, err
		}
		for obj := it.Next(); obj != nil; obj = it.Next() {
			p := obj.(models.Country)
			output = append(output, &p)
		}
	}

	return output, nil
}
