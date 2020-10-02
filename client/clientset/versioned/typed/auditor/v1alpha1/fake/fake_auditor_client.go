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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeshield.dev/auditor/client/clientset/versioned/typed/auditor/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAuditorV1alpha1 struct {
	*testing.Fake
}

func (c *FakeAuditorV1alpha1) AuditRegistrations(namespace string) v1alpha1.AuditRegistrationInterface {
	return &FakeAuditRegistrations{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuditorV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
