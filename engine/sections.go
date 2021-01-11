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

const (
	tplHeader = `
	{{- if .Settings.ShowHeader -}}
		{{- with .Module.Header -}}
			{{ . }}
			{{ printf "\n" }}
		{{- end -}}
	{{ end -}}
	`

	tplInputs = `
	{{- if .Settings.ShowInputs -}}
		{{ indent 0 "#" }} Inputs
		{{ if not .Module.Inputs }}
			No inputs.
		{{ else }}
			{{- range .Module.Inputs }}
			{{ printf "\n" }}
			{{ indent 1 "#" }} {{ name .Name }} ` + "`" + `{{ .Type }}` + "`" + `

			Description: {{ .Description }}
			Default: {{ .Default }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tplOutputs = `
	{{- if .Settings.ShowOutputs -}}
		{{ indent 0 "#" }} Outputs
		{{ if not .Module.Outputs }}
			No outputs.
		{{ else }}
			| Name | Description | Type | Default |
			|------|-------------|------|---------|
			{{- range .Module.Outputs }}
				- {{ .Name }}
				- {{ .Description }}
			{{- end }}
		{{ end }}
	{{ end -}}
	`

	tplRequirements = `
	{{- if .Settings.ShowRequirements -}}
	{{ indent 0 "#" }} Requirements
	
	Intentionally don't show requirements to demonstrate the flexibility of the plugin.
	{{ end -}}
	`

	tplResources = `
	{{- if .Settings.ShowResources -}}
	{{ indent 0 "#" }} Resources
		{{ if not .Module.Resources }}
			No resources.
		{{ else }}
			{{- range .Module.Resources }}
				- {{ .Name }}
			{{- end }}
		{{- end }}
	{{ end -}}
	`

	tplProviders = `
	{{- if .Settings.ShowProviders -}}
		{{ indent 0 "#" }} Providers
		{{ if not .Module.Providers }}
			No providers.
		{{ else }}
			{{- range .Module.Providers }}
				- {{ .ProviderName }}
			{{- end }}
		{{- end }}
	{{ end -}}
	`

	tplCustom = `
	{{- template "header" . -}}
	{{- template "inputs" . -}}
	{{- template "outputs" . -}}
	{{- template "requirements" . -}}
	{{- template "providers" . -}}
	{{- template "resources" . -}}
	`
)
