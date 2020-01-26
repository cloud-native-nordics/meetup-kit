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
				},
			},
			//SponsorTier Schema
			"sponsorTier": &memdb.TableSchema{
				Name: "sponsorTier",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
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
				},
			},
			//SpeakerToCompany Schema
			"speakerToCompany": &memdb.TableSchema{
				Name: "speakerToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"speakerID": &memdb.IndexSchema{
						Name:    "speakerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
					"companyID": &memdb.IndexSchema{
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
				},
			},
			//SponsorTierToMeetupGroup Schema
			"sponsorTierToMeetupGroup": &memdb.TableSchema{
				Name: "sponsorTierToMeetupGroup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": &memdb.IndexSchema{
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"sponsorTierID": &memdb.IndexSchema{
						Name:    "sponsorTierID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorTierID"},
					},
				},
			},
			//MeetupGroupToOrganizer Schema
			"meetupGroupToOrganizer": &memdb.TableSchema{
				Name: "meetupGroupToOrganizer",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": &memdb.IndexSchema{
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"organizerID": &memdb.IndexSchema{
						Name:    "organizerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "OrganizerID"},
					},
				},
			},
			//MeetupGroupToMeetup Schema
			"meetupGroupToMeetup": &memdb.TableSchema{
				Name: "meetupGroupToMeetup",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": &memdb.IndexSchema{
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"meetupID": &memdb.IndexSchema{
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
				},
			},
			//MeetupToSponsor Schema
			"meetupToSponsor": &memdb.TableSchema{
				Name: "meetupToSponsor",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": &memdb.IndexSchema{
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
					"sponsorID": &memdb.IndexSchema{
						Name:    "sponsorID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorID"},
					},
				},
			},
			//SponsorToCompany Schema
			"sponsorToCompany": &memdb.TableSchema{
				Name: "sponsorToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"companyID": &memdb.IndexSchema{
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
					"sponsorID": &memdb.IndexSchema{
						Name:    "sponsorID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorID"},
					},
				},
			},
			//MeetupToPresentation Schema
			"meetupToPresentation": &memdb.TableSchema{
				Name: "meetupToPresentation",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupID": &memdb.IndexSchema{
						Name:    "meetupID",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "MeetupID"},
					},
					"presentationID": &memdb.IndexSchema{
						Name:    "presentationID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "PresentationID"},
					},
				},
			},
			//PresentationToSpeaker Schema
			"presentationToSpeaker": &memdb.TableSchema{
				Name: "presentationToSpeaker",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"presentationID": &memdb.IndexSchema{
						Name:    "presentationID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "PresentationID"},
					},
					"speakerID": &memdb.IndexSchema{
						Name:    "speakerID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SpeakerID"},
					},
				},
			},
			//MeetupGroupToEcosystemMember Schema
			"meetupGroupToEcosystemMember": &memdb.TableSchema{
				Name: "meetupGroupToEcosystemMember",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"meetupGroupID": &memdb.IndexSchema{
						Name:    "meetupGroupID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "MeetupGroupID"},
					},
					"companyID": &memdb.IndexSchema{
						Name:    "companyID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CompanyID"},
					},
				},
			},
			//SponsorTierToCompany Schema
			"sponsorTierToCompany": &memdb.TableSchema{
				Name: "sponsorTierToCompany",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"sponsorTierID": &memdb.IndexSchema{
						Name:    "sponsorTierID",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "SponsorTierID"},
					},
					"companyID": &memdb.IndexSchema{
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
