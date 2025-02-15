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

package v1beta1

import (
	"testing"
)

func TestHcloudCluster_ValidateUpdate(t *testing.T) {
	tests := []struct {
		name       string
		oldCluster *HcloudCluster
		newCluster *HcloudCluster
		wantErr    bool
	}{
		{
			name: "locations is immutable",
			oldCluster: &HcloudCluster{
				Spec: HcloudClusterSpec{
					Locations: []HcloudLocation{"nbg1"},
				},
			},
			newCluster: &HcloudCluster{
				Spec: HcloudClusterSpec{
					Locations: []HcloudLocation{"us-east-2"},
				},
			},
			wantErr: true,
		},
		{
			name: "reordering is allowed",
			oldCluster: &HcloudCluster{
				Spec: HcloudClusterSpec{
					Locations: []HcloudLocation{"fsn1", "nbg1"},
				},
			},
			newCluster: &HcloudCluster{
				Spec: HcloudClusterSpec{
					Locations: []HcloudLocation{"nbg1", "fsn1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.newCluster.ValidateUpdate(tt.oldCluster); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
