// Code generated by tools/generator. DO NOT EDIT.

/*
Copyright 2023 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"fmt"
	"time"
)

var incidentresolvedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/incident-resolved-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.incident.resolved.0.1.0"],"default":"dev.cdevents.incident.resolved.0.1.0"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["incident"],"default":"incident"},"content":{"properties":{"description":{"type":"string"},"environment":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"service":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"}},"additionalProperties":false,"type":"object","required":["id"]},"artifactId":{"type":"string","minLength":1}},"additionalProperties":false,"type":"object","required":["environment"]}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// IncidentResolved event v0.1.0
	IncidentResolvedEventV1 CDEventType = CDEventType{
		Subject:   "incident",
		Predicate: "resolved",
		Version:   "0.1.0",
	}
)

type IncidentResolvedSubjectContent struct {
	ArtifactId string `json:"artifactId" validate:"purl"`

	Description string `json:"description"`

	Environment Reference `json:"environment"`

	Service Reference `json:"service"`
}

type IncidentResolvedSubject struct {
	SubjectBase
	Content IncidentResolvedSubjectContent `json:"content"`
}

func (sc IncidentResolvedSubject) GetSubjectType() SubjectType {
	return "incident"
}

type IncidentResolvedEvent struct {
	Context Context                 `json:"context"`
	Subject IncidentResolvedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e IncidentResolvedEvent) GetType() CDEventType {
	return IncidentResolvedEventV1
}

func (e IncidentResolvedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e IncidentResolvedEvent) GetId() string {
	return e.Context.Id
}

func (e IncidentResolvedEvent) GetSource() string {
	return e.Context.Source
}

func (e IncidentResolvedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e IncidentResolvedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e IncidentResolvedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e IncidentResolvedEvent) GetSubject() Subject {
	return e.Subject
}

func (e IncidentResolvedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e IncidentResolvedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e IncidentResolvedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e IncidentResolvedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *IncidentResolvedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *IncidentResolvedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *IncidentResolvedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *IncidentResolvedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *IncidentResolvedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *IncidentResolvedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e IncidentResolvedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), incidentresolvedschema
}

// Set subject custom fields

func (e *IncidentResolvedEvent) SetSubjectArtifactId(artifactId string) {
	e.Subject.Content.ArtifactId = artifactId
}

func (e *IncidentResolvedEvent) SetSubjectDescription(description string) {
	e.Subject.Content.Description = description
}

func (e *IncidentResolvedEvent) SetSubjectEnvironment(environment Reference) {
	e.Subject.Content.Environment = environment
}

func (e *IncidentResolvedEvent) SetSubjectService(service Reference) {
	e.Subject.Content.Service = service
}

// New creates a new IncidentResolvedEvent
func NewIncidentResolvedEvent() (*IncidentResolvedEvent, error) {
	e := &IncidentResolvedEvent{
		Context: Context{
			Type:    IncidentResolvedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: IncidentResolvedSubject{
			SubjectBase: SubjectBase{
				Type: "incident",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}