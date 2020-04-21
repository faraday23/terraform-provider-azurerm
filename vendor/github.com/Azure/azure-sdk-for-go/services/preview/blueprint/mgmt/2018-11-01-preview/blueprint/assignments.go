package blueprint

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssignmentsClient is the blueprint Client
type AssignmentsClient struct {
	BaseClient
}

// NewAssignmentsClient creates an instance of the AssignmentsClient client.
func NewAssignmentsClient() AssignmentsClient {
	return NewAssignmentsClientWithBaseURI(DefaultBaseURI)
}

// NewAssignmentsClientWithBaseURI creates an instance of the AssignmentsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAssignmentsClientWithBaseURI(baseURI string) AssignmentsClient {
	return AssignmentsClient{NewWithBaseURI(baseURI)}
}

// CreateOrUpdate create or update a blueprint assignment.
// Parameters:
// resourceScope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the blueprint assignment.
// assignment - blueprint assignment object to save.
func (client AssignmentsClient) CreateOrUpdate(ctx context.Context, resourceScope string, assignmentName string, assignment Assignment) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: assignment,
			Constraints: []validation.Constraint{{Target: "assignment.Identity", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "assignment.AssignmentProperties", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "assignment.AssignmentProperties.Parameters", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "assignment.AssignmentProperties.ResourceGroups", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("blueprint.AssignmentsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceScope, assignmentName, assignment)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AssignmentsClient) CreateOrUpdatePreparer(ctx context.Context, resourceScope string, assignmentName string, assignment Assignment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName": autorest.Encode("path", assignmentName),
		"resourceScope":  resourceScope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", pathParameters),
		autorest.WithJSON(assignment),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) CreateOrUpdateResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a blueprint assignment.
// Parameters:
// resourceScope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the blueprint assignment.
func (client AssignmentsClient) Delete(ctx context.Context, resourceScope string, assignmentName string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceScope, assignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AssignmentsClient) DeletePreparer(ctx context.Context, resourceScope string, assignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName": autorest.Encode("path", assignmentName),
		"resourceScope":  resourceScope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) DeleteResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get a blueprint assignment.
// Parameters:
// resourceScope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the blueprint assignment.
func (client AssignmentsClient) Get(ctx context.Context, resourceScope string, assignmentName string) (result Assignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceScope, assignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssignmentsClient) GetPreparer(ctx context.Context, resourceScope string, assignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName": autorest.Encode("path", assignmentName),
		"resourceScope":  resourceScope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) GetResponder(resp *http.Response) (result Assignment, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list blueprint assignments within a subscription or a management group.
// Parameters:
// resourceScope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
func (client AssignmentsClient) List(ctx context.Context, resourceScope string) (result AssignmentListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.List")
		defer func() {
			sc := -1
			if result.al.Response.Response != nil {
				sc = result.al.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceScope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.al.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.al, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AssignmentsClient) ListPreparer(ctx context.Context, resourceScope string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceScope": resourceScope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) ListResponder(resp *http.Response) (result AssignmentList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AssignmentsClient) listNextResults(ctx context.Context, lastResults AssignmentList) (result AssignmentList, err error) {
	req, err := lastResults.assignmentListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AssignmentsClient) ListComplete(ctx context.Context, resourceScope string) (result AssignmentListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceScope)
	return
}

// WhoIsBlueprint get Blueprints service SPN objectId
// Parameters:
// resourceScope - the scope of the resource. Valid scopes are: management group (format:
// '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format:
// '/subscriptions/{subscriptionId}').
// assignmentName - name of the blueprint assignment.
func (client AssignmentsClient) WhoIsBlueprint(ctx context.Context, resourceScope string, assignmentName string) (result WhoIsBlueprintContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssignmentsClient.WhoIsBlueprint")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.WhoIsBlueprintPreparer(ctx, resourceScope, assignmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "WhoIsBlueprint", nil, "Failure preparing request")
		return
	}

	resp, err := client.WhoIsBlueprintSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "WhoIsBlueprint", resp, "Failure sending request")
		return
	}

	result, err = client.WhoIsBlueprintResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blueprint.AssignmentsClient", "WhoIsBlueprint", resp, "Failure responding to request")
	}

	return
}

// WhoIsBlueprintPreparer prepares the WhoIsBlueprint request.
func (client AssignmentsClient) WhoIsBlueprintPreparer(ctx context.Context, resourceScope string, assignmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"assignmentName": autorest.Encode("path", assignmentName),
		"resourceScope":  resourceScope,
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceScope}/providers/Microsoft.Blueprint/blueprintAssignments/{assignmentName}/WhoIsBlueprint", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// WhoIsBlueprintSender sends the WhoIsBlueprint request. The method will close the
// http.Response Body if it receives an error.
func (client AssignmentsClient) WhoIsBlueprintSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// WhoIsBlueprintResponder handles the response to the WhoIsBlueprint request. The method always
// closes the http.Response Body.
func (client AssignmentsClient) WhoIsBlueprintResponder(resp *http.Response) (result WhoIsBlueprintContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
