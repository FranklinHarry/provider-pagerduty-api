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
	v1alpha1 "kubeform.dev/provider-pagerduty-api/apis/team/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MembershipLister helps list Memberships.
// All objects returned here must be treated as read-only.
type MembershipLister interface {
	// List lists all Memberships in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Membership, err error)
	// Memberships returns an object that can list and get Memberships.
	Memberships(namespace string) MembershipNamespaceLister
	MembershipListerExpansion
}

// membershipLister implements the MembershipLister interface.
type membershipLister struct {
	indexer cache.Indexer
}

// NewMembershipLister returns a new MembershipLister.
func NewMembershipLister(indexer cache.Indexer) MembershipLister {
	return &membershipLister{indexer: indexer}
}

// List lists all Memberships in the indexer.
func (s *membershipLister) List(selector labels.Selector) (ret []*v1alpha1.Membership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Membership))
	})
	return ret, err
}

// Memberships returns an object that can list and get Memberships.
func (s *membershipLister) Memberships(namespace string) MembershipNamespaceLister {
	return membershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MembershipNamespaceLister helps list and get Memberships.
// All objects returned here must be treated as read-only.
type MembershipNamespaceLister interface {
	// List lists all Memberships in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Membership, err error)
	// Get retrieves the Membership from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Membership, error)
	MembershipNamespaceListerExpansion
}

// membershipNamespaceLister implements the MembershipNamespaceLister
// interface.
type membershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Memberships in the indexer for a given namespace.
func (s membershipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Membership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Membership))
	})
	return ret, err
}

// Get retrieves the Membership from the indexer for a given namespace and name.
func (s membershipNamespaceLister) Get(name string) (*v1alpha1.Membership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("membership"), name)
	}
	return obj.(*v1alpha1.Membership), nil
}