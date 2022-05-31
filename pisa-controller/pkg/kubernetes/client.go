// Copyright 2022 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sync"
)

var client *KClient
var once sync.Once
var initErr error

func GetInClusterClient() (*KClient, error) {
	once.Do(func() {
		config, err := rest.InClusterConfig()
		if err != nil {
			initErr = err
			return
		}
		clientset, err := dynamic.NewForConfig(config)
		if err != nil {
			initErr = err
			return
		}
		client = &KClient{}
		client.Client = clientset
	})

	return client, initErr
}