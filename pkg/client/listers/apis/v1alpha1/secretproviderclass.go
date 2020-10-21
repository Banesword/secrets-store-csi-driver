/*
Copyright 2020 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1alpha1"
)

// SecretProviderClassLister helps list SecretProviderClasses.
type SecretProviderClassLister interface {
	// List lists all SecretProviderClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClass, err error)
	// SecretProviderClasses returns an object that can list and get SecretProviderClasses.
	SecretProviderClasses(namespace string) SecretProviderClassNamespaceLister
	SecretProviderClassListerExpansion
}

// secretProviderClassLister implements the SecretProviderClassLister interface.
type secretProviderClassLister struct {
	indexer cache.Indexer
}

// NewSecretProviderClassLister returns a new SecretProviderClassLister.
func NewSecretProviderClassLister(indexer cache.Indexer) SecretProviderClassLister {
	return &secretProviderClassLister{indexer: indexer}
}

// List lists all SecretProviderClasses in the indexer.
func (s *secretProviderClassLister) List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretProviderClass))
	})
	return ret, err
}

// SecretProviderClasses returns an object that can list and get SecretProviderClasses.
func (s *secretProviderClassLister) SecretProviderClasses(namespace string) SecretProviderClassNamespaceLister {
	return secretProviderClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretProviderClassNamespaceLister helps list and get SecretProviderClasses.
type SecretProviderClassNamespaceLister interface {
	// List lists all SecretProviderClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClass, err error)
	// Get retrieves the SecretProviderClass from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecretProviderClass, error)
	SecretProviderClassNamespaceListerExpansion
}

// secretProviderClassNamespaceLister implements the SecretProviderClassNamespaceLister
// interface.
type secretProviderClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretProviderClasses in the indexer for a given namespace.
func (s secretProviderClassNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretProviderClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretProviderClass))
	})
	return ret, err
}

// Get retrieves the SecretProviderClass from the indexer for a given namespace and name.
func (s secretProviderClassNamespaceLister) Get(name string) (*v1alpha1.SecretProviderClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretproviderclass"), name)
	}
	return obj.(*v1alpha1.SecretProviderClass), nil
}