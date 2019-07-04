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
