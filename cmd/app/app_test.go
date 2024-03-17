package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn/internal/config"

	"github.com/stretchr/testify/assert"
)

func runTestServer() *httptest.Server {
	config.InitialiseDatabase()

	return httptest.NewServer(setupServer())
}

func Test_post_api_integration_tests_store_endpoint(t *testing.T) {
	ts := runTestServer()
	defer ts.Close()

	t.Run("it should return 200 when health is ok", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/health", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode)
	})

}
