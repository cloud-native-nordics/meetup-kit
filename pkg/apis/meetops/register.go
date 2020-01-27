package meetops

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeBuilder the schema builder
	SchemeBuilder = runtime.NewSchemeBuilder(
		addKnownTypes,
	)

	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

const (
	// GroupName is the group name use in this package
	GroupName = "meetops.cncf.io"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: runtime.APIVersionInternal,
}

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Company{},
		&Speaker{},
		&MeetupGroup{},
		&Meetup{},
	)

	return nil
}
