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
	xiaofengv1 "carrier/pkg/apis/xiaofeng/v1"
	versioned "carrier/pkg/generated/clientset/versioned"
	internalinterfaces "carrier/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "carrier/pkg/generated/listers/xiaofeng/v1"
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CarrierInformer provides access to a shared informer and lister for
// Carriers.
type CarrierInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CarrierLister
}

type carrierInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCarrierInformer constructs a new informer for Carrier type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCarrierInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCarrierInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCarrierInformer constructs a new informer for Carrier type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCarrierInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.XiaofengV1().Carriers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.XiaofengV1().Carriers(namespace).Watch(context.TODO(), options)
			},
		},
		&xiaofengv1.Carrier{},
		resyncPeriod,
		indexers,
	)
}

func (f *carrierInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCarrierInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *carrierInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&xiaofengv1.Carrier{}, f.defaultInformer)
}

func (f *carrierInformer) Lister() v1.CarrierLister {
	return v1.NewCarrierLister(f.Informer().GetIndexer())
}
