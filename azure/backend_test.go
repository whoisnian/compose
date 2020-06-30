/*
   Copyright 2020 Docker, Inc.

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

package azure

import (
	"testing"

	"github.com/stretchr/testify/suite"

	. "github.com/onsi/gomega"
)

type BackendSuiteTest struct {
	suite.Suite
}

func (suite *BackendSuiteTest) TestGetContainerName() {
	group, container := getGroupAndContainerName("docker1234")
	Expect(group).To(Equal("docker1234"))
	Expect(container).To(Equal(singleContainerName))

	group, container = getGroupAndContainerName("compose_service1")
	Expect(group).To(Equal("compose"))
	Expect(container).To(Equal("service1"))

	group, container = getGroupAndContainerName("compose_stack_service1")
	Expect(group).To(Equal("compose_stack"))
	Expect(container).To(Equal("service1"))
}

func TestBackendSuite(t *testing.T) {
	RegisterTestingT(t)
	suite.Run(t, new(BackendSuiteTest))
}