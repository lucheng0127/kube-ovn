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

package v1

import (
	"context"
	"time"

	v1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	scheme "github.com/kubeovn/kube-ovn/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VpcNatGatewayIpipsGetter has a method to return a VpcNatGatewayIpipInterface.
// A group's client should implement this interface.
type VpcNatGatewayIpipsGetter interface {
	VpcNatGatewayIpips() VpcNatGatewayIpipInterface
}

// VpcNatGatewayIpipInterface has methods to work with VpcNatGatewayIpip resources.
type VpcNatGatewayIpipInterface interface {
	Create(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.CreateOptions) (*v1.VpcNatGatewayIpip, error)
	Update(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.UpdateOptions) (*v1.VpcNatGatewayIpip, error)
	UpdateStatus(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.UpdateOptions) (*v1.VpcNatGatewayIpip, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VpcNatGatewayIpip, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VpcNatGatewayIpipList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VpcNatGatewayIpip, err error)
	VpcNatGatewayIpipExpansion
}

// vpcNatGatewayIpips implements VpcNatGatewayIpipInterface
type vpcNatGatewayIpips struct {
	client rest.Interface
}

// newVpcNatGatewayIpips returns a VpcNatGatewayIpips
func newVpcNatGatewayIpips(c *KubeovnV1Client) *vpcNatGatewayIpips {
	return &vpcNatGatewayIpips{
		client: c.RESTClient(),
	}
}

// Get takes name of the vpcNatGatewayIpip, and returns the corresponding vpcNatGatewayIpip object, and an error if there is any.
func (c *vpcNatGatewayIpips) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VpcNatGatewayIpip, err error) {
	result = &v1.VpcNatGatewayIpip{}
	err = c.client.Get().
		Resource("vpc-nat-gateway-ipip").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VpcNatGatewayIpips that match those selectors.
func (c *vpcNatGatewayIpips) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VpcNatGatewayIpipList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VpcNatGatewayIpipList{}
	err = c.client.Get().
		Resource("vpc-nat-gateway-ipip").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vpcNatGatewayIpips.
func (c *vpcNatGatewayIpips) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("vpc-nat-gateway-ipip").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vpcNatGatewayIpip and creates it.  Returns the server's representation of the vpcNatGatewayIpip, and an error, if there is any.
func (c *vpcNatGatewayIpips) Create(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.CreateOptions) (result *v1.VpcNatGatewayIpip, err error) {
	result = &v1.VpcNatGatewayIpip{}
	err = c.client.Post().
		Resource("vpc-nat-gateway-ipip").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcNatGatewayIpip).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vpcNatGatewayIpip and updates it. Returns the server's representation of the vpcNatGatewayIpip, and an error, if there is any.
func (c *vpcNatGatewayIpips) Update(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.UpdateOptions) (result *v1.VpcNatGatewayIpip, err error) {
	result = &v1.VpcNatGatewayIpip{}
	err = c.client.Put().
		Resource("vpc-nat-gateway-ipip").
		Name(vpcNatGatewayIpip.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcNatGatewayIpip).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vpcNatGatewayIpips) UpdateStatus(ctx context.Context, vpcNatGatewayIpip *v1.VpcNatGatewayIpip, opts metav1.UpdateOptions) (result *v1.VpcNatGatewayIpip, err error) {
	result = &v1.VpcNatGatewayIpip{}
	err = c.client.Put().
		Resource("vpc-nat-gateway-ipip").
		Name(vpcNatGatewayIpip.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vpcNatGatewayIpip).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vpcNatGatewayIpip and deletes it. Returns an error if one occurs.
func (c *vpcNatGatewayIpips) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vpc-nat-gateway-ipip").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vpcNatGatewayIpips) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("vpc-nat-gateway-ipip").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vpcNatGatewayIpip.
func (c *vpcNatGatewayIpips) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VpcNatGatewayIpip, err error) {
	result = &v1.VpcNatGatewayIpip{}
	err = c.client.Patch(pt).
		Resource("vpc-nat-gateway-ipip").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
