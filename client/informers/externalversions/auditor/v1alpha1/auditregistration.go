/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	auditorv1alpha1 "kubeshield.dev/auditor/apis/auditor/v1alpha1"
	versioned "kubeshield.dev/auditor/client/clientset/versioned"
	internalinterfaces "kubeshield.dev/auditor/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeshield.dev/auditor/client/listers/auditor/v1alpha1"
)

// AuditRegistrationInformer provides access to a shared informer and lister for
// AuditRegistrations.
type AuditRegistrationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AuditRegistrationLister
}

type auditRegistrationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAuditRegistrationInformer constructs a new informer for AuditRegistration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAuditRegistrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAuditRegistrationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAuditRegistrationInformer constructs a new informer for AuditRegistration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAuditRegistrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuditorV1alpha1().AuditRegistrations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AuditorV1alpha1().AuditRegistrations(namespace).Watch(context.TODO(), options)
			},
		},
		&auditorv1alpha1.AuditRegistration{},
		resyncPeriod,
		indexers,
	)
}

func (f *auditRegistrationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAuditRegistrationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *auditRegistrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&auditorv1alpha1.AuditRegistration{}, f.defaultInformer)
}

func (f *auditRegistrationInformer) Lister() v1alpha1.AuditRegistrationLister {
	return v1alpha1.NewAuditRegistrationLister(f.Informer().GetIndexer())
}
