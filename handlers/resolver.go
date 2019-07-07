package handlers

import (
	"context"

	"github.com/cloud-native-nordics/stats-api/generated"
	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/cloud-native-nordics/stats-api/repositories"
	"github.com/golang/glog"
)

type Resolver struct {
	statsRepository *repositories.StatsRepository
}

func NewResolver(statsRepository *repositories.StatsRepository) *Resolver {
	return &Resolver{statsRepository: statsRepository}
}

func (r *Resolver) Company() generated.CompanyResolver {
	return &companyResolver{r}
}

func (r *Resolver) Organizer() generated.OrganizerResolver {
	return &organizerResolver{r}
}
func (r *Resolver) Other() generated.OtherResolver {
	return &otherResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Speaker() generated.SpeakerResolver {
	return &speakerResolver{r}
}

func (r *Resolver) Venue() generated.VenueResolver {
	return &venueResolver{r}
}

type companyResolver struct{ *Resolver }

func (r *companyResolver) Countries(ctx context.Context, obj *models.Company) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "company")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

type organizerResolver struct{ *Resolver }

func (r *organizerResolver) Countries(ctx context.Context, obj *models.Organizer) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "organizer")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

type otherResolver struct{ *Resolver }

func (r *otherResolver) Countries(ctx context.Context, obj *models.Other) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "other")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
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
func (r *queryResolver) Organizers(ctx context.Context) ([]*models.Organizer, error) {
	organizers, err := r.statsRepository.GetAllOrganizers()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return organizers, nil
}
func (r *queryResolver) Organizer(ctx context.Context, id string) (*models.Organizer, error) {
	organizer, err := r.statsRepository.GetOrganizer(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return organizer, nil
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
func (r *queryResolver) Sponsors(ctx context.Context) ([]*models.Sponsor, error) {
	sponsors, err := r.statsRepository.GetAllSponsors()

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return sponsors, nil
}
func (r *queryResolver) Sponsor(ctx context.Context, id string) (*models.Sponsor, error) {
	sponsor, err := r.statsRepository.GetSponsor(id)

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return sponsor, nil
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

type speakerResolver struct{ *Resolver }

func (r *speakerResolver) Countries(ctx context.Context, obj *models.Speaker) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "speaker")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

type venueResolver struct{ *Resolver }

func (r *venueResolver) Countries(ctx context.Context, obj *models.Venue) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "venue")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}
