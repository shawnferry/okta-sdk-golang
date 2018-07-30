/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
)

type SamlAttributeStatement struct {
	Name string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Type string `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
}

func (m *SamlAttributeStatement) WithName(v string) *SamlAttributeStatement {
	m.Name = v
	return m
}

func (m *SamlAttributeStatement) WithNamespace(v string) *SamlAttributeStatement {
	m.Namespace = v
	return m
}

func (m *SamlAttributeStatement) WithType(v string) *SamlAttributeStatement {
	m.Type = v
	return m
}

func (m *SamlAttributeStatement) WithValues(v []string) *SamlAttributeStatement {
	m.Values = v
	return m
}

