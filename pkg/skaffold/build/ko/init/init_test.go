/*
Copyright 2022 The Skaffold Authors

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

package init

import (
	"testing"

	"github.com/GoogleContainerTools/skaffold/testutil"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		description string
		path        string
		expect      bool
	}{
		{
			description: "go.mod in current directory",
			path:        "go.mod",
			expect:      true,
		},
		{
			description: "go.mod in subdirectory",
			path:        "foo/go.mod",
			expect:      true,
		},
		{
			description: "not go.mod",
			path:        "Gopkg.toml",
			expect:      false,
		},
	}
	for _, test := range tests {
		testutil.Run(t, test.description, func(t *testutil.T) {
			t.CheckDeepEqual(test.expect, Validate(test.path))
		})
	}
}
