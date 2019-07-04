package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"os"
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
	meetupGroups    []models.MeetupGroup
	organizers      []models.Organizer
	companies       []models.Company
	meetups         []models.Meetup
	sponsors        []models.Sponsor
	speakers        []models.Speaker
	presentations   []models.Presentation
	countries       []models.Country
	entityToCountry []models.EntityToCountry
}

type jsonStructure struct {
	MeetupGroups []models.MeetupGroup `json:"meetupGroups"`
}

//NewStatsManager fetches the stats.json file, marshals to structs, creates in-mem db
//and then returns a reference to the in-mem db.
func NewStatsManager(URL string) *memdb.MemDB {
	sm := &StatsManager{
		URL:      URL,
		filepath: "./data/stats.json",
	}
	schema := sm.createDatabaseSchema()

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
	var meetupGroups jsonStructure
	output := &unmarshalledData{}

	file, err := ioutil.ReadFile(sm.filepath)

	if err != nil {
		return nil, err
	}

	glog.V(5).Info("Unmarhsalling stats.json")
	err = json.Unmarshal(file, &meetupGroups)

	if err != nil {
		return nil, err
	}
	var denmark = "denmark"
	var finland = "finland"
	var sweden = "sweden"

	countries := []models.Country{
		models.Country{&denmark, &denmark},
		models.Country{&finland, &finland},
		models.Country{&sweden, &sweden},
	}

	output.countries = countries
	output.meetupGroups = meetupGroups.MeetupGroups

	for _, meetupGroup := range meetupGroups.MeetupGroups {
		//Add organizers
		for _, organizer := range meetupGroup.Organizers {
			organizerToBe := organizer
			organizerToBe.MeetupGroupID = &meetupGroup.MeetupID
			output.organizers = append(output.organizers, *organizerToBe)

			for _, country := range organizer.Countries {
				output.entityToCountry = append(output.entityToCountry,
					models.EntityToCountry{
						ID:         &organizer.ID,
						CountryID:  country,
						EntityType: "organizer",
					})
			}

			if organizer.Company != nil {
				output.companies = append(output.companies, *organizer.Company)

				for _, country := range organizer.Company.Countries {
					output.entityToCountry = append(output.entityToCountry,
						models.EntityToCountry{
							ID:         &organizer.Company.ID,
							CountryID:  country,
							EntityType: "company",
						})
				}
			}
		}

		//Add meetups
		for _, meetup := range meetupGroup.Meetups {
			meetupToBe := meetup
			meetupToBe.MeetupGroupID = &meetupGroup.MeetupID
			output.meetups = append(output.meetups, *meetupToBe)

			//Add sponsors
			if meetup.Sponsors != nil {
				sponsorsToBe := meetup.Sponsors
				sponsorsToBe.ID = strconv.Itoa(rand.Intn(100000000000))
				sponsorsToBe.MeetupID = meetup.ID

				if sponsorsToBe.Venue != nil {
					sponsorsToBe.Venue.SponsorID = sponsorsToBe.ID

					for _, country := range sponsorsToBe.Venue.Countries {
						output.entityToCountry = append(output.entityToCountry,
							models.EntityToCountry{
								ID:         &sponsorsToBe.Venue.ID,
								CountryID:  country,
								EntityType: "sponsor",
							})
					}
				}

				//Add other sponsor
				for _, other := range sponsorsToBe.Other {
					other.SponsorID = other.ID
					for _, country := range other.Countries {
						output.entityToCountry = append(output.entityToCountry,
							models.EntityToCountry{
								ID:         &other.ID,
								CountryID:  country,
								EntityType: "sponsor",
							})
					}
				}

				output.sponsors = append(output.sponsors, *sponsorsToBe)
			}

			//Add presentations
			for _, presentation := range meetup.Presentations {
				presentationToBe := presentation
				presentationToBe.MeetupID = meetup.ID
				presentationToBe.ID = strconv.Itoa(rand.Intn(100000000000))

				output.presentations = append(output.presentations, *presentationToBe)

				for _, speaker := range presentation.Speakers {
					speakerToBe := speaker
					speakerToBe.PresentationID = &presentation.ID

					output.speakers = append(output.speakers, *speakerToBe)

					if speaker.Company != nil {
						output.companies = append(output.companies, *speaker.Company)

						for _, country := range speaker.Company.Countries {
							output.entityToCountry = append(output.entityToCountry,
								models.EntityToCountry{
									ID:         &speaker.Company.ID,
									CountryID:  country,
									EntityType: "company",
								})
						}
					}
				}
			}
		}
	}

	return output, nil
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

	// Insert Sponsors
	glog.V(5).Infof("Inserting %d Sponsors", len(data.sponsors))
	for _, sponsor := range data.sponsors {
		if err := txn.Insert("sponsor", sponsor); err != nil {
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

//Create database schema
func (sm *StatsManager) createDatabaseSchema() *memdb.DBSchema {
	glog.V(5).Info("Creating database schema")
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			//MeetupGroup Schema
			"meetupGroup": &memdb.TableSchema{
				Name: "meetupGroup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"city": &memdb.IndexSchema{
						Name:         "city",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "City"},
					},
					"country": &memdb.IndexSchema{
						Name:         "country",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Country"},
					},
				},
			},
			//Organizer Schema
			"organizer": &memdb.TableSchema{
				Name: "organizer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"title": &memdb.IndexSchema{
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"email": &memdb.IndexSchema{
						Name:         "email",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Email"},
					},
					"github": &memdb.IndexSchema{
						Name:         "github",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Github"},
					},
					"twitter": &memdb.IndexSchema{
						Name:         "twitter",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Twitter"},
					},
					"speakersBureau": &memdb.IndexSchema{
						Name:         "speakersBureau",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakersBureau"},
					},
					"meetupGroupId": &memdb.IndexSchema{
						Name:         "meetupGroupId",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
				},
			},
			//Company Schema
			"company": &memdb.TableSchema{
				Name: "company",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"websiteURL": &memdb.IndexSchema{
						Name:         "websiteURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "WebsiteURL"},
					},
					"logoURL": &memdb.IndexSchema{
						Name:         "logoURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "LogoURL"},
					},
					"organizerID": &memdb.IndexSchema{
						Name:         "organizerID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "OrganizerID"},
					},
					"speakerID": &memdb.IndexSchema{
						Name:         "speakerID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
				},
			},
			//Meetup Schema
			"meetup": &memdb.TableSchema{
				Name: "meetup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"date": &memdb.IndexSchema{
						Name:         "date",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Date"},
					},
					"duration": &memdb.IndexSchema{
						Name:         "duration",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Duration"},
					},
					"attendees": &memdb.IndexSchema{
						Name:         "attendees",
						Unique:       false,
						AllowMissing: false,
						Indexer:      &memdb.IntFieldIndex{Field: "Attendees"},
					},
					"address": &memdb.IndexSchema{
						Name:         "address",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Address"},
					},
					"meetupGroupID": &memdb.IndexSchema{
						Name:         "meetupGroupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
				},
			},
			//Sponsor Schema
			"sponsor": &memdb.TableSchema{
				Name: "sponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": &memdb.IndexSchema{
						Name:         "meetupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//Presentation Schema
			"presentation": &memdb.TableSchema{
				Name: "presentation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"duration": &memdb.IndexSchema{
						Name:         "duration",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Duration"},
					},
					"title": &memdb.IndexSchema{
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"slides": &memdb.IndexSchema{
						Name:         "slides",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Slides"},
					},
					"meetupID": &memdb.IndexSchema{
						Name:         "meetupID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//Speaker Schema
			"speaker": &memdb.TableSchema{
				Name: "speaker",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
					"title": &memdb.IndexSchema{
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"email": &memdb.IndexSchema{
						Name:         "email",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Email"},
					},
					"github": &memdb.IndexSchema{
						Name:         "github",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Github"},
					},
					"speakersBureau": &memdb.IndexSchema{
						Name:         "speakersBureau",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SpeakersBureau"},
					},
					"presentationID": &memdb.IndexSchema{
						Name:         "presentationID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "PresentationID"},
					},
				},
			},
			//Country Schema
			"country": &memdb.TableSchema{
				Name: "country",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:         "name",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			//CountryToEntity Schema
			"entityToCountry": &memdb.TableSchema{
				Name: "entityToCountry",
				Indexes: map[string]*memdb.IndexSchema{
					"countryId": &memdb.IndexSchema{
						Name:    "countryId",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "CountryID"},
					},
					"id": &memdb.IndexSchema{
						Name:   "id",
						Unique: true,

						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"entityType": &memdb.IndexSchema{
						Name:    "entityType",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "EntityType"},
					},
				},
			},
		},
	}
	return schema
}
