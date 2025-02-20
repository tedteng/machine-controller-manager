/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package internalversion

import (
	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineDeploymentLister helps list MachineDeployments.
// All objects returned here must be treated as read-only.
type MachineDeploymentLister interface {
	// List lists all MachineDeployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.MachineDeployment, err error)
	// MachineDeployments returns an object that can list and get MachineDeployments.
	MachineDeployments(namespace string) MachineDeploymentNamespaceLister
	MachineDeploymentListerExpansion
}

// machineDeploymentLister implements the MachineDeploymentLister interface.
type machineDeploymentLister struct {
	indexer cache.Indexer
}

// NewMachineDeploymentLister returns a new MachineDeploymentLister.
func NewMachineDeploymentLister(indexer cache.Indexer) MachineDeploymentLister {
	return &machineDeploymentLister{indexer: indexer}
}

// List lists all MachineDeployments in the indexer.
func (s *machineDeploymentLister) List(selector labels.Selector) (ret []*machine.MachineDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineDeployment))
	})
	return ret, err
}

// MachineDeployments returns an object that can list and get MachineDeployments.
func (s *machineDeploymentLister) MachineDeployments(namespace string) MachineDeploymentNamespaceLister {
	return machineDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineDeploymentNamespaceLister helps list and get MachineDeployments.
// All objects returned here must be treated as read-only.
type MachineDeploymentNamespaceLister interface {
	// List lists all MachineDeployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*machine.MachineDeployment, err error)
	// Get retrieves the MachineDeployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*machine.MachineDeployment, error)
	MachineDeploymentNamespaceListerExpansion
}

// machineDeploymentNamespaceLister implements the MachineDeploymentNamespaceLister
// interface.
type machineDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineDeployments in the indexer for a given namespace.
func (s machineDeploymentNamespaceLister) List(selector labels.Selector) (ret []*machine.MachineDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*machine.MachineDeployment))
	})
	return ret, err
}

// Get retrieves the MachineDeployment from the indexer for a given namespace and name.
func (s machineDeploymentNamespaceLister) Get(name string) (*machine.MachineDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(machine.Resource("machinedeployment"), name)
	}
	return obj.(*machine.MachineDeployment), nil
}
