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

var testsuitefinishedschema = `{"$schema":"https://json-schema.org/draft/2020-12/schema","$id":"https://cdevents.dev/0.2.0/schema/test-suite-finished-event","properties":{"context":{"properties":{"version":{"type":"string","minLength":1},"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","enum":["dev.cdevents.testsuite.finished.0.1.1"],"default":"dev.cdevents.testsuite.finished.0.1.1"},"timestamp":{"type":"string","format":"date-time"}},"additionalProperties":false,"type":"object","required":["version","id","source","type","timestamp"]},"subject":{"properties":{"id":{"type":"string","minLength":1},"source":{"type":"string","minLength":1,"format":"uri-reference"},"type":{"type":"string","minLength":1,"enum":["testSuite"],"default":"testSuite"},"content":{"properties":{},"additionalProperties":false,"type":"object"}},"additionalProperties":false,"type":"object","required":["id","type","content"]},"customData":{"oneOf":[{"type":"object"},{"type":"string","contentEncoding":"base64"}]},"customDataContentType":{"type":"string"}},"additionalProperties":false,"type":"object","required":["context","subject"]}`

var (
	// TestSuiteFinished event v0.1.1
	TestSuiteFinishedEventV1 CDEventType = CDEventType{
		Subject:   "testsuite",
		Predicate: "finished",
		Version:   "0.1.1",
	}
)

type TestSuiteFinishedSubjectContent struct {
}

type TestSuiteFinishedSubject struct {
	SubjectBase
	Content TestSuiteFinishedSubjectContent `json:"content"`
}

func (sc TestSuiteFinishedSubject) GetSubjectType() SubjectType {
	return "testSuite"
}

type TestSuiteFinishedEvent struct {
	Context Context                  `json:"context"`
	Subject TestSuiteFinishedSubject `json:"subject"`
	CDEventCustomData
}

// CDEventsReader implementation

func (e TestSuiteFinishedEvent) GetType() CDEventType {
	return TestSuiteFinishedEventV1
}

func (e TestSuiteFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestSuiteFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e TestSuiteFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestSuiteFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestSuiteFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestSuiteFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestSuiteFinishedEvent) GetSubject() Subject {
	return e.Subject
}

func (e TestSuiteFinishedEvent) GetCustomData() (interface{}, error) {
	return GetCustomData(e.CustomDataContentType, e.CustomData)
}

func (e TestSuiteFinishedEvent) GetCustomDataAs(receiver interface{}) error {
	return GetCustomDataAs(e, receiver)
}

func (e TestSuiteFinishedEvent) GetCustomDataRaw() ([]byte, error) {
	return GetCustomDataRaw(e.CustomDataContentType, e.CustomData)
}

func (e TestSuiteFinishedEvent) GetCustomDataContentType() string {
	return e.CustomDataContentType
}

// CDEventsWriter implementation

func (e *TestSuiteFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestSuiteFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestSuiteFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestSuiteFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestSuiteFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func (e *TestSuiteFinishedEvent) SetCustomData(contentType string, data interface{}) error {
	err := CheckCustomData(contentType, data)
	if err != nil {
		return err
	}
	e.CustomData = data
	e.CustomDataContentType = contentType
	return nil
}

func (e TestSuiteFinishedEvent) GetSchema() (string, string) {
	eType := e.GetType()
	return fmt.Sprintf(CDEventsSchemaURLTemplate, CDEventsSpecVersion, eType.Subject, eType.Predicate), testsuitefinishedschema
}

// Set subject custom fields

// New creates a new TestSuiteFinishedEvent
func NewTestSuiteFinishedEvent() (*TestSuiteFinishedEvent, error) {
	e := &TestSuiteFinishedEvent{
		Context: Context{
			Type:    TestSuiteFinishedEventV1.String(),
			Version: CDEventsSpecVersion,
		},
		Subject: TestSuiteFinishedSubject{
			SubjectBase: SubjectBase{
				Type: "testSuite",
			},
		},
	}
	_, err := initCDEvent(e)
	if err != nil {
		return nil, err
	}
	return e, nil
}