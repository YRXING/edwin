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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	edwinv1 "github.com/YRXING/edwin/pkg/apis/edwin/v1"
	versioned "github.com/YRXING/edwin/pkg/client/clientset/versioned"
	internalinterfaces "github.com/YRXING/edwin/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/YRXING/edwin/pkg/client/listers/edwin/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PacketFilterInformer provides access to a shared informer and lister for
// PacketFilters.
type PacketFilterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PacketFilterLister
}

type packetFilterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPacketFilterInformer constructs a new informer for PacketFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPacketFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPacketFilterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPacketFilterInformer constructs a new informer for PacketFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPacketFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdwinV1().PacketFilters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdwinV1().PacketFilters(namespace).Watch(context.TODO(), options)
			},
		},
		&edwinv1.PacketFilter{},
		resyncPeriod,
		indexers,
	)
}

func (f *packetFilterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPacketFilterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *packetFilterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edwinv1.PacketFilter{}, f.defaultInformer)
}

func (f *packetFilterInformer) Lister() v1.PacketFilterLister {
	return v1.NewPacketFilterLister(f.Informer().GetIndexer())
}
