{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "wallet-settings.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "wallet-settings.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "wallet-settings.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "wallet-settings.labels" -}}
helm.sh/chart: {{ include "wallet-settings.chart" . }}
{{ include "wallet-settings.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "wallet-settings.selectorLabels" -}}
app.kubernetes.io/name: {{ include "wallet-settings.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "wallet-settings.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "wallet-settings.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create tag name of the image
*/}}
{{- define "wallet-settings.imageTag" -}}
{{ .Values.image.tag | default .Chart.AppVersion }}
{{- end }}

{{/*
Create the name of the image repository
*/}}
{{- define "wallet-settings.imageRepository" -}}
{{ .Values.image.repository | default (printf "velmie/%s" .Chart.Name) }}
{{- end }}

{{/*
Create full image repository name including tag
*/}}
{{- define "wallet-settings.imageRepositoryWithTag" -}}
{{ include "wallet-settings.imageRepository" . }}:{{ include "wallet-settings.imageTag" . }}
{{- end }}

{{/*
Create full database migration image repository name
*/}}
{{- define "wallet-settings.dbMigrationImageRepositoryWithTag" -}}
{{ include "wallet-settings.imageRepository" . }}-db-migration:{{ include "wallet-settings.imageTag" . }}
{{- end }}