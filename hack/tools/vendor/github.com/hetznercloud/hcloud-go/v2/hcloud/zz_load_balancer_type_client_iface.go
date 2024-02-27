// Code generated by ifacemaker; DO NOT EDIT.

package hcloud

import (
	"context"
)

// ILoadBalancerTypeClient ...
type ILoadBalancerTypeClient interface {
	// GetByID retrieves a Load Balancer type by its ID. If the Load Balancer type does not exist, nil is returned.
	GetByID(ctx context.Context, id int64) (*LoadBalancerType, *Response, error)
	// GetByName retrieves a Load Balancer type by its name. If the Load Balancer type does not exist, nil is returned.
	GetByName(ctx context.Context, name string) (*LoadBalancerType, *Response, error)
	// Get retrieves a Load Balancer type by its ID if the input can be parsed as an integer, otherwise it
	// retrieves a Load Balancer type by its name. If the Load Balancer type does not exist, nil is returned.
	Get(ctx context.Context, idOrName string) (*LoadBalancerType, *Response, error)
	// List returns a list of Load Balancer types for a specific page.
	//
	// Please note that filters specified in opts are not taken into account
	// when their value corresponds to their zero value or when they are empty.
	List(ctx context.Context, opts LoadBalancerTypeListOpts) ([]*LoadBalancerType, *Response, error)
	// All returns all Load Balancer types.
	All(ctx context.Context) ([]*LoadBalancerType, error)
	// AllWithOpts returns all Load Balancer types for the given options.
	AllWithOpts(ctx context.Context, opts LoadBalancerTypeListOpts) ([]*LoadBalancerType, error)
}
