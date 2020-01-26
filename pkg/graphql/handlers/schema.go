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
			"meetupGroup": {
				Name: "meetupGroup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//Meetup Schema
			"meetup": {
				Name: "meetup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
			//Company Schema
			"company": {
				Name: "company",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			//Sponsor Schema
			"sponsor": {
				Name: "sponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			//SponsorTier Schema
			"sponsorTier": {
				Name: "sponsorTier",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			//Presentation Schema
			"presentation": {
				Name: "presentation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			//Speaker Schema
			"speaker": {
				Name: "speaker",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			//SpeakerToCompany Schema
			"speakerToCompany": {
				Name: "speakerToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"speakerID": {
						Name:    "speakerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
					"companyID": {
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
				},
			},
			//SponsorTierToMeetupGroup Schema
			"sponsorTierToMeetupGroup": {
				Name: "sponsorTierToMeetupGroup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": {
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"sponsorTierID": {
						Name:    "sponsorTierID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorTierID"},
					},
				},
			},
			//MeetupGroupToOrganizer Schema
			"meetupGroupToOrganizer": {
				Name: "meetupGroupToOrganizer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": {
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"organizerID": {
						Name:    "organizerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "OrganizerID"},
					},
				},
			},
			//MeetupGroupToMeetup Schema
			"meetupGroupToMeetup": {
				Name: "meetupGroupToMeetup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": {
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"meetupID": {
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//MeetupToSponsor Schema
			"meetupToSponsor": {
				Name: "meetupToSponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": {
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
					"sponsorID": {
						Name:    "sponsorID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorID"},
					},
				},
			},
			//SponsorToCompany Schema
			"sponsorToCompany": {
				Name: "sponsorToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"companyID": {
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
					"sponsorID": {
						Name:    "sponsorID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorID"},
					},
				},
			},
			//MeetupToPresentation Schema
			"meetupToPresentation": {
				Name: "meetupToPresentation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": {
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
					"presentationID": {
						Name:    "presentationID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "PresentationID"},
					},
				},
			},
			//PresentationToSpeaker Schema
			"presentationToSpeaker": {
				Name: "presentationToSpeaker",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"presentationID": {
						Name:    "presentationID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "PresentationID"},
					},
					"speakerID": {
						Name:    "speakerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
				},
			},
			//MeetupGroupToEcosystemMember Schema
			"meetupGroupToEcosystemMember": {
				Name: "meetupGroupToEcosystemMember",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": {
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"companyID": {
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
				},
			},
			//SponsorTierToCompany Schema
			"sponsorTierToCompany": {
				Name: "sponsorTierToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"sponsorTierID": {
						Name:    "sponsorTierID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorTierID"},
					},
					"companyID": {
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
				},
			},
		},
	}
	return schema
}
