// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "kpt.dev/configsync/clientgen/apis"
	configmanagementv1 "kpt.dev/configsync/clientgen/apis/typed/configmanagement/v1"
	fakeconfigmanagementv1 "kpt.dev/configsync/clientgen/apis/typed/configmanagement/v1/fake"
	configsyncv1alpha1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1alpha1"
	fakeconfigsyncv1alpha1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1alpha1/fake"
	configsyncv1beta1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1beta1"
	fakeconfigsyncv1beta1 "kpt.dev/configsync/clientgen/apis/typed/configsync/v1beta1/fake"
	hubv1 "kpt.dev/configsync/clientgen/apis/typed/hub/v1"
	fakehubv1 "kpt.dev/configsync/clientgen/apis/typed/hub/v1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// ConfigmanagementV1 retrieves the ConfigmanagementV1Client
func (c *Clientset) ConfigmanagementV1() configmanagementv1.ConfigmanagementV1Interface {
	return &fakeconfigmanagementv1.FakeConfigmanagementV1{Fake: &c.Fake}
}

// ConfigsyncV1beta1 retrieves the ConfigsyncV1beta1Client
func (c *Clientset) ConfigsyncV1beta1() configsyncv1beta1.ConfigsyncV1beta1Interface {
	return &fakeconfigsyncv1beta1.FakeConfigsyncV1beta1{Fake: &c.Fake}
}

// ConfigsyncV1alpha1 retrieves the ConfigsyncV1alpha1Client
func (c *Clientset) ConfigsyncV1alpha1() configsyncv1alpha1.ConfigsyncV1alpha1Interface {
	return &fakeconfigsyncv1alpha1.FakeConfigsyncV1alpha1{Fake: &c.Fake}
}

// HubV1 retrieves the HubV1Client
func (c *Clientset) HubV1() hubv1.HubV1Interface {
	return &fakehubv1.FakeHubV1{Fake: &c.Fake}
}
