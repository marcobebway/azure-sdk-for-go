package trainingapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.1/customvision/training"
	"github.com/Azure/go-autorest/autorest"
	"github.com/gofrs/uuid"
	"io"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CreateImageRegions(ctx context.Context, projectID uuid.UUID, batch training.ImageRegionCreateBatch) (result training.ImageRegionCreateSummary, err error)
	CreateImagesFromData(ctx context.Context, projectID uuid.UUID, imageData io.ReadCloser, tagIds []string) (result training.ImageCreateSummary, err error)
	CreateImagesFromFiles(ctx context.Context, projectID uuid.UUID, batch training.ImageFileCreateBatch) (result training.ImageCreateSummary, err error)
	CreateImagesFromPredictions(ctx context.Context, projectID uuid.UUID, batch training.ImageIDCreateBatch) (result training.ImageCreateSummary, err error)
	CreateImagesFromUrls(ctx context.Context, projectID uuid.UUID, batch training.ImageURLCreateBatch) (result training.ImageCreateSummary, err error)
	CreateImageTags(ctx context.Context, projectID uuid.UUID, batch training.ImageTagCreateBatch) (result training.ImageTagCreateSummary, err error)
	CreateProject(ctx context.Context, name string, description string, domainID *uuid.UUID, classificationType string) (result training.Project, err error)
	CreateTag(ctx context.Context, projectID uuid.UUID, name string, description string) (result training.Tag, err error)
	DeleteImageRegions(ctx context.Context, projectID uuid.UUID, regionIds []string) (result autorest.Response, err error)
	DeleteImages(ctx context.Context, projectID uuid.UUID, imageIds []string) (result autorest.Response, err error)
	DeleteImageTags(ctx context.Context, projectID uuid.UUID, imageIds []string, tagIds []string) (result autorest.Response, err error)
	DeleteIteration(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID) (result autorest.Response, err error)
	DeletePrediction(ctx context.Context, projectID uuid.UUID, ids []string) (result autorest.Response, err error)
	DeleteProject(ctx context.Context, projectID uuid.UUID) (result autorest.Response, err error)
	DeleteTag(ctx context.Context, projectID uuid.UUID, tagID uuid.UUID) (result autorest.Response, err error)
	ExportIteration(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID, platform string, flavor string) (result training.Export, err error)
	GetDomain(ctx context.Context, domainID uuid.UUID) (result training.Domain, err error)
	GetDomains(ctx context.Context) (result training.ListDomain, err error)
	GetExports(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID) (result training.ListExport, err error)
	GetImagePerformanceCount(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID, tagIds []string) (result training.Int32, err error)
	GetImagePerformances(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID, tagIds []string, orderBy string, take *int32, skip *int32) (result training.ListImagePerformance, err error)
	GetImageRegionProposals(ctx context.Context, projectID uuid.UUID, imageID uuid.UUID) (result training.ImageRegionProposal, err error)
	GetImagesByIds(ctx context.Context, projectID uuid.UUID, imageIds []string, iterationID *uuid.UUID) (result training.ListImage, err error)
	GetIteration(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID) (result training.Iteration, err error)
	GetIterationPerformance(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID, threshold *float64, overlapThreshold *float64) (result training.IterationPerformance, err error)
	GetIterations(ctx context.Context, projectID uuid.UUID) (result training.ListIteration, err error)
	GetProject(ctx context.Context, projectID uuid.UUID) (result training.Project, err error)
	GetProjects(ctx context.Context) (result training.ListProject, err error)
	GetTag(ctx context.Context, projectID uuid.UUID, tagID uuid.UUID, iterationID *uuid.UUID) (result training.Tag, err error)
	GetTaggedImageCount(ctx context.Context, projectID uuid.UUID, iterationID *uuid.UUID, tagIds []string) (result training.Int32, err error)
	GetTaggedImages(ctx context.Context, projectID uuid.UUID, iterationID *uuid.UUID, tagIds []string, orderBy string, take *int32, skip *int32) (result training.ListImage, err error)
	GetTags(ctx context.Context, projectID uuid.UUID, iterationID *uuid.UUID) (result training.ListTag, err error)
	GetUntaggedImageCount(ctx context.Context, projectID uuid.UUID, iterationID *uuid.UUID) (result training.Int32, err error)
	GetUntaggedImages(ctx context.Context, projectID uuid.UUID, iterationID *uuid.UUID, orderBy string, take *int32, skip *int32) (result training.ListImage, err error)
	QueryPredictions(ctx context.Context, projectID uuid.UUID, query training.PredictionQueryToken) (result training.PredictionQueryResult, err error)
	QuickTestImage(ctx context.Context, projectID uuid.UUID, imageData io.ReadCloser, iterationID *uuid.UUID) (result training.ImagePrediction, err error)
	QuickTestImageURL(ctx context.Context, projectID uuid.UUID, imageURL training.ImageURL, iterationID *uuid.UUID) (result training.ImagePrediction, err error)
	TrainProject(ctx context.Context, projectID uuid.UUID) (result training.Iteration, err error)
	UpdateIteration(ctx context.Context, projectID uuid.UUID, iterationID uuid.UUID, updatedIteration training.Iteration) (result training.Iteration, err error)
	UpdateProject(ctx context.Context, projectID uuid.UUID, updatedProject training.Project) (result training.Project, err error)
	UpdateTag(ctx context.Context, projectID uuid.UUID, tagID uuid.UUID, updatedTag training.Tag) (result training.Tag, err error)
}

var _ BaseClientAPI = (*training.BaseClient)(nil)
