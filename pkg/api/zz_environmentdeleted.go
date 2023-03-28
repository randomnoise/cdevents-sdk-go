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
	_ "embed"
	"fmt"
	"time"
)

//go:embed spec/schemas/environmentdeleted.json
var environmentdeletedschema string

var (
	// EnvironmentDeleted event v0.1.1
	EnvironmentDeletedEventV1 CDEventType = CDEventType{
		Subject:   "environment",
		Predicate: "deleted",
		Version:   "0.1.1",
	}
)

type EnvironmentDeletedSubjectContent struct {
	Name string `json:"name"`
}

type EnvironmentDeletedSubject struct {
	SubjectBase
	Content EnvironmentDeletedSubjectContent `json:"content"`
}

func (sc EnvironmentDeletedSubject) GetSubjectType() SubjectType {
	return "environment"
}

type EnvironmentDeletedEvent struct {
	Context Context                   `json:"context"`
	Subject EnvironmentDeletedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e EnvironmentDeletedEvent) GetType() CDEventType {
	return EnvironmentDeletedEventV1
}

func (e EnvironmentDeletedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e EnvironmentDeletedEvent) GetId() string {
	return e.Context.Id
}

func (e EnvironmentDeletedEvent) GetSource() string {
	return e.Context.Source
}

func (e EnvironmentDeletedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e EnvironmentDeletedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e EnvironmentDeletedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e EnvironmentDeletedEvent) GetSubject() Subject {
	return e.Subject
}

func (e EnvironmentDeletedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentDeletedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e EnvironmentDeletedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e EnvironmentDeletedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *EnvironmentDeletedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *EnvironmentDeletedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *EnvironmentDeletedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *EnvironmentDeletedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *EnvironmentDeletedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *EnvironmentDeletedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e EnvironmentDeletedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), environmentdeletedschema
}

// Set subject custom fields

func (e *EnvironmentDeletedEvent) SetSubjectName(name string) {
	e.Subject.Content.Name = name
}

// New creates a new EnvironmentDeletedEvent
func NewEnvironmentDeletedEvent() (*EnvironmentDeletedEvent, error) {
	e := &EnvironmentDeletedEvent{
		Context: Context{
			Type:    EnvironmentDeletedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: EnvironmentDeletedSubject{
			SubjectBase: SubjectBase{
				Type: "environment",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}
