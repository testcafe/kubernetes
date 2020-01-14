/*
Copyright 2015 The Kubernetes Authors.

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

package oom

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"

	"github.com/google/cadvisor/utils/oomparser"
	"github.com/stretchr/testify/assert"
)

// TestBasic verifies that the OOMWatch works without error.
func TestBasic(t *testing.T) {
	fakeRecorder := &record.FakeRecorder{}
	node := &v1.ObjectReference{}

	// TODO: Substitute this `oomStreamer` out for a fake, and then write
	// more comprehensive unit tests of the actual behavior.
	oomStreamer, err := oomparser.New()
	assert.NoError(t, err)

	oomWatcher := &realWatcher{
		recorder:    fakeRecorder,
		oomStreamer: oomStreamer,
	}
	assert.NoError(t, oomWatcher.Start(node))
}