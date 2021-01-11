/*
Copyright 2021 The terraform-docs Authors.

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

package engine

import (
	"github.com/terraform-docs/plugin-sdk/print"
	"github.com/terraform-docs/plugin-sdk/template"
	"github.com/terraform-docs/plugin-sdk/terraform"
)

type engine struct{}

// New returns an implementation which satisfies print.Engine interface.
func New() print.Engine {
	return &engine{}
}

// Print the custom format template. You have all the flexibility to generate
// the output however you choose to.
func (e *engine) Print(module terraform.Module, settings *print.Settings) (string, error) {
	tpl := template.New(settings,
		&template.Item{Name: "custom", Text: tplCustom},
		&template.Item{Name: "header", Text: tplHeader},
		&template.Item{Name: "inputs", Text: tplInputs},
		&template.Item{Name: "outputs", Text: tplOutputs},
		&template.Item{Name: "requirements", Text: tplRequirements},
		&template.Item{Name: "providers", Text: tplProviders},
	)

	rendered, err := tpl.Render(module)
	if err != nil {
		return "", err
	}
	return rendered, nil
}
