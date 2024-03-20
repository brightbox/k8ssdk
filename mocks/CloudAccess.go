// Copyright 2023 Brightbox Systems Ltd
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

// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	brightbox "github.com/brightbox/gobrightbox/v2"

	mock "github.com/stretchr/testify/mock"
)

// CloudAccess is an autogenerated mock type for the CloudAccess type
type CloudAccess struct {
	mock.Mock
}

// AddServersToServerGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudAccess) AddServersToServerGroup(_a0 context.Context, _a1 string, _a2 brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.ServerGroupMemberList) *brightbox.ServerGroup); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, brightbox.ServerGroupMemberList) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIP provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CloudIP(_a0 context.Context, _a1 string) (*brightbox.CloudIP, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.CloudIP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.CloudIP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIPs provides a mock function with given fields: _a0
func (_m *CloudAccess) CloudIPs(_a0 context.Context) ([]brightbox.CloudIP, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.CloudIP, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.CloudIP); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigMap provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) ConfigMap(_a0 context.Context, _a1 string) (*brightbox.ConfigMap, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.ConfigMap, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.ConfigMap); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigMaps provides a mock function with given fields: _a0
func (_m *CloudAccess) ConfigMaps(_a0 context.Context) ([]brightbox.ConfigMap, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.ConfigMap, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.ConfigMap); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCloudIP provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateCloudIP(_a0 context.Context, _a1 brightbox.CloudIPOptions) (*brightbox.CloudIP, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.CloudIPOptions) (*brightbox.CloudIP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.CloudIPOptions) *brightbox.CloudIP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.CloudIPOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallPolicy provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateFirewallPolicy(_a0 context.Context, _a1 brightbox.FirewallPolicyOptions) (*brightbox.FirewallPolicy, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.FirewallPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallPolicyOptions) (*brightbox.FirewallPolicy, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallPolicyOptions) *brightbox.FirewallPolicy); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.FirewallPolicyOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallRule provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateFirewallRule(_a0 context.Context, _a1 brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.FirewallRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallRuleOptions) *brightbox.FirewallRule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.FirewallRuleOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoadBalancer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateLoadBalancer(_a0 context.Context, _a1 brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.LoadBalancerOptions) *brightbox.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.LoadBalancerOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateServer(_a0 context.Context, _a1 brightbox.ServerOptions) (*brightbox.Server, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.ServerOptions) (*brightbox.Server, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.ServerOptions) *brightbox.Server); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.ServerOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServerGroup provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) CreateServerGroup(_a0 context.Context, _a1 brightbox.ServerGroupOptions) (*brightbox.ServerGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.ServerGroupOptions) (*brightbox.ServerGroup, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.ServerGroupOptions) *brightbox.ServerGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.ServerGroupOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyCloudIP provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) DestroyCloudIP(_a0 context.Context, _a1 string) (*brightbox.CloudIP, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.CloudIP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.CloudIP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyFirewallPolicy provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) DestroyFirewallPolicy(_a0 context.Context, _a1 string) (*brightbox.FirewallPolicy, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.FirewallPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.FirewallPolicy, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.FirewallPolicy); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyLoadBalancer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) DestroyLoadBalancer(_a0 context.Context, _a1 string) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.LoadBalancer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyServer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) DestroyServer(_a0 context.Context, _a1 string) (*brightbox.Server, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.Server, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.Server); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyServerGroup provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) DestroyServerGroup(_a0 context.Context, _a1 string) (*brightbox.ServerGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.ServerGroup, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.ServerGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirewallPolicies provides a mock function with given fields: _a0
func (_m *CloudAccess) FirewallPolicies(_a0 context.Context) ([]brightbox.FirewallPolicy, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.FirewallPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.FirewallPolicy, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.FirewallPolicy); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.FirewallPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Images provides a mock function with given fields: _a0
func (_m *CloudAccess) Images(_a0 context.Context) ([]brightbox.Image, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.Image, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.Image); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) LoadBalancer(_a0 context.Context, _a1 string) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.LoadBalancer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancers provides a mock function with given fields: _a0
func (_m *CloudAccess) LoadBalancers(_a0 context.Context) ([]brightbox.LoadBalancer, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.LoadBalancer, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.LoadBalancer); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MapCloudIP provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudAccess) MapCloudIP(_a0 context.Context, _a1 string, _a2 brightbox.CloudIPAttachment) (*brightbox.CloudIP, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.CloudIPAttachment) (*brightbox.CloudIP, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.CloudIPAttachment) *brightbox.CloudIP); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, brightbox.CloudIPAttachment) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveServersFromServerGroup provides a mock function with given fields: _a0, _a1, _a2
func (_m *CloudAccess) RemoveServersFromServerGroup(_a0 context.Context, _a1 string, _a2 brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.ServerGroupMemberList) (*brightbox.ServerGroup, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, brightbox.ServerGroupMemberList) *brightbox.ServerGroup); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, brightbox.ServerGroupMemberList) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Server provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) Server(_a0 context.Context, _a1 string) (*brightbox.Server, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.Server, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.Server); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroup provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) ServerGroup(_a0 context.Context, _a1 string) (*brightbox.ServerGroup, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.ServerGroup, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.ServerGroup); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroups provides a mock function with given fields: _a0
func (_m *CloudAccess) ServerGroups(_a0 context.Context) ([]brightbox.ServerGroup, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.ServerGroup, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.ServerGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerType provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) ServerType(_a0 context.Context, _a1 string) (*brightbox.ServerType, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.ServerType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.ServerType, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.ServerType); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.ServerType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerTypes provides a mock function with given fields: _a0
func (_m *CloudAccess) ServerTypes(_a0 context.Context) ([]brightbox.ServerType, error) {
	ret := _m.Called(_a0)

	var r0 []brightbox.ServerType
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]brightbox.ServerType, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []brightbox.ServerType); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]brightbox.ServerType)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnMapCloudIP provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) UnMapCloudIP(_a0 context.Context, _a1 string) (*brightbox.CloudIP, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*brightbox.CloudIP, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *brightbox.CloudIP); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFirewallRule provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) UpdateFirewallRule(_a0 context.Context, _a1 brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.FirewallRule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallRuleOptions) (*brightbox.FirewallRule, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.FirewallRuleOptions) *brightbox.FirewallRule); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.FirewallRule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.FirewallRuleOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLoadBalancer provides a mock function with given fields: _a0, _a1
func (_m *CloudAccess) UpdateLoadBalancer(_a0 context.Context, _a1 brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *brightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.LoadBalancerOptions) (*brightbox.LoadBalancer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, brightbox.LoadBalancerOptions) *brightbox.LoadBalancer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*brightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, brightbox.LoadBalancerOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCloudAccess interface {
	mock.TestingT
	Cleanup(func())
}

// NewCloudAccess creates a new instance of CloudAccess. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCloudAccess(t mockConstructorTestingTNewCloudAccess) *CloudAccess {
	mock := &CloudAccess{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
