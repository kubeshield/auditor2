/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeshield.dev/auditor/apis/auditor/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DashboardTemplateLister helps list DashboardTemplates.
type DashboardTemplateLister interface {
	// List lists all DashboardTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DashboardTemplate, err error)
	// DashboardTemplates returns an object that can list and get DashboardTemplates.
	DashboardTemplates(namespace string) DashboardTemplateNamespaceLister
	DashboardTemplateListerExpansion
}

// dashboardTemplateLister implements the DashboardTemplateLister interface.
type dashboardTemplateLister struct {
	indexer cache.Indexer
}

// NewDashboardTemplateLister returns a new DashboardTemplateLister.
func NewDashboardTemplateLister(indexer cache.Indexer) DashboardTemplateLister {
	return &dashboardTemplateLister{indexer: indexer}
}

// List lists all DashboardTemplates in the indexer.
func (s *dashboardTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.DashboardTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DashboardTemplate))
	})
	return ret, err
}

// DashboardTemplates returns an object that can list and get DashboardTemplates.
func (s *dashboardTemplateLister) DashboardTemplates(namespace string) DashboardTemplateNamespaceLister {
	return dashboardTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DashboardTemplateNamespaceLister helps list and get DashboardTemplates.
type DashboardTemplateNamespaceLister interface {
	// List lists all DashboardTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DashboardTemplate, err error)
	// Get retrieves the DashboardTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DashboardTemplate, error)
	DashboardTemplateNamespaceListerExpansion
}

// dashboardTemplateNamespaceLister implements the DashboardTemplateNamespaceLister
// interface.
type dashboardTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DashboardTemplates in the indexer for a given namespace.
func (s dashboardTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DashboardTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DashboardTemplate))
	})
	return ret, err
}

// Get retrieves the DashboardTemplate from the indexer for a given namespace and name.
func (s dashboardTemplateNamespaceLister) Get(name string) (*v1alpha1.DashboardTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dashboardtemplate"), name)
	}
	return obj.(*v1alpha1.DashboardTemplate), nil
}
