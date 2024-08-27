package snapshot

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/opensearch-project/opensearch-go/v2"

	"github.com/dim-ops/opensearch-snapshot/internal/config"
)

type MockTransport struct {
	Response *http.Response
	Err      error
}

func (m *MockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return m.Response, m.Err
}

func TestCreateRepository(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		responseBody     string
		expectedError    string
		opensearchConfig config.OpenSearchConfig
	}{
		{
			name:          "Successful repository creation",
			statusCode:    200,
			responseBody:  `{"acknowledged": true}`,
			expectedError: "",
			opensearchConfig: config.OpenSearchConfig{
				Bucket:   "test-bucket",
				Region:   "us-west-2",
				RoleARN:  "arn:aws:iam::123456789012:role/OpenSearchRole",
				Clusters: []string{"https://test-cluster.us-west-2.es.amazonaws.com"},
			},
		},
		{
			name:          "Failed repository creation",
			statusCode:    400,
			responseBody:  `{"error": "Bad Request"}`,
			expectedError: "HTTP - 400",
			opensearchConfig: config.OpenSearchConfig{
				Bucket:   "test-bucket",
				Region:   "us-west-2",
				RoleARN:  "arn:aws:iam::123456789012:role/OpenSearchRole",
				Clusters: []string{"https://test-cluster.us-west-2.es.amazonaws.com"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := &config.Config{
				Opensearch: tc.opensearchConfig,
			}

			mockTransport := &MockTransport{
				Response: &http.Response{
					StatusCode: tc.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tc.responseBody)),
				},
			}

			client, _ := opensearch.NewClient(opensearch.Config{
				Transport: mockTransport,
			})

			err := CreateRepository(0, client, cfg)

			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Error("Expected an error, got nil")
				} else if err.Error() != tc.expectedError {
					t.Errorf("Expected error '%s', got '%v'", tc.expectedError, err)
				}
			}
		})
	}
}

func TestCreateSnapshot(t *testing.T) {
	testCases := []struct {
		name          string
		statusCode    int
		responseBody  string
		expectedError string
	}{
		{
			name:          "Successful snapshot creation",
			statusCode:    200,
			responseBody:  `{"acknowledged": true}`,
			expectedError: "",
		},
		{
			name:          "Failed snapshot creation",
			statusCode:    400,
			responseBody:  `{"error": "Bad Request"}`,
			expectedError: "HTTP - 400",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockTransport := &MockTransport{
				Response: &http.Response{
					StatusCode: tc.statusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tc.responseBody)),
				},
			}

			client, _ := opensearch.NewClient(opensearch.Config{
				Transport: mockTransport,
			})

			err := CreateSnapshot(client)

			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Error("Expected an error, got nil")
				} else if err.Error() != tc.expectedError {
					t.Errorf("Expected error '%s', got '%v'", tc.expectedError, err)
				}
			}
		})
	}
}
