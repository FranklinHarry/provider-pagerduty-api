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
	v1alpha1 "kubeform.dev/provider-pagerduty-api/apis/maintenance/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WindowLister helps list Windows.
// All objects returned here must be treated as read-only.
type WindowLister interface {
	// List lists all Windows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Window, err error)
	// Windows returns an object that can list and get Windows.
	Windows(namespace string) WindowNamespaceLister
	WindowListerExpansion
}

// windowLister implements the WindowLister interface.
type windowLister struct {
	indexer cache.Indexer
}

// NewWindowLister returns a new WindowLister.
func NewWindowLister(indexer cache.Indexer) WindowLister {
	return &windowLister{indexer: indexer}
}

// List lists all Windows in the indexer.
func (s *windowLister) List(selector labels.Selector) (ret []*v1alpha1.Window, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Window))
	})
	return ret, err
}

// Windows returns an object that can list and get Windows.
func (s *windowLister) Windows(namespace string) WindowNamespaceLister {
	return windowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WindowNamespaceLister helps list and get Windows.
// All objects returned here must be treated as read-only.
type WindowNamespaceLister interface {
	// List lists all Windows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Window, err error)
	// Get retrieves the Window from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Window, error)
	WindowNamespaceListerExpansion
}

// windowNamespaceLister implements the WindowNamespaceLister
// interface.
type windowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Windows in the indexer for a given namespace.
func (s windowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Window, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Window))
	})
	return ret, err
}

// Get retrieves the Window from the indexer for a given namespace and name.
func (s windowNamespaceLister) Get(name string) (*v1alpha1.Window, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("window"), name)
	}
	return obj.(*v1alpha1.Window), nil
}