// Copyright 2019 Brightbox Systems Ltd
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

package mocks

import (
	brightbox "github.com/brightbox/gobrightbox"

	mock "github.com/stretchr/testify/mock"
)

// CloudAccess is an autogenerated mock type for the CloudAccess type
type CloudAccess struct {
	mock.Mock
}

// AddServersToServerGroup provides a mock function with given fields: identifier, serverIds
func (_m *CloudAccess) AddServersToServerGroup(identifier string, serverIds []string) (*brightbox.ServerGroup, error) {
	ret := _m.Called(identifier, serverIds)

	var r0 *brightbox.ServerGroup
	if rf, ok := ret.Get(0).(func(string, []string) *brightbox.ServerGroup); ok {
		r0 = rf(identifier, serverIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(identifier, serverIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) CloudIP(identifier string) (*brightbox.CloudIP, error) {
	ret := _m.Called(identifier)

	var r0 *brightbox.CloudIP
	if rf, ok := ret.Get(0).(func(string) *brightbox.CloudIP); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIPs provides a mock function with given fields:
func (_m *CloudAccess) CloudIPs() ([]brightbox.CloudIP, error) {
	ret := _m.Called()

	var r0 []brightbox.CloudIP
	if rf, ok := ret.Get(0).(func() []brightbox.CloudIP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.CloudIP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCloudIP provides a mock function with given fields: newCloudIP
func (_m *CloudAccess) CreateCloudIP(newCloudIP *brightbox.CloudIPOptions) (*brightbox.CloudIP, error) {
	ret := _m.Called(newCloudIP)

	var r0 *brightbox.CloudIP
	if rf, ok := ret.Get(0).(func(*brightbox.CloudIPOptions) *brightbox.CloudIP); ok {
		r0 = rf(newCloudIP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.CloudIPOptions) error); ok {
		r1 = rf(newCloudIP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallPolicy provides a mock function with given fields: policyOptions
func (_m *CloudAccess) CreateFirewallPolicy(policyOptions *brightbox.FirewallPolicyOptions) (*brightbox.FirewallPolicy, error) {
	ret := _m.Called(policyOptions)

	var r0 *brightbox.FirewallPolicy
	if rf, ok := ret.Get(0).(func(*brightbox.FirewallPolicyOptions) *brightbox.FirewallPolicy); ok {
		r0 = rf(policyOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.FirewallPolicyOptions) error); ok {
		r1 = rf(policyOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallRule provides a mock function with given fields: ruleOptions
func (_m *CloudAccess) CreateFirewallRule(ruleOptions *brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error) {
	ret := _m.Called(ruleOptions)

	var r0 *brightbox.FirewallRule
	if rf, ok := ret.Get(0).(func(*brightbox.FirewallRuleOptions) *brightbox.FirewallRule); ok {
		r0 = rf(ruleOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.FirewallRuleOptions) error); ok {
		r1 = rf(ruleOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoadBalancer provides a mock function with given fields: newDetails
func (_m *CloudAccess) CreateLoadBalancer(newDetails *brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(newDetails)

	var r0 *brightbox.LoadBalancer
	if rf, ok := ret.Get(0).(func(*brightbox.LoadBalancerOptions) *brightbox.LoadBalancer); ok {
		r0 = rf(newDetails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.LoadBalancerOptions) error); ok {
		r1 = rf(newDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServer provides a mock function with given fields: newServer
func (_m *CloudAccess) CreateServer(newServer *brightbox.ServerOptions) (*brightbox.Server, error) {
	ret := _m.Called(newServer)

	var r0 *brightbox.Server
	if rf, ok := ret.Get(0).(func(*brightbox.ServerOptions) *brightbox.Server); ok {
		r0 = rf(newServer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.ServerOptions) error); ok {
		r1 = rf(newServer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServerGroup provides a mock function with given fields: newServerGroup
func (_m *CloudAccess) CreateServerGroup(newServerGroup *brightbox.ServerGroupOptions) (*brightbox.ServerGroup, error) {
	ret := _m.Called(newServerGroup)

	var r0 *brightbox.ServerGroup
	if rf, ok := ret.Get(0).(func(*brightbox.ServerGroupOptions) *brightbox.ServerGroup); ok {
		r0 = rf(newServerGroup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.ServerGroupOptions) error); ok {
		r1 = rf(newServerGroup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyCloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyCloudIP(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyFirewallPolicy provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyFirewallPolicy(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyLoadBalancer provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyLoadBalancer(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyServer provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyServer(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyServerGroup provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyServerGroup(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirewallPolicies provides a mock function with given fields:
func (_m *CloudAccess) FirewallPolicies() ([]brightbox.FirewallPolicy, error) {
	ret := _m.Called()

	var r0 []brightbox.FirewallPolicy
	if rf, ok := ret.Get(0).(func() []brightbox.FirewallPolicy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.FirewallPolicy)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancer provides a mock function with given fields: identifier
func (_m *CloudAccess) LoadBalancer(identifier string) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(identifier)

	var r0 *brightbox.LoadBalancer
	if rf, ok := ret.Get(0).(func(string) *brightbox.LoadBalancer); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancers provides a mock function with given fields:
func (_m *CloudAccess) LoadBalancers() ([]brightbox.LoadBalancer, error) {
	ret := _m.Called()

	var r0 []brightbox.LoadBalancer
	if rf, ok := ret.Get(0).(func() []brightbox.LoadBalancer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MapCloudIP provides a mock function with given fields: identifier, destination
func (_m *CloudAccess) MapCloudIP(identifier string, destination string) error {
	ret := _m.Called(identifier, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(identifier, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveServersFromServerGroup provides a mock function with given fields: identifier, serverIds
func (_m *CloudAccess) RemoveServersFromServerGroup(identifier string, serverIds []string) (*brightbox.ServerGroup, error) {
	ret := _m.Called(identifier, serverIds)

	var r0 *brightbox.ServerGroup
	if rf, ok := ret.Get(0).(func(string, []string) *brightbox.ServerGroup); ok {
		r0 = rf(identifier, serverIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(identifier, serverIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Server provides a mock function with given fields: identifier
func (_m *CloudAccess) Server(identifier string) (*brightbox.Server, error) {
	ret := _m.Called(identifier)

	var r0 *brightbox.Server
	if rf, ok := ret.Get(0).(func(string) *brightbox.Server); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroup provides a mock function with given fields: identifier
func (_m *CloudAccess) ServerGroup(identifier string) (*brightbox.ServerGroup, error) {
	ret := _m.Called(identifier)

	var r0 *brightbox.ServerGroup
	if rf, ok := ret.Get(0).(func(string) *brightbox.ServerGroup); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroups provides a mock function with given fields:
func (_m *CloudAccess) ServerGroups() ([]brightbox.ServerGroup, error) {
	ret := _m.Called()

	var r0 []brightbox.ServerGroup
	if rf, ok := ret.Get(0).(func() []brightbox.ServerGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.ServerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnMapCloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) UnMapCloudIP(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFirewallRule provides a mock function with given fields: ruleOptions
func (_m *CloudAccess) UpdateFirewallRule(ruleOptions *brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error) {
	ret := _m.Called(ruleOptions)

	var r0 *brightbox.FirewallRule
	if rf, ok := ret.Get(0).(func(*brightbox.FirewallRuleOptions) *brightbox.FirewallRule); ok {
		r0 = rf(ruleOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallRule)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.FirewallRuleOptions) error); ok {
		r1 = rf(ruleOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLoadBalancer provides a mock function with given fields: newDetails
func (_m *CloudAccess) UpdateLoadBalancer(newDetails *brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(newDetails)

	var r0 *brightbox.LoadBalancer
	if rf, ok := ret.Get(0).(func(*brightbox.LoadBalancerOptions) *brightbox.LoadBalancer); ok {
		r0 = rf(newDetails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*brightbox.LoadBalancerOptions) error); ok {
		r1 = rf(newDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
