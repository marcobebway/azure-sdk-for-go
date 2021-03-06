// +build go1.9

// Copyright 2021 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package policy

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Type = original.Type

const (
	BuiltIn      Type = original.BuiltIn
	Custom       Type = original.Custom
	NotSpecified Type = original.NotSpecified
)

type Assignment = original.Assignment
type AssignmentListResult = original.AssignmentListResult
type AssignmentListResultIterator = original.AssignmentListResultIterator
type AssignmentListResultPage = original.AssignmentListResultPage
type AssignmentProperties = original.AssignmentProperties
type AssignmentsClient = original.AssignmentsClient
type BaseClient = original.BaseClient
type Definition = original.Definition
type DefinitionListResult = original.DefinitionListResult
type DefinitionListResultIterator = original.DefinitionListResultIterator
type DefinitionListResultPage = original.DefinitionListResultPage
type DefinitionProperties = original.DefinitionProperties
type DefinitionsClient = original.DefinitionsClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAssignmentListResultIterator(page AssignmentListResultPage) AssignmentListResultIterator {
	return original.NewAssignmentListResultIterator(page)
}
func NewAssignmentListResultPage(cur AssignmentListResult, getNextPage func(context.Context, AssignmentListResult) (AssignmentListResult, error)) AssignmentListResultPage {
	return original.NewAssignmentListResultPage(cur, getNextPage)
}
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDefinitionListResultIterator(page DefinitionListResultPage) DefinitionListResultIterator {
	return original.NewDefinitionListResultIterator(page)
}
func NewDefinitionListResultPage(cur DefinitionListResult, getNextPage func(context.Context, DefinitionListResult) (DefinitionListResult, error)) DefinitionListResultPage {
	return original.NewDefinitionListResultPage(cur, getNextPage)
}
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClient(subscriptionID)
}
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return original.NewDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2017-03-09"
}
func Version() string {
	return original.Version()
}
