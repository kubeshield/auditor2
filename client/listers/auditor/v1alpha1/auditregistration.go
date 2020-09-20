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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeshield.dev/auditor/apis/auditor/v1alpha1"
)

// AuditRegistrationLister helps list AuditRegistrations.
type AuditRegistrationLister interface {
	// List lists all AuditRegistrations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AuditRegistration, err error)
	// AuditRegistrations returns an object that can list and get AuditRegistrations.
	AuditRegistrations(namespace string) AuditRegistrationNamespaceLister
	AuditRegistrationListerExpansion
}

// auditRegistrationLister implements the AuditRegistrationLister interface.
type auditRegistrationLister struct {
	indexer cache.Indexer
}

// NewAuditRegistrationLister returns a new AuditRegistrationLister.
func NewAuditRegistrationLister(indexer cache.Indexer) AuditRegistrationLister {
	return &auditRegistrationLister{indexer: indexer}
}

// List lists all AuditRegistrations in the indexer.
func (s *auditRegistrationLister) List(selector labels.Selector) (ret []*v1alpha1.AuditRegistration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuditRegistration))
	})
	return ret, err
}

// AuditRegistrations returns an object that can list and get AuditRegistrations.
func (s *auditRegistrationLister) AuditRegistrations(namespace string) AuditRegistrationNamespaceLister {
	return auditRegistrationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuditRegistrationNamespaceLister helps list and get AuditRegistrations.
type AuditRegistrationNamespaceLister interface {
	// List lists all AuditRegistrations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AuditRegistration, err error)
	// Get retrieves the AuditRegistration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AuditRegistration, error)
	AuditRegistrationNamespaceListerExpansion
}

// auditRegistrationNamespaceLister implements the AuditRegistrationNamespaceLister
// interface.
type auditRegistrationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuditRegistrations in the indexer for a given namespace.
func (s auditRegistrationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AuditRegistration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AuditRegistration))
	})
	return ret, err
}

// Get retrieves the AuditRegistration from the indexer for a given namespace and name.
func (s auditRegistrationNamespaceLister) Get(name string) (*v1alpha1.AuditRegistration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("auditregistration"), name)
	}
	return obj.(*v1alpha1.AuditRegistration), nil
}
