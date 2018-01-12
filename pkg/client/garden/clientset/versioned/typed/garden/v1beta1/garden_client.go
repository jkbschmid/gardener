package v1beta1

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	"github.com/gardener/gardener/pkg/client/garden/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type GardenV1beta1Interface interface {
	RESTClient() rest.Interface
	CloudProfilesGetter
	CrossSecretBindingsGetter
	PrivateSecretBindingsGetter
	QuotasGetter
	SeedsGetter
	ShootsGetter
}

// GardenV1beta1Client is used to interact with features provided by the garden.sapcloud.io group.
type GardenV1beta1Client struct {
	restClient rest.Interface
}

func (c *GardenV1beta1Client) CloudProfiles() CloudProfileInterface {
	return newCloudProfiles(c)
}

func (c *GardenV1beta1Client) CrossSecretBindings(namespace string) CrossSecretBindingInterface {
	return newCrossSecretBindings(c, namespace)
}

func (c *GardenV1beta1Client) PrivateSecretBindings(namespace string) PrivateSecretBindingInterface {
	return newPrivateSecretBindings(c, namespace)
}

func (c *GardenV1beta1Client) Quotas(namespace string) QuotaInterface {
	return newQuotas(c, namespace)
}

func (c *GardenV1beta1Client) Seeds() SeedInterface {
	return newSeeds(c)
}

func (c *GardenV1beta1Client) Shoots(namespace string) ShootInterface {
	return newShoots(c, namespace)
}

// NewForConfig creates a new GardenV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*GardenV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GardenV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new GardenV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GardenV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GardenV1beta1Client for the given RESTClient.
func New(c rest.Interface) *GardenV1beta1Client {
	return &GardenV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *GardenV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}