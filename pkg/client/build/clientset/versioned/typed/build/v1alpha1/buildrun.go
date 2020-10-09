/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	scheme "github.com/shipwright-io/build/pkg/client/build/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BuildRunsGetter has a method to return a BuildRunInterface.
// A group's client should implement this interface.
type BuildRunsGetter interface {
	BuildRuns(namespace string) BuildRunInterface
}

// BuildRunInterface has methods to work with BuildRun resources.
type BuildRunInterface interface {
	Create(*v1alpha1.BuildRun) (*v1alpha1.BuildRun, error)
	Update(*v1alpha1.BuildRun) (*v1alpha1.BuildRun, error)
	UpdateStatus(*v1alpha1.BuildRun) (*v1alpha1.BuildRun, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BuildRun, error)
	List(opts v1.ListOptions) (*v1alpha1.BuildRunList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildRun, err error)
	BuildRunExpansion
}

// buildRuns implements BuildRunInterface
type buildRuns struct {
	client rest.Interface
	ns     string
}

// newBuildRuns returns a BuildRuns
func newBuildRuns(c *BuildV1alpha1Client, namespace string) *buildRuns {
	return &buildRuns{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the buildRun, and returns the corresponding buildRun object, and an error if there is any.
func (c *buildRuns) Get(name string, options v1.GetOptions) (result *v1alpha1.BuildRun, err error) {
	result = &v1alpha1.BuildRun{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildruns").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BuildRuns that match those selectors.
func (c *buildRuns) List(opts v1.ListOptions) (result *v1alpha1.BuildRunList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BuildRunList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("buildruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested buildRuns.
func (c *buildRuns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("buildruns").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a buildRun and creates it.  Returns the server's representation of the buildRun, and an error, if there is any.
func (c *buildRuns) Create(buildRun *v1alpha1.BuildRun) (result *v1alpha1.BuildRun, err error) {
	result = &v1alpha1.BuildRun{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("buildruns").
		Body(buildRun).
		Do().
		Into(result)
	return
}

// Update takes the representation of a buildRun and updates it. Returns the server's representation of the buildRun, and an error, if there is any.
func (c *buildRuns) Update(buildRun *v1alpha1.BuildRun) (result *v1alpha1.BuildRun, err error) {
	result = &v1alpha1.BuildRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildruns").
		Name(buildRun.Name).
		Body(buildRun).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *buildRuns) UpdateStatus(buildRun *v1alpha1.BuildRun) (result *v1alpha1.BuildRun, err error) {
	result = &v1alpha1.BuildRun{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("buildruns").
		Name(buildRun.Name).
		SubResource("status").
		Body(buildRun).
		Do().
		Into(result)
	return
}

// Delete takes name of the buildRun and deletes it. Returns an error if one occurs.
func (c *buildRuns) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildruns").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *buildRuns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("buildruns").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched buildRun.
func (c *buildRuns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildRun, err error) {
	result = &v1alpha1.BuildRun{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("buildruns").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}