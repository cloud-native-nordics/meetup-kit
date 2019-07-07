package handlers

import (
	"github.com/golang/glog"
	"github.com/hashicorp/go-memdb"
)

//Create database schema
func createDatabaseSchema() *memdb.DBSchema {
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
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
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
			//Venue Schema
			"venue": &memdb.TableSchema{
				Name: "venue",
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
					"sponsorID": &memdb.IndexSchema{
						Name:         "sponsorID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SponsorID"},
					},
				},
			},
			//Other Schema
			"other": &memdb.TableSchema{
				Name: "other",
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
					"sponsorID": &memdb.IndexSchema{
						Name:         "sponsorID",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "SponsorID"},
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
