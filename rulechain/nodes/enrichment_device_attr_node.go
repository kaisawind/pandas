//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package nodes

type enrichmentDeviceAttrNodeFactory struct{}

func (f enrichmentDeviceAttrNodeFactory) Name() string     { return "EnrichmentDeviceAttrbute" }
func (f enrichmentDeviceAttrNodeFactory) Category() string { return NODE_CATEGORY_ENRICHMENT }
func (f enrichmentDeviceAttrNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}
