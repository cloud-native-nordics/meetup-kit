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
	it, err := txn.Get("meetupGroup", "meetupID")
	if err != nil {
		return nil, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(models.MeetupGroup)
		output = append(output, &p)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetupGroup(id string) (*models.MeetupGroup, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get meetup group by id
	it, err := txn.First("meetupGroup", "meetupID", id)
	if err != nil {
		return nil, err
	}

	out := it.(models.MeetupGroup)
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
