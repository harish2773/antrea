/*
Copyright 2021 Antrea Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	multiclusterv1alpha1 "antrea.io/antrea/multicluster/apis/multicluster/v1alpha1"
	versioned "antrea.io/antrea/multicluster/pkg/client/clientset/versioned"
	internalinterfaces "antrea.io/antrea/multicluster/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "antrea.io/antrea/multicluster/pkg/client/listers/multicluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceExportInformer provides access to a shared informer and lister for
// ResourceExports.
type ResourceExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceExportLister
}

type resourceExportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceExportInformer constructs a new informer for ResourceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceExportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceExportInformer constructs a new informer for ResourceExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1alpha1().ResourceExports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterV1alpha1().ResourceExports(namespace).Watch(context.TODO(), options)
			},
		},
		&multiclusterv1alpha1.ResourceExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceExportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceExportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceExportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiclusterv1alpha1.ResourceExport{}, f.defaultInformer)
}

func (f *resourceExportInformer) Lister() v1alpha1.ResourceExportLister {
	return v1alpha1.NewResourceExportLister(f.Informer().GetIndexer())
}
