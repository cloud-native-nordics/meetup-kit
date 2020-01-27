/*
	Note: This file is autogenerated! Do not edit it manually!
	Edit client_meetup_template.go instead, and run
	hack/generate-client.sh afterwards.
*/

package client

import (
	"fmt"

	api "github.com/cloud-native-nordics/meetup-kit/pkg/apis/meetops"

	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/gitops-toolkit/pkg/runtime"
	"github.com/weaveworks/gitops-toolkit/pkg/storage"
	"github.com/weaveworks/gitops-toolkit/pkg/storage/filterer"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// MeetupClient is an interface for accessing Meetup-specific API objects
type MeetupClient interface {
	// New returns a new Meetup
	New() *api.Meetup
	// Get returns the Meetup matching given UID from the storage
	Get(runtime.UID) (*api.Meetup, error)
	// Set saves the given Meetup into persistent storage
	Set(*api.Meetup) error
	// Patch performs a strategic merge patch on the object with
	// the given UID, using the byte-encoded patch given
	Patch(runtime.UID, []byte) error
	// Find returns the Meetup matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	Find(filter filterer.BaseFilter) (*api.Meetup, error)
	// FindAll returns multiple Meetups matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	FindAll(filter filterer.BaseFilter) ([]*api.Meetup, error)
	// Delete deletes the Meetup with the given UID from the storage
	Delete(uid runtime.UID) error
	// List returns a list of all Meetups available
	List() ([]*api.Meetup, error)
}

// Meetups returns the MeetupClient for the IgniteInternalClient instance
func (c *MeetOpsInternalClient) Meetups() MeetupClient {
	if c.meetupClient == nil {
		c.meetupClient = newMeetupClient(c.storage, c.gv)
	}

	return c.meetupClient
}

// meetupClient is a struct implementing the MeetupClient interface
// It uses a shared storage instance passed from the Client together with its own Filterer
type meetupClient struct {
	storage  storage.Storage
	filterer *filterer.Filterer
	gvk      schema.GroupVersionKind
}

// newMeetupClient builds the meetupClient struct using the storage implementation and a new Filterer
func newMeetupClient(s storage.Storage, gv schema.GroupVersion) MeetupClient {
	return &meetupClient{
		storage:  s,
		filterer: filterer.NewFilterer(s),
		gvk:      gv.WithKind(api.KindMeetup.Title()),
	}
}

// New returns a new Object of its kind
func (c *meetupClient) New() *api.Meetup {
	log.Tracef("Client.New; GVK: %v", c.gvk)
	obj, err := c.storage.New(c.gvk)
	if err != nil {
		panic(fmt.Sprintf("Client.New must not return an error: %v", err))
	}
	return obj.(*api.Meetup)
}

// Find returns a single Meetup based on the given Filter
func (c *meetupClient) Find(filter filterer.BaseFilter) (*api.Meetup, error) {
	log.Tracef("Client.Find; GVK: %v", c.gvk)
	object, err := c.filterer.Find(c.gvk, filter)
	if err != nil {
		return nil, err
	}

	return object.(*api.Meetup), nil
}

// FindAll returns multiple Meetups based on the given Filter
func (c *meetupClient) FindAll(filter filterer.BaseFilter) ([]*api.Meetup, error) {
	log.Tracef("Client.FindAll; GVK: %v", c.gvk)
	matches, err := c.filterer.FindAll(c.gvk, filter)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Meetup, 0, len(matches))
	for _, item := range matches {
		results = append(results, item.(*api.Meetup))
	}

	return results, nil
}

// Get returns the Meetup matching given UID from the storage
func (c *meetupClient) Get(uid runtime.UID) (*api.Meetup, error) {
	log.Tracef("Client.Get; UID: %q, GVK: %v", uid, c.gvk)
	object, err := c.storage.Get(c.gvk, uid)
	if err != nil {
		return nil, err
	}

	return object.(*api.Meetup), nil
}

// Set saves the given Meetup into the persistent storage
func (c *meetupClient) Set(meetup *api.Meetup) error {
	log.Tracef("Client.Set; UID: %q, GVK: %v", meetup.GetUID(), c.gvk)
	return c.storage.Set(c.gvk, meetup)
}

// Patch performs a strategic merge patch on the object with
// the given UID, using the byte-encoded patch given
func (c *meetupClient) Patch(uid runtime.UID, patch []byte) error {
	return c.storage.Patch(c.gvk, uid, patch)
}

// Delete deletes the Meetup from the storage
func (c *meetupClient) Delete(uid runtime.UID) error {
	log.Tracef("Client.Delete; UID: %q, GVK: %v", uid, c.gvk)
	return c.storage.Delete(c.gvk, uid)
}

// List returns a list of all Meetups available
func (c *meetupClient) List() ([]*api.Meetup, error) {
	log.Tracef("Client.List; GVK: %v", c.gvk)
	list, err := c.storage.List(c.gvk)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Meetup, 0, len(list))
	for _, item := range list {
		results = append(results, item.(*api.Meetup))
	}

	return results, nil
}