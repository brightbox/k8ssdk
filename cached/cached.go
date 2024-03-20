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

package cached

import (
	"context"
	"time"

	brightbox "github.com/brightbox/gobrightbox/v2"
	cache "github.com/patrickmn/go-cache"
	klog "k8s.io/klog/v2"
)

const (
	expirationTime = 5 * time.Second
	purgeTime      = 30 * time.Second
)

// Client is a cached brightbox Client
type Client struct {
	clientCache *cache.Cache
	brightbox.Client
}

// Connect allocates and configures a Client for interacting with the API.
func Connect(ctx context.Context, config brightbox.Oauth2) (*Client, error) {
	cl, err := brightbox.Connect(ctx, config)
	if err != nil {
		return nil, err
	}
	return &Client{
		clientCache: cache.New(expirationTime, purgeTime),
		Client:      *cl,
	}, err
}

// Server fetches a server by id
func (c *Client) Server(ctx context.Context, identifier string) (*brightbox.Server, error) {
	if cachedServer, found := c.clientCache.Get(identifier); found {
		klog.V(4).Infof("Cache hit %q", identifier)
		return cachedServer.(*brightbox.Server), nil
	}
	server, err := c.Client.Server(ctx, identifier)
	if err != nil {
		return nil, err
	}
	klog.V(4).Infof("Cacheing %q", identifier)
	c.clientCache.Set(identifier, server, cache.DefaultExpiration)
	return server, nil
}

// ServerGroup fetches a server group by id
func (c *Client) ServerGroup(ctx context.Context, identifier string) (*brightbox.ServerGroup, error) {
	if cachedServerGroup, found := c.clientCache.Get(identifier); found {
		klog.V(4).Infof("Cache hit %q", identifier)
		return cachedServerGroup.(*brightbox.ServerGroup), nil
	}
	serverGroup, err := c.Client.ServerGroup(ctx, identifier)
	if err != nil {
		return nil, err
	}
	klog.V(4).Infof("Cacheing %q", identifier)
	c.clientCache.Set(identifier, serverGroup, cache.DefaultExpiration)
	return serverGroup, nil
}

// ConfigMap fetches a config map by id
func (c *Client) ConfigMap(ctx context.Context, identifier string) (*brightbox.ConfigMap, error) {
	if cachedConfigMap, found := c.clientCache.Get(identifier); found {
		klog.V(4).Infof("Cache hit %q", identifier)
		return cachedConfigMap.(*brightbox.ConfigMap), nil
	}
	configMap, err := c.Client.ConfigMap(ctx, identifier)
	if err != nil {
		return nil, err
	}
	klog.V(4).Infof("Cacheing %q", identifier)
	c.clientCache.Set(identifier, configMap, cache.DefaultExpiration)
	return configMap, nil
}

// DestroyServer removes a server by id
func (c *Client) DestroyServer(ctx context.Context, identifier string) (*brightbox.Server, error) {
	server, err := c.Client.DestroyServer(ctx, identifier)
	if err == nil {
		c.clientCache.Delete(identifier)
	}
	return server, err
}

// DestroyServerGroup removes a server group by id
func (c *Client) DestroyServerGroup(ctx context.Context, identifier string) (*brightbox.ServerGroup, error) {
	group, err := c.Client.DestroyServerGroup(ctx, identifier)
	if err == nil {
		c.clientCache.Delete(identifier)
	}
	return group, err
}
