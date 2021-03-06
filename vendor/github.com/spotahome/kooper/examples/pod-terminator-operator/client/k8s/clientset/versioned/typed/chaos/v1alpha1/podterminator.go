package v1alpha1

import (
	v1alpha1 "github.com/spotahome/kooper/examples/pod-terminator-operator/apis/chaos/v1alpha1"
	scheme "github.com/spotahome/kooper/examples/pod-terminator-operator/client/k8s/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PodTerminatorsGetter has a method to return a PodTerminatorInterface.
// A group's client should implement this interface.
type PodTerminatorsGetter interface {
	PodTerminators() PodTerminatorInterface
}

// PodTerminatorInterface has methods to work with PodTerminator resources.
type PodTerminatorInterface interface {
	Create(*v1alpha1.PodTerminator) (*v1alpha1.PodTerminator, error)
	Update(*v1alpha1.PodTerminator) (*v1alpha1.PodTerminator, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PodTerminator, error)
	List(opts v1.ListOptions) (*v1alpha1.PodTerminatorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodTerminator, err error)
	PodTerminatorExpansion
}

// podTerminators implements PodTerminatorInterface
type podTerminators struct {
	client rest.Interface
}

// newPodTerminators returns a PodTerminators
func newPodTerminators(c *ChaosV1alpha1Client) *podTerminators {
	return &podTerminators{
		client: c.RESTClient(),
	}
}

// Get takes name of the podTerminator, and returns the corresponding podTerminator object, and an error if there is any.
func (c *podTerminators) Get(name string, options v1.GetOptions) (result *v1alpha1.PodTerminator, err error) {
	result = &v1alpha1.PodTerminator{}
	err = c.client.Get().
		Resource("podterminators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodTerminators that match those selectors.
func (c *podTerminators) List(opts v1.ListOptions) (result *v1alpha1.PodTerminatorList, err error) {
	result = &v1alpha1.PodTerminatorList{}
	err = c.client.Get().
		Resource("podterminators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podTerminators.
func (c *podTerminators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("podterminators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a podTerminator and creates it.  Returns the server's representation of the podTerminator, and an error, if there is any.
func (c *podTerminators) Create(podTerminator *v1alpha1.PodTerminator) (result *v1alpha1.PodTerminator, err error) {
	result = &v1alpha1.PodTerminator{}
	err = c.client.Post().
		Resource("podterminators").
		Body(podTerminator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a podTerminator and updates it. Returns the server's representation of the podTerminator, and an error, if there is any.
func (c *podTerminators) Update(podTerminator *v1alpha1.PodTerminator) (result *v1alpha1.PodTerminator, err error) {
	result = &v1alpha1.PodTerminator{}
	err = c.client.Put().
		Resource("podterminators").
		Name(podTerminator.Name).
		Body(podTerminator).
		Do().
		Into(result)
	return
}

// Delete takes name of the podTerminator and deletes it. Returns an error if one occurs.
func (c *podTerminators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("podterminators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podTerminators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("podterminators").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched podTerminator.
func (c *podTerminators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodTerminator, err error) {
	result = &v1alpha1.PodTerminator{}
	err = c.client.Patch(pt).
		Resource("podterminators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
