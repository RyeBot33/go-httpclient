package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T){
	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "testing")
	client.Headers = commonHeaders
	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "abc123")
	finalHeaders := client.getRequestHeaders(requestHeaders)
	// Validation
	if len(finalHeaders) != 3 {
		t.Error("Expected 3 headers, got ", len(finalHeaders))
	}
}