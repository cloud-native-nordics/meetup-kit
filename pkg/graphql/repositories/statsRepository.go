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

// ### Meetup Groups ###
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

func (sr *StatsRepository) GetMeetupGroupForMeetup(id int) (*models.MeetupGroup, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	//Get meetup group by id
	relations, err := txn.First("meetupGroupToMeetup", "meetupID", id)
	if err != nil {
		return nil, err
	}

	relation := relations.(models.MeetupGroupToMeetup)
	it, err := txn.First("meetupGroup", "id", relation.MeetupGroupID)
	if err != nil {
		return nil, err
	}

	out := it.(models.MeetupGroup)
	return &out, nil
}

func (sr *StatsRepository) GetSponsorTiersForMeetupGroup(id string) ([]*models.SponsorTier, error) {
	output := []*models.SponsorTier{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all sponsor tier relations for meetup group
	relations, err := txn.Get("sponsorTierToMeetupGroup", "meetupGroupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.SponsorTierToMeetupGroup)
		it, err := txn.First("sponsorTier", "id", relation.SponsorTierID)
		if err != nil {
			return nil, err
		}
		result := it.(models.SponsorTier)
		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetupGroupsForSponsorTier(id string) ([]*models.MeetupGroup, error) {
	output := []*models.MeetupGroup{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	// List all sponsor tier relations for meetup group
	relations, err := txn.Get("sponsorTierToMeetupGroup", "sponsorTierID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.SponsorTierToMeetupGroup)
		it, err := txn.Get("meetupGroup", "id", relation.MeetupGroupID)
		if err != nil {
			return nil, err
		}
		for obj := it.Next(); obj != nil; obj = it.Next() {
			meetupGroup := obj.(models.MeetupGroup)
			output = append(output, &meetupGroup)
		}
	}
	return output, nil
}

func (sr *StatsRepository) GetOrganizersForMeetupGroup(id string) ([]*models.Speaker, error) {
	output := []*models.Speaker{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("meetupGroupToOrganizer", "meetupGroupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.MeetupGroupToOrganizer)
		it, err := txn.First("speaker", "id", relation.OrganizerID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Speaker)
		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetEcosystemMembersForMeetupGroup(id string) ([]*models.Company, error) {
	output := []*models.Company{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("meetupGroupToEcosystemMember", "meetupGroupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.MeetupGroupToEcosystemMember)
		it, err := txn.First("company", "id", relation.CompanyID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Company)
		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetupsForMeetupGroup(id string) ([]*models.Meetup, error) {
	output := []*models.Meetup{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("meetupGroupToMeetup", "meetupGroupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.MeetupGroupToMeetup)
		it, err := txn.First("meetup", "id", relation.MeetupID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Meetup)
		output = append(output, &result)
	}

	return output, nil
}

// ### Sponsor Tier ###
func (sr *StatsRepository) GetCompanyForSponsorTier(id string) (*models.Company, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.First("sponsorTierToCompany", "sponsorTierID", id)
	if err != nil {
		return nil, err
	}

	relation := relations.(models.SponsorTierToCompany)
	it, err := txn.First("company", "id", relation.CompanyID)
	if err != nil {
		return nil, err
	}
	result := it.(models.Company)

	return &result, nil
}

// ### Companies ###
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

func (sr *StatsRepository) GetCountriesForCompany(id string) ([]*string, error) {
	var output []*string
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	sponsorTierRelations, err := txn.Get("sponsorTierToCompany", "companyID", id)
	if err != nil {
		return nil, err
	}

	for sponTierRelObj := sponsorTierRelations.Next(); sponTierRelObj != nil; sponTierRelObj = sponsorTierRelations.Next() {
		sponTierRel := sponTierRelObj.(models.SponsorTierToCompany)
		meetupGrpRelations, err := txn.Get("sponsorTierToMeetupGroup", "sponsorTierID", sponTierRel.SponsorTierID)

		if err != nil {
			return nil, err
		}

		for meetGrpRelObj := meetupGrpRelations.Next(); meetGrpRelObj != nil; meetGrpRelObj = meetupGrpRelations.Next() {
			meetGrpRel := meetGrpRelObj.(models.SponsorTierToMeetupGroup)

			it, err := txn.First("meetupGroup", "id", meetGrpRel.MeetupGroupID)
			if err != nil {
				return nil, err
			}

			meetGrp := it.(models.MeetupGroup)
			if !contains(output, &meetGrp.Country) {
				output = append(output, &meetGrp.Country)
			}
		}
	}

	return output, nil
}

func (sr *StatsRepository) GetSponsorTiersForCompany(id string) ([]*models.SponsorTier, error) {
	output := []*models.SponsorTier{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	sponsorTierRelations, err := txn.Get("sponsorTierToCompany", "companyID", id)
	if err != nil {
		return nil, err
	}

	for sponTierRelObj := sponsorTierRelations.Next(); sponTierRelObj != nil; sponTierRelObj = sponsorTierRelations.Next() {
		sponTierRel := sponTierRelObj.(models.SponsorTierToCompany)

		it, err := txn.First("sponsorTier", "id", sponTierRel.SponsorTierID)
		if err != nil {
			return nil, err
		}
		result := it.(models.SponsorTier)

		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetSpeakersForCompany(id string) ([]*models.Speaker, error) {
	output := []*models.Speaker{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	speakerRelations, err := txn.Get("speakerToCompany", "companyID", id)
	if err != nil {
		return nil, err
	}

	for speakTierRelObj := speakerRelations.Next(); speakTierRelObj != nil; speakTierRelObj = speakerRelations.Next() {
		speakRel := speakTierRelObj.(models.SpeakerToCompany)

		it, err := txn.First("speaker", "id", speakRel.SpeakerID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Speaker)

		output = append(output, &result)
	}

	return output, nil
}

// ### Meetups ###
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

func (sr *StatsRepository) GetSponsorsForMeetup(id int) ([]*models.Sponsor, error) {
	output := []*models.Sponsor{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("meetupToSponsor", "meetupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.MeetupToSponsor)
		it, err := txn.First("sponsor", "id", relation.SponsorID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Sponsor)
		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetPresentationsForMeetup(id int) ([]*models.Presentation, error) {
	output := []*models.Presentation{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("meetupToPresentation", "meetupID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.MeetupToPresentation)
		it, err := txn.First("presentation", "id", relation.PresentationID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Presentation)
		output = append(output, &result)
	}

	return output, nil
}

/// ### Presentations ###
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

func (sr *StatsRepository) GetSpeakersForPresentation(id string) ([]*models.Speaker, error) {
	output := []*models.Speaker{}
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.Get("presentationToSpeaker", "presentationID", id)
	if err != nil {
		return nil, err
	}

	for obj := relations.Next(); obj != nil; obj = relations.Next() {
		relation := obj.(models.PresentationToSpeaker)
		it, err := txn.First("speaker", "id", relation.SpeakerID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Speaker)
		output = append(output, &result)
	}

	return output, nil
}

func (sr *StatsRepository) GetMeetupForPresentation(id string) (*models.Meetup, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.First("meetupToPresentation", "presentationID", id)
	if err != nil {
		return nil, err
	}

	relation := relations.(models.MeetupToPresentation)
	it, err := txn.First("meetup", "id", relation.MeetupID)
	if err != nil {
		return nil, err
	}
	result := it.(models.Meetup)

	return &result, nil
}

// ### Speakers ###
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

func (sr *StatsRepository) GetCompanyForSpeaker(id string) (*models.Company, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.First("speakerToCompany", "speakerID", id)
	if err != nil {
		return nil, err
	}

	relation, done := relations.(models.SpeakerToCompany)
	if done {
		it, err := txn.First("company", "id", relation.CompanyID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Company)

		return &result, nil
	}
	return nil, nil
}

func (sr *StatsRepository) GetCountriesForSpeaker(id string) ([]*string, error) {
	var output []*string
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	presentationRelations, err := txn.Get("presentationToSpeaker", "speakerID", id)
	if err != nil {
		return nil, err
	}

	for presRelObj := presentationRelations.Next(); presRelObj != nil; presRelObj = presentationRelations.Next() {
		presRel := presRelObj.(models.PresentationToSpeaker)
		meetupRelations, err := txn.Get("meetupToPresentation", "presentationID", presRel.PresentationID)

		if err != nil {
			return nil, err
		}

		for meetRelObj := meetupRelations.Next(); meetRelObj != nil; meetRelObj = meetupRelations.Next() {
			meetRel := meetRelObj.(models.MeetupToPresentation)

			meetGrpRelations, err := txn.Get("meetupGroupToMeetup", "meetupID", meetRel.MeetupID)
			if err != nil {
				return nil, err
			}

			for meetGrpRelObj := meetGrpRelations.Next(); meetGrpRelObj != nil; meetGrpRelObj = meetGrpRelations.Next() {
				meetGrpRel := meetGrpRelObj.(models.MeetupGroupToMeetup)

				it, err := txn.First("meetupGroup", "id", meetGrpRel.MeetupGroupID)
				if err != nil {
					return nil, err
				}

				meetGrp := it.(models.MeetupGroup)
				if !contains(output, &meetGrp.Country) {
					output = append(output, &meetGrp.Country)
				}
			}
		}
	}
	return output, nil
}

func (sr *StatsRepository) GetPresentationsForSpeaker(id string) ([]*models.Presentation, error) {
	var output []*models.Presentation
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	presentationRelations, err := txn.Get("presentationToSpeaker", "speakerID", id)
	if err != nil {
		return nil, err
	}

	for presRelObj := presentationRelations.Next(); presRelObj != nil; presRelObj = presentationRelations.Next() {
		presRel := presRelObj.(models.PresentationToSpeaker)

		it, err := txn.First("presentation", "id", presRel.PresentationID)
		if err != nil {
			return nil, err
		}

		result := it.(models.Presentation)

		output = append(output, &result)
	}
	return output, nil
}

// ### Sponsor ###
func (sr *StatsRepository) GetCompanyForSponsor(id string) (*models.Company, error) {
	// Create read-only transaction
	txn := sr.db.Txn(false)
	defer txn.Abort()

	relations, err := txn.First("sponsorToCompany", "sponsorID", id)
	if err != nil {
		return nil, err
	}

	relation, done := relations.(models.SponsorToCompany)
	if done {
		it, err := txn.First("company", "id", relation.CompanyID)
		if err != nil {
			return nil, err
		}
		result := it.(models.Company)

		return &result, nil
	}
	return nil, nil
}

// ### Helpers ###
func contains(slice []*string, item *string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[*s] = struct{}{}
	}

	_, ok := set[*item]
	return ok
}
