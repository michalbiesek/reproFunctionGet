package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
)

// Config holds the SDK client and server configuration
type Config struct {
	Client    *criblcontrolplanesdkgo.CriblControlPlane
	ServerURL string
}

// loginResponse represents the response from the login endpoint
type loginResponse struct {
	Token string `json:"token"`
}

// login authenticates with the Cribl server and returns a bearer token
func login(baseURL, username, password string) (string, error) {
	loginURL := baseURL + "/api/v1/auth/login"

	loginData := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(loginData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login data: %w", err)
	}

	req, err := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute login request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body to check what we got
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var loginResp loginResponse
	if err := json.Unmarshal(bodyBytes, &loginResp); err != nil {
		return "", fmt.Errorf("failed to decode login response: %w, response body: %s", err, string(bodyBytes))
	}

	if loginResp.Token == "" {
		return "", fmt.Errorf("login response did not contain a token, response: %s", string(bodyBytes))
	}

	return loginResp.Token, nil
}

// newClient creates a new SDK client with configuration from environment variables
func newClient() (*Config, error) {
	// Get base server URL from environment or use default
	baseURL := os.Getenv("CRIBL_SERVER_URL")
	if baseURL == "" {
		baseURL = "http://localhost:9000"
	}

	// Remove trailing slash and ensure we have the base URL without /api/v1
	baseURL = strings.TrimSuffix(baseURL, "/")
	baseURL = strings.TrimSuffix(baseURL, "/api/v1")

	// Construct server URL with /api/v1
	serverURL := baseURL + "/api/v1"

	// Get username and password from environment, default to "admin"
	username := os.Getenv("CRIBL_USERNAME")
	if username == "" {
		username = os.Getenv("CRIBL_USER")
	}
	if username == "" {
		username = "admin"
	}

	password := os.Getenv("CRIBL_PASSWORD")
	if password == "" {
		password = os.Getenv("CRIBL_PASS")
	}
	if password == "" {
		password = "admin"
	}

	// Login to get auth token
	fmt.Println("🔐 Authenticating with Cribl server...")
	authToken, err := login(baseURL, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}
	fmt.Println("✅ Authentication successful!")

	// Create SDK client with security
	security := components.Security{
		BearerAuth: &authToken,
	}

	sdk := criblcontrolplanesdkgo.New(
		serverURL,
		criblcontrolplanesdkgo.WithSecurity(security),
	)

	return &Config{
		Client:    sdk,
		ServerURL: serverURL,
	}, nil
}

func main() {
	fmt.Println("\n=== Go SDK Function Get Example ===")
	fmt.Println("====================================")

	// Step 1: Create authenticated SDK client
	// This will login to the Cribl server and create a client with bearer token authentication
	config, err := newClient()
	if err != nil {
		log.Fatalf("❌ Failed to create client: %v", err)
	}

	fmt.Printf("🌐 Connected to: %s\n", config.ServerURL)
	fmt.Println()

	ctx := context.Background()

	// Step 2: Get function by ID
	// This demonstrates the Functions.Get() method which should return the function
	// with its schema populated (in beta version) or empty (in RC version)
	fmt.Println("--- Get Function ---")
	fmt.Println("Fetching function: aggregate_metrics")
	fmt.Println()

	// Note: Don't use WithServerURL here as it may interfere with URL construction
	// The SDK client was already initialized with the correct serverURL (baseURL + "/api/v1")
	functionResponse, err := config.Client.Functions.Get(ctx, "aggregate_metrics")
	if err != nil {
		log.Fatalf("❌ Failed to get function: %v", err)
	}

	// Step 3: Inspect the response
	// The response should contain a CountedFunctionResponse with items
	// Each item should have a schema field that is populated in beta but empty in RC
	if functionResponse.CountedFunctionResponse != nil {
		responseJSON, err := json.MarshalIndent(functionResponse.CountedFunctionResponse, "", "  ")
		if err != nil {
			log.Fatalf("❌ Failed to marshal response: %v", err)
		}
		fmt.Println("Response:")
		fmt.Println(string(responseJSON))
		fmt.Println()
		fmt.Println("Note: Check the 'schema' field in the response.")
		fmt.Println("  - Beta version (v0.5.0-beta.21): schema contains properties")
		fmt.Println("  - RC version (v0.5.0-rc.23): schema is empty {}")
	} else {
		fmt.Println("⚠️ No function response received")
	}

	fmt.Println()
	fmt.Println("🎉 Example completed!")
}
