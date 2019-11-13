package handlers

import (
	"context"

	"github.com/cloud-native-nordics/stats-api/generated"
	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/cloud-native-nordics/stats-api/repositories"
	"github.com/golang/glog"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	statsRepository *repositories.StatsRepository
	slackRepository *repositories.SlackRepository
}

func NewResolver(statsRepository *repositories.StatsRepository, slackRepository *repositories.SlackRepository) *Resolver {
	return &Resolver{statsRepository: statsRepository, slackRepository: slackRepository}
}

func (r *Resolver) Meetup() generated.MeetupResolver {
	return &meetupResolver{r}
}
func (r *Resolver) MeetupGroup() generated.MeetupGroupResolver {
	return &meetupGroupResolver{r}
}
func (r *Resolver) Presentation() generated.PresentationResolver {
	return &presentationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Speaker() generated.SpeakerResolver {
	return &speakerResolver{r}
}
func (r *Resolver) Sponsor() generated.SponsorResolver {
	return &sponsorResolver{r}
}
func (r *Resolver) Company() generated.CompanyResolver {
	return &companyResolver{r}
}
func (r *Resolver) SponsorTier() generated.SponsorTierResolver {
	return &sponsorTierResolver{r}
}

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) Sponsors(ctx context.Context, obj *models.Meetup) ([]*models.Sponsor, error) {
	sponsors, err := r.statsRepository.GetSponsorsForMeetup(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return sponsors, nil
}
func (r *meetupResolver) Presentations(ctx context.Context, obj *models.Meetup) ([]*models.Presentation, error) {
	presentations, err := r.statsRepository.GetPresentationsForMeetup(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return presentations, nil
}
func (r *meetupResolver) MeetupGroup(ctx context.Context, obj *models.Meetup) (*models.MeetupGroup, error) {
	meetupGroup, err := r.statsRepository.GetMeetupGroupForMeetup(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetupGroup, nil
}

type meetupGroupResolver struct{ *Resolver }

func (r *meetupGroupResolver) SponsorTiers(ctx context.Context, obj *models.MeetupGroup) ([]*models.SponsorTier, error) {
	sponsorTiers, err := r.statsRepository.GetSponsorTiersForMeetupGroup(obj.MeetupID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return sponsorTiers, nil
}
func (r *meetupGroupResolver) Organizers(ctx context.Context, obj *models.MeetupGroup) ([]*models.Speaker, error) {
	organizers, err := r.statsRepository.GetOrganizersForMeetupGroup(obj.MeetupID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return organizers, nil
}
func (r *meetupGroupResolver) EcosystemMembers(ctx context.Context, obj *models.MeetupGroup) ([]*models.Company, error) {
	companies, err := r.statsRepository.GetEcosystemMembersForMeetupGroup(obj.MeetupID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return companies, nil
}

func (r *meetupGroupResolver) Meetups(ctx context.Context, obj *models.MeetupGroup) ([]*models.Meetup, error) {
	meetups, err := r.statsRepository.GetMeetupsForMeetupGroup(obj.MeetupID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetups, nil
}

func (r *meetupGroupResolver) MemberCount(ctx context.Context, obj *models.MeetupGroup) (int, error) {
	count, err := repositories.GetMeetupInfoFromAPI(*obj)

	if err != nil {
		glog.V(1).Info(err)
		return 0, err
	}

	return count.Members, nil
}

type presentationResolver struct{ *Resolver }

func (r *presentationResolver) Speakers(ctx context.Context, obj *models.Presentation) ([]*models.Speaker, error) {
	speakers, err := r.statsRepository.GetSpeakersForPresentation(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return speakers, nil
}

func (r *presentationResolver) Meetup(ctx context.Context, obj *models.Presentation) (*models.Meetup, error) {
	meetup, err := r.statsRepository.GetMeetupForPresentation(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetup, nil
}

type companyResolver struct{ *Resolver }

func (r *companyResolver) Countries(ctx context.Context, obj *models.Company) ([]*string, error) {
	countries, err := r.statsRepository.GetCountriesForCompany(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

func (r *companyResolver) SponsorTiers(ctx context.Context, obj *models.Company) ([]*models.SponsorTier, error) {
	sponsorTiers, err := r.statsRepository.GetSponsorTiersForCompany(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return sponsorTiers, nil
}

func (r *companyResolver) Speakers(ctx context.Context, obj *models.Company) ([]*models.Speaker, error) {
	speakers, err := r.statsRepository.GetSpeakersForCompany(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return speakers, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) MeetupGroups(ctx context.Context) ([]*models.MeetupGroup, error) {
	meetupGroups, err := r.statsRepository.GetAllMeetupGroups()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetupGroups, nil
}
func (r *queryResolver) MeetupGroup(ctx context.Context, meetupID string) (*models.MeetupGroup, error) {
	meetupGroups, err := r.statsRepository.GetMeetupGroup(meetupID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetupGroups, nil
}
func (r *queryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	companies, err := r.statsRepository.GetAllCompanies()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return companies, nil
}
func (r *queryResolver) Company(ctx context.Context, id string) (*models.Company, error) {
	company, err := r.statsRepository.GetCompany(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return company, nil
}
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	meetups, err := r.statsRepository.GetAllMeetups()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetups, nil
}
func (r *queryResolver) Meetup(ctx context.Context, id int) (*models.Meetup, error) {
	meetup, err := r.statsRepository.GetMeetup(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetup, nil
}
func (r *queryResolver) Presentations(ctx context.Context) ([]*models.Presentation, error) {
	presentations, err := r.statsRepository.GetAllPresentations()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return presentations, nil
}
func (r *queryResolver) Presentation(ctx context.Context, id string) (*models.Presentation, error) {
	presentation, err := r.statsRepository.GetPresentation(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return presentation, nil
}
func (r *queryResolver) Speakers(ctx context.Context) ([]*models.Speaker, error) {
	speakers, err := r.statsRepository.GetAllSpeakers()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return speakers, nil
}
func (r *queryResolver) Speaker(ctx context.Context, id string) (*models.Speaker, error) {
	speaker, err := r.statsRepository.GetSpeaker(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return speaker, nil
}

func (r *queryResolver) SlackInvite(ctx context.Context, email string) (string, error) {
	res := r.slackRepository.DoInvite(email)
	return res, nil
}

type speakerResolver struct{ *Resolver }

func (r *speakerResolver) Company(ctx context.Context, obj *models.Speaker) (*models.Company, error) {
	company, err := r.statsRepository.GetCompanyForSpeaker(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return company, nil
}

func (r *speakerResolver) Countries(ctx context.Context, obj *models.Speaker) ([]*string, error) {
	countries, err := r.statsRepository.GetCountriesForSpeaker(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

func (r *speakerResolver) Presentations(ctx context.Context, obj *models.Speaker) ([]*models.Presentation, error) {
	presentations, err := r.statsRepository.GetPresentationsForSpeaker(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return presentations, nil
}

type sponsorResolver struct{ *Resolver }

func (r *sponsorResolver) Company(ctx context.Context, obj *models.Sponsor) (*models.Company, error) {
	company, err := r.statsRepository.GetCompanyForSponsor(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return company, nil
}

type sponsorTierResolver struct{ *Resolver }

func (r *sponsorTierResolver) Company(ctx context.Context, obj *models.SponsorTier) (*models.Company, error) {
	company, err := r.statsRepository.GetCompanyForSponsorTier(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return company, nil
}

func (r *sponsorTierResolver) MeetupGroups(ctx context.Context, obj *models.SponsorTier) ([]*models.MeetupGroup, error) {
	meetupGroups, err := r.statsRepository.GetMeetupGroupsForSponsorTier(obj.ID)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return meetupGroups, nil
}
