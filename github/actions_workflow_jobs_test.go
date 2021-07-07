// Copyright 2020 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestActionsService_ListWorkflowJobs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/runs/29679449/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"jobs":[{"id":399444496,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"},{"id":399444497,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}]}`)
	})

	opts := &ListWorkflowJobsOptions{ListOptions: ListOptions{Page: 2, PerPage: 2}}
	ctx := context.Background()
	jobs, _, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflowJobs returned error: %v", err)
	}

	want := &Jobs{
		TotalCount: Int(4),
		Jobs: []*WorkflowJob{
			{ID: Int64(399444496), RunID: Int64(29679449), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
			{ID: Int64(399444497), RunID: Int64(29679449), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
		},
	}
	if !cmp.Equal(jobs, want) {
		t.Errorf("Actions.ListWorkflowJobs returned %+v, want %+v", jobs, want)
	}

	const methodName = "ListWorkflowJobs"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.ListWorkflowJobs(ctx, "\n", "\n", 29679449, opts)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_ListWorkflowJobs_Filter(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/runs/29679449/jobs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"filter": "all", "per_page": "2", "page": "2"})
		fmt.Fprint(w, `{"total_count":4,"jobs":[{"id":399444496,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"},{"id":399444497,"run_id":29679449,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}]}`)
	})

	opts := &ListWorkflowJobsOptions{Filter: "all", ListOptions: ListOptions{Page: 2, PerPage: 2}}
	ctx := context.Background()
	jobs, _, err := client.Actions.ListWorkflowJobs(ctx, "o", "r", 29679449, opts)
	if err != nil {
		t.Errorf("Actions.ListWorkflowJobs returned error: %v", err)
	}

	want := &Jobs{
		TotalCount: Int(4),
		Jobs: []*WorkflowJob{
			{ID: Int64(399444496), RunID: Int64(29679449), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
			{ID: Int64(399444497), RunID: Int64(29679449), StartedAt: &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)}, CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)}},
		},
	}
	if !cmp.Equal(jobs, want) {
		t.Errorf("Actions.ListWorkflowJobs returned %+v, want %+v", jobs, want)
	}
}

func TestActionsService_GetWorkflowJobByID(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/jobs/399444496", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":399444496,"started_at":"2019-01-02T15:04:05Z","completed_at":"2020-01-02T15:04:05Z"}`)
	})

	ctx := context.Background()
	job, _, err := client.Actions.GetWorkflowJobByID(ctx, "o", "r", 399444496)
	if err != nil {
		t.Errorf("Actions.GetWorkflowJobByID returned error: %v", err)
	}

	want := &WorkflowJob{
		ID:          Int64(399444496),
		StartedAt:   &Timestamp{time.Date(2019, time.January, 02, 15, 04, 05, 0, time.UTC)},
		CompletedAt: &Timestamp{time.Date(2020, time.January, 02, 15, 04, 05, 0, time.UTC)},
	}
	if !cmp.Equal(job, want) {
		t.Errorf("Actions.GetWorkflowJobByID returned %+v, want %+v", job, want)
	}

	const methodName = "GetWorkflowJobByID"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.GetWorkflowJobByID(ctx, "\n", "\n", 399444496)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Actions.GetWorkflowJobByID(ctx, "o", "r", 399444496)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestActionsService_GetWorkflowJobLogs(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		http.Redirect(w, r, "http://github.com/a", http.StatusFound)
	})

	ctx := context.Background()
	url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, true)
	if err != nil {
		t.Errorf("Actions.GetWorkflowJobLogs returned error: %v", err)
	}
	if resp.StatusCode != http.StatusFound {
		t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusFound)
	}
	want := "http://github.com/a"
	if url.String() != want {
		t.Errorf("Actions.GetWorkflowJobLogs returned %+v, want %+v", url.String(), want)
	}

	const methodName = "GetWorkflowJobLogs"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.GetWorkflowJobLogs(ctx, "\n", "\n", 399444496, true)
		return err
	})

	// Add custom round tripper
	client.client.Transport = roundTripperFunc(func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("failed to get workflow logs")
	})
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, true)
		return err
	})
}

func TestActionsService_GetWorkflowJobLogs_StatusMovedPermanently_dontFollowRedirects(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		http.Redirect(w, r, "http://github.com/a", http.StatusMovedPermanently)
	})

	ctx := context.Background()
	_, resp, _ := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, false)
	if resp.StatusCode != http.StatusMovedPermanently {
		t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusMovedPermanently)
	}
}

func TestActionsService_GetWorkflowJobLogs_StatusMovedPermanently_followRedirects(t *testing.T) {
	client, mux, serverURL, teardown := setup()
	defer teardown()

	// Mock a redirect link, which leads to an archive link
	mux.HandleFunc("/repos/o/r/actions/jobs/399444496/logs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		redirectURL, _ := url.Parse(serverURL + baseURLPath + "/redirect")
		http.Redirect(w, r, redirectURL.String(), http.StatusMovedPermanently)
	})

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		http.Redirect(w, r, "http://github.com/a", http.StatusFound)
	})

	ctx := context.Background()
	url, resp, err := client.Actions.GetWorkflowJobLogs(ctx, "o", "r", 399444496, true)
	if err != nil {
		t.Errorf("Actions.GetWorkflowJobLogs returned error: %v", err)
	}

	if resp.StatusCode != http.StatusFound {
		t.Errorf("Actions.GetWorkflowJobLogs returned status: %d, want %d", resp.StatusCode, http.StatusFound)
	}

	want := "http://github.com/a"
	if url.String() != want {
		t.Errorf("Actions.GetWorkflowJobLogs returned %+v, want %+v", url.String(), want)
	}
}
