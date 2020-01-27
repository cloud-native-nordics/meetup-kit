/*
	Note: This file is autogenerated! Do not edit it manually!
	Edit client_speaker_template.go instead, and run
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

// SpeakerClient is an interface for accessing Speaker-specific API objects
type SpeakerClient interface {
	// New returns a new Speaker
	New() *api.Speaker
	// Get returns the Speaker matching given UID from the storage
	Get(runtime.UID) (*api.Speaker, error)
	// Set saves the given Speaker into persistent storage
	Set(*api.Speaker) error
	// Patch performs a strategic merge patch on the object with
	// the given UID, using the byte-encoded patch given
	Patch(runtime.UID, []byte) error
	// Find returns the Speaker matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	Find(filter filterer.BaseFilter) (*api.Speaker, error)
	// FindAll returns multiple Speakers matching the given filter, filters can
	// match e.g. the Object's Name, UID or a specific property
	FindAll(filter filterer.BaseFilter) ([]*api.Speaker, error)
	// Delete deletes the Speaker with the given UID from the storage
	Delete(uid runtime.UID) error
	// List returns a list of all Speakers available
	List() ([]*api.Speaker, error)
}

// Speakers returns the SpeakerClient for the IgniteInternalClient instance
func (c *MeetOpsInternalClient) Speakers() SpeakerClient {
	if c.speakerClient == nil {
		c.speakerClient = newSpeakerClient(c.storage, c.gv)
	}

	return c.speakerClient
}

// speakerClient is a struct implementing the SpeakerClient interface
// It uses a shared storage instance passed from the Client together with its own Filterer
type speakerClient struct {
	storage  storage.Storage
	filterer *filterer.Filterer
	gvk      schema.GroupVersionKind
}

// newSpeakerClient builds the speakerClient struct using the storage implementation and a new Filterer
func newSpeakerClient(s storage.Storage, gv schema.GroupVersion) SpeakerClient {
	return &speakerClient{
		storage:  s,
		filterer: filterer.NewFilterer(s),
		gvk:      gv.WithKind(api.KindSpeaker.Title()),
	}
}

// New returns a new Object of its kind
func (c *speakerClient) New() *api.Speaker {
	log.Tracef("Client.New; GVK: %v", c.gvk)
	obj, err := c.storage.New(c.gvk)
	if err != nil {
		panic(fmt.Sprintf("Client.New must not return an error: %v", err))
	}
	return obj.(*api.Speaker)
}

// Find returns a single Speaker based on the given Filter
func (c *speakerClient) Find(filter filterer.BaseFilter) (*api.Speaker, error) {
	log.Tracef("Client.Find; GVK: %v", c.gvk)
	object, err := c.filterer.Find(c.gvk, filter)
	if err != nil {
		return nil, err
	}

	return object.(*api.Speaker), nil
}

// FindAll returns multiple Speakers based on the given Filter
func (c *speakerClient) FindAll(filter filterer.BaseFilter) ([]*api.Speaker, error) {
	log.Tracef("Client.FindAll; GVK: %v", c.gvk)
	matches, err := c.filterer.FindAll(c.gvk, filter)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Speaker, 0, len(matches))
	for _, item := range matches {
		results = append(results, item.(*api.Speaker))
	}

	return results, nil
}

// Get returns the Speaker matching given UID from the storage
func (c *speakerClient) Get(uid runtime.UID) (*api.Speaker, error) {
	log.Tracef("Client.Get; UID: %q, GVK: %v", uid, c.gvk)
	object, err := c.storage.Get(c.gvk, uid)
	if err != nil {
		return nil, err
	}

	return object.(*api.Speaker), nil
}

// Set saves the given Speaker into the persistent storage
func (c *speakerClient) Set(speaker *api.Speaker) error {
	log.Tracef("Client.Set; UID: %q, GVK: %v", speaker.GetUID(), c.gvk)
	return c.storage.Set(c.gvk, speaker)
}

// Patch performs a strategic merge patch on the object with
// the given UID, using the byte-encoded patch given
func (c *speakerClient) Patch(uid runtime.UID, patch []byte) error {
	return c.storage.Patch(c.gvk, uid, patch)
}

// Delete deletes the Speaker from the storage
func (c *speakerClient) Delete(uid runtime.UID) error {
	log.Tracef("Client.Delete; UID: %q, GVK: %v", uid, c.gvk)
	return c.storage.Delete(c.gvk, uid)
}

// List returns a list of all Speakers available
func (c *speakerClient) List() ([]*api.Speaker, error) {
	log.Tracef("Client.List; GVK: %v", c.gvk)
	list, err := c.storage.List(c.gvk)
	if err != nil {
		return nil, err
	}

	results := make([]*api.Speaker, 0, len(list))
	for _, item := range list {
		results = append(results, item.(*api.Speaker))
	}

	return results, nil
}