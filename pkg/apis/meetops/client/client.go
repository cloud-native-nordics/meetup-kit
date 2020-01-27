// TODO: Docs
package client

import (
	api "github.com/cloud-native-nordics/meetup-kit/pkg/apis/meetops"
	"github.com/weaveworks/gitops-toolkit/pkg/client"
	"github.com/weaveworks/gitops-toolkit/pkg/runtime"
	"github.com/weaveworks/gitops-toolkit/pkg/storage"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: Autogenerate this!

// NewClient creates a client for the specified storage
func NewClient(s storage.Storage) *Client {
	return &Client{
		MeetOpsInternalClient: NewMeetOpsInternalClient(s),
	}
}

// Client is a struct providing high-level access to objects in a storage
// The resource-specific client interfaces are automatically generated based
// off client_resource_template.go. The auto-generation can be done with hack/client.sh
// At the moment MeetOpsInternalClient is the default client. If more than this client
// is created in the future, the MeetOpsInternalClient will be accessible under
// Client.MeetOpsInternal() instead.
type Client struct {
	*MeetOpsInternalClient
}

func NewMeetOpsInternalClient(s storage.Storage) *MeetOpsInternalClient {
	return &MeetOpsInternalClient{
		storage:        s,
		dynamicClients: map[schema.GroupVersionKind]client.DynamicClient{},
		gv:             api.SchemeGroupVersion,
	}
}

type MeetOpsInternalClient struct {
	storage           storage.Storage
	gv                schema.GroupVersion
	speakerClient     SpeakerClient
	companyClient     CompanyClient
	meetupClient      MeetupClient
	meetupgroupClient MeetupGroupClient

	dynamicClients map[schema.GroupVersionKind]client.DynamicClient
}

// Dynamic returns the DynamicClient for the Client instance, for the specific kind
func (c *MeetOpsInternalClient) Dynamic(kind runtime.Kind) (dc client.DynamicClient) {
	var ok bool
	gvk := c.gv.WithKind(kind.Title())
	if dc, ok = c.dynamicClients[gvk]; !ok {
		dc = client.NewDynamicClient(c.storage, gvk)
		c.dynamicClients[gvk] = dc
	}

	return
}
