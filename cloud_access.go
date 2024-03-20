// Copyright 2020 Brightbox Systems Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8ssdk

import (
	"context"

	brightbox "github.com/brightbox/gobrightbox/v2"
)

// CloudAccess is an abstraction over the Brightbox API to allow testing
//
//go:generate mockery --name CloudAccess --boilerplate-file copyright_header
type CloudAccess interface {
	//Fetch a server
	Server(context.Context, string) (*brightbox.Server, error)

	//creates a new server
	CreateServer(context.Context, brightbox.ServerOptions) (*brightbox.Server, error)

	//Fetch a list of LoadBalancers
	LoadBalancers(context.Context) ([]brightbox.LoadBalancer, error)

	//Retrieves a detailed view of one load balancer
	LoadBalancer(context.Context, string) (*brightbox.LoadBalancer, error)

	//Creates a new load balancer
	CreateLoadBalancer(context.Context, brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error)

	//Updates an existing load balancer
	UpdateLoadBalancer(context.Context, brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error)

	//Retrieves a list of all cloud IPs
	CloudIPs(context.Context) ([]brightbox.CloudIP, error)

	//retrieves a detailed view of one cloud ip
	CloudIP(context.Context, string) (*brightbox.CloudIP, error)

	//Issues a request to map the cloud ip to the destination
	MapCloudIP(context.Context, string, brightbox.CloudIPAttachment) (*brightbox.CloudIP, error)

	//UnMapCloudIP issues a request to unmap the cloud ip
	UnMapCloudIP(context.Context, string) (*brightbox.CloudIP, error)

	//Creates a new Cloud IP
	CreateCloudIP(context.Context, brightbox.CloudIPOptions) (*brightbox.CloudIP, error)

	//adds servers to an existing server group
	AddServersToServerGroup(context.Context, string, brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error)

	//removes servers from an existing server group
	RemoveServersFromServerGroup(context.Context, string, brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error)

	// ServerGroups retrieves a list of all server groups
	ServerGroups(context.Context) ([]brightbox.ServerGroup, error)

	//Fetch a server group
	ServerGroup(context.Context, string) (*brightbox.ServerGroup, error)

	//creates a new server group
	CreateServerGroup(context.Context, brightbox.ServerGroupOptions) (*brightbox.ServerGroup, error)

	//creates a new firewall policy
	CreateFirewallPolicy(context.Context, brightbox.FirewallPolicyOptions) (*brightbox.FirewallPolicy, error)

	//creates a new firewall rule
	CreateFirewallRule(context.Context, brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error)

	//updates an existing firewall rule
	UpdateFirewallRule(context.Context, brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error)

	//retrieves a list of all firewall policies
	FirewallPolicies(context.Context) ([]brightbox.FirewallPolicy, error)

	// DestroyServer destroys an existing server
	DestroyServer(context.Context, string) (*brightbox.Server, error)

	// DestroyServerGroup destroys an existing server group
	DestroyServerGroup(context.Context, string) (*brightbox.ServerGroup, error)

	// DestroyFirewallPolicy issues a request to destroy the firewall policy
	DestroyFirewallPolicy(context.Context, string) (*brightbox.FirewallPolicy, error)

	// DestroyLoadBalancer issues a request to destroy the load balancer
	DestroyLoadBalancer(context.Context, string) (*brightbox.LoadBalancer, error)

	// DestroyCloudIP issues a request to destroy the cloud ip
	DestroyCloudIP(context.Context, string) (*brightbox.CloudIP, error)

	// ConfigMaps retrieves a list of all config maps
	Images(context.Context) ([]brightbox.Image, error)

	// ConfigMaps retrieves a list of all config maps
	ConfigMaps(context.Context) ([]brightbox.ConfigMap, error)

	// ConfigMap retrieves a detailed view on one config map
	ConfigMap(context.Context, string) (*brightbox.ConfigMap, error)

	// ServerTypes retrieves a list of all server types
	ServerTypes(context.Context) ([]brightbox.ServerType, error)

	// ServerType retrieves a detailed view on one server type
	ServerType(context.Context, string) (*brightbox.ServerType, error)
}
