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
func (r *Resolver) Meetup() generated.MeetupResolver {
	return &meetupResolver{r}
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
func (r *Resolver) Sponsor() generated.SponsorResolver {
	return &sponsorResolver{r}
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

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) Sponsors(ctx context.Context, obj *models.Meetup) ([]*models.Sponsor, error) {
	panic("not implemented")
}

type otherResolver struct{ *Resolver }

func (r *otherResolver) Countries(ctx context.Context, obj *models.Other) ([]*models.Country, error) {
	panic("not implemented")
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
	panic("not implemented")
}
func (r *queryResolver) Organizers(ctx context.Context) ([]*models.Organizer, error) {
	panic("not implemented")
}
func (r *queryResolver) Organizer(ctx context.Context, id string) (*models.Organizer, error) {
	panic("not implemented")
}
func (r *queryResolver) Companies(ctx context.Context) ([]*models.Company, error) {
	panic("not implemented")
}
func (r *queryResolver) Company(ctx context.Context, id string) (*models.Company, error) {
	panic("not implemented")
}
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	panic("not implemented")
}
func (r *queryResolver) Meetup(ctx context.Context, id int) (*models.Meetup, error) {
	panic("not implemented")
}
func (r *queryResolver) Sponsors(ctx context.Context) ([]*models.Sponsor, error) {
	panic("not implemented")
}
func (r *queryResolver) Sponsor(ctx context.Context, id string) (*models.Sponsor, error) {
	panic("not implemented")
}
func (r *queryResolver) Presentations(ctx context.Context) ([]*models.Presentation, error) {
	panic("not implemented")
}
func (r *queryResolver) Presentation(ctx context.Context, title string) (*models.Presentation, error) {
	panic("not implemented")
}
func (r *queryResolver) Speakers(ctx context.Context) ([]*models.Speaker, error) {
	panic("not implemented")
}
func (r *queryResolver) Speaker(ctx context.Context, title string) (*models.Speaker, error) {
	panic("not implemented")
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

type sponsorResolver struct{ *Resolver }

func (r *sponsorResolver) Other(ctx context.Context, obj *models.Sponsor) (*models.Other, error) {
	panic("not implemented")
}

func (r *sponsorResolver) Countries(ctx context.Context, obj *models.Sponsor) ([]*models.Country, error) {
	countries, err := r.statsRepository.GetCountriesForEntity(obj.ID, "sponsor")

	if err != nil {
		glog.V(1).Info(err)
		return nil, err
	}

	return countries, nil
}

type venueResolver struct{ *Resolver }

func (r *venueResolver) Countries(ctx context.Context, obj *models.Venue) ([]*models.Country, error) {
	panic("not implemented")
}
