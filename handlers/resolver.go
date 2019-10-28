package handlers

import (
	"context"

	"github.com/cloud-native-nordics/stats-api/generated"
	"github.com/cloud-native-nordics/stats-api/models"
	"github.com/cloud-native-nordics/stats-api/repositories"
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
func (r *Resolver) SponsorTier() generated.SponsorTierResolver {
	return &sponsorTierResolver{r}
}

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) Sponsors(ctx context.Context, obj *models.Meetup) ([]*models.Sponsor, error) {
	panic("not implemented")
}
func (r *meetupResolver) Presentations(ctx context.Context, obj *models.Meetup) ([]*models.Presentation, error) {
	panic("not implemented")
}

type meetupGroupResolver struct{ *Resolver }

func (r *meetupGroupResolver) SponsorTiers(ctx context.Context, obj *models.MeetupGroup) ([]*models.SponsorTier, error) {
	panic("not implemented")
}
func (r *meetupGroupResolver) Organizers(ctx context.Context, obj *models.MeetupGroup) ([]*models.Speaker, error) {
	panic("not implemented")
}
func (r *meetupGroupResolver) EcosystemMembers(ctx context.Context, obj *models.MeetupGroup) ([]*models.Company, error) {
	panic("not implemented")
}
func (r *meetupGroupResolver) Meetups(ctx context.Context, obj *models.MeetupGroup) ([]*models.Meetup, error) {
	panic("not implemented")
}

type presentationResolver struct{ *Resolver }

func (r *presentationResolver) Speakers(ctx context.Context, obj *models.Presentation) ([]*models.Speaker, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) MeetupGroups(ctx context.Context) ([]*models.MeetupGroup, error) {
	panic("not implemented")
}
func (r *queryResolver) MeetupGroup(ctx context.Context, meetupID string) (*models.MeetupGroup, error) {
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
func (r *queryResolver) Presentations(ctx context.Context) ([]*models.Presentation, error) {
	panic("not implemented")
}
func (r *queryResolver) Presentation(ctx context.Context, id string) (*models.Presentation, error) {
	panic("not implemented")
}
func (r *queryResolver) Speakers(ctx context.Context) ([]*models.Speaker, error) {
	panic("not implemented")
}
func (r *queryResolver) Speaker(ctx context.Context, id string) (*models.Speaker, error) {
	panic("not implemented")
}

func (r *queryResolver) SlackInvite(ctx context.Context, email string) (string, error) {
	res := r.slackRepository.DoInvite(email)
	return res, nil
}

type speakerResolver struct{ *Resolver }

func (r *speakerResolver) Company(ctx context.Context, obj *models.Speaker) (*models.Company, error) {
	panic("not implemented")
}

type sponsorResolver struct{ *Resolver }

func (r *sponsorResolver) Company(ctx context.Context, obj *models.Sponsor) (*models.Company, error) {
	panic("not implemented")
}

type sponsorTierResolver struct{ *Resolver }

func (r *sponsorTierResolver) Company(ctx context.Context, obj *models.SponsorTier) (*models.Company, error) {
	panic("not implemented")
}
