/*
Copyright IBM Corp. 2016 All Rights Reserved.

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

package sharedconfig

import ab "github.com/hyperledger/fabric/protos/orderer"
import "time"

// Manager is a mock implementation of sharedconfig.Manager
type Manager struct {
	// ConsensusTypeVal is returned as the result of ConsensusType()
	ConsensusTypeVal string
	// BatchSizeVal is returned as the result of BatchSize()
	BatchSizeVal *ab.BatchSize
	// BatchTimeoutVal is returned as the result of BatchTimeout()
	BatchTimeoutVal time.Duration
	// ChainCreatorsVal is returned as the result of ChainCreators()
	ChainCreatorsVal []string
	// KafkaBrokersVal is returned as the result of KafkaBrokers()
	KafkaBrokersVal []string
	// IngressPolicyVal is returned as the result of IngressPolicy()
	IngressPolicyVal string
}

// ConsensusType returns the ConsensusTypeVal
func (scm *Manager) ConsensusType() string {
	return scm.ConsensusTypeVal
}

// BatchSize returns the BatchSizeVal
func (scm *Manager) BatchSize() *ab.BatchSize {
	return scm.BatchSizeVal
}

// BatchTimeout returns the BatchTimeoutVal
func (scm *Manager) BatchTimeout() time.Duration {
	return scm.BatchTimeoutVal
}

// ChainCreators returns the ChainCreatorsVal
func (scm *Manager) ChainCreators() []string {
	return scm.ChainCreatorsVal
}

// KafkaBrokers returns the KafkaBrokersVal
func (scm *Manager) KafkaBrokers() []string {
	return scm.KafkaBrokersVal
}

// IngressPolicy returns the IngressPolicyVal
func (scm *Manager) IngressPolicy() string {
	return scm.IngressPolicyVal
}
