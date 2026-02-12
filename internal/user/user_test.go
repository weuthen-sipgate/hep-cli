package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"hepic-cli/internal/api"
)

func TestList_ReturnsUsers(t *testing.T) {
	expected := map[string]interface{}{
		"count": float64(2),
		"data": []interface{}{
			map[string]interface{}{
				"username":   "admin",
				"email":      "admin@example.com",
				"firstname":  "Admin",
				"lastname":   "User",
				"usergroup":  "admin",
				"department": "IT",
			},
			map[string]interface{}{
				"username":   "user1",
				"email":      "user1@example.com",
				"firstname":  "Test",
				"lastname":   "User",
				"usergroup":  "user",
				"department": "Support",
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/users" {
			t.Errorf("expected path /users, got %s", r.URL.Path)
		}
		if r.Header.Get("Auth-Token") != "test-token" {
			t.Errorf("missing or incorrect Auth-Token header")
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := List(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	count, ok := parsed["count"].(float64)
	if !ok || count != 2 {
		t.Errorf("expected count 2, got %v", parsed["count"])
	}

	data, ok := parsed["data"].([]interface{})
	if !ok || len(data) != 2 {
		t.Errorf("expected 2 users, got %v", parsed["data"])
	}
}

func TestCreate_SendsCorrectPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/users" {
			t.Errorf("expected path /users, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)

		if req["username"] != "newuser" {
			t.Errorf("expected username 'newuser', got %v", req["username"])
		}
		if req["email"] != "new@example.com" {
			t.Errorf("expected email 'new@example.com', got %v", req["email"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "af7943f1-bdd4-4c52-8e53-1c3e834c0f58",
			"message": "successfully created user",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"username":   "newuser",
		"email":      "new@example.com",
		"password":   "secret123",
		"firstname":  "New",
		"lastname":   "User",
		"usergroup":  "user",
		"department": "Dev",
		"partid":     10,
	}

	result, err := Create(context.Background(), client, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully created user" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestUpdate_SendsCorrectPut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/users/test-uuid-123" {
			t.Errorf("expected path /users/test-uuid-123, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)

		if req["email"] != "updated@example.com" {
			t.Errorf("expected email 'updated@example.com', got %v", req["email"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "test-uuid-123",
			"message": "successfully updated user",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"email": "updated@example.com",
	}

	result, err := Update(context.Background(), client, "test-uuid-123", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully updated user" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestDelete_SendsCorrectDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/users/delete-uuid-456" {
			t.Errorf("expected path /users/delete-uuid-456, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "delete-uuid-456",
			"message": "successfully deleted user",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Delete(context.Background(), client, "delete-uuid-456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully deleted user" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestUpdatePassword_SendsCorrectPut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/users/update/password/pw-uuid" {
			t.Errorf("expected path /users/update/password/pw-uuid, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)

		if req["password"] != "newpass123" {
			t.Errorf("expected password 'newpass123', got %v", req["password"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "pw-uuid",
			"message": "successfully updated password",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"password": "newpass123",
	}

	result, err := UpdatePassword(context.Background(), client, "pw-uuid", data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully updated password" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestGroups_ReturnsGroupList(t *testing.T) {
	expected := map[string]interface{}{
		"count": float64(3),
		"data":  []interface{}{"admin", "user", "viewer"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/users/groups" {
			t.Errorf("expected path /users/groups, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Groups(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	data, ok := parsed["data"].([]interface{})
	if !ok || len(data) != 3 {
		t.Errorf("expected 3 groups, got %v", parsed["data"])
	}
}

func TestAuthTypes_ReturnsList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/auth/type/list" {
			t.Errorf("expected path /auth/type/list, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": []interface{}{"internal", "ldap"},
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := AuthTypes(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	data, ok := parsed["data"].([]interface{})
	if !ok || len(data) != 2 {
		t.Errorf("expected 2 auth types, got %v", parsed["data"])
	}
}

func TestExport_ReturnsRawBody(t *testing.T) {
	csvData := "username,email,firstname,lastname\nadmin,admin@example.com,Admin,User\n"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/users/export" {
			t.Errorf("expected path /users/export, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "text/csv")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(csvData))
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	body, err := Export(context.Background(), client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if string(data) != csvData {
		t.Errorf("expected CSV data, got %s", string(data))
	}
}

func TestImport_SendsMultipartPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/users/import" {
			t.Errorf("expected path /users/import, got %s", r.URL.Path)
		}

		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			t.Error("missing Content-Type header")
		}

		if r.Header.Get("Auth-Token") != "test-token" {
			t.Error("missing or incorrect Auth-Token header")
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			t.Fatalf("failed to parse multipart form: %v", err)
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			t.Fatalf("failed to get form file: %v", err)
		}
		defer file.Close()

		if header.Filename != "users.csv" {
			t.Errorf("expected filename 'users.csv', got %s", header.Filename)
		}

		content, _ := io.ReadAll(file)
		if string(content) != "username,email\ntest,test@example.com\n" {
			t.Errorf("unexpected file content: %s", string(content))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "ok",
			"message": "successfully imported users",
		})
	}))
	defer server.Close()

	// Create a temp CSV file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "users.csv")
	if err := os.WriteFile(tmpFile, []byte("username,email\ntest,test@example.com\n"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	client := api.NewClientWith(server.URL, "test-token")
	result, err := Import(context.Background(), client, tmpFile)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully imported users" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestImport_FileNotFound(t *testing.T) {
	client := api.NewClientWith("http://localhost:1", "test-token")
	_, err := Import(context.Background(), client, "/nonexistent/file.csv")
	if err == nil {
		t.Fatal("expected error for non-existent file")
	}
}

func TestCreateToken_SendsCorrectPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/token/auth" {
			t.Errorf("expected path /token/auth, got %s", r.URL.Path)
		}

		body, _ := io.ReadAll(r.Body)
		var req map[string]interface{}
		json.Unmarshal(body, &req)

		if req["name"] != "ci-token" {
			t.Errorf("expected name 'ci-token', got %v", req["name"])
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data":    "token-uuid-789",
			"message": "successfully created token",
			"token":   "abc123def456",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	data := map[string]interface{}{
		"name": "ci-token",
	}

	result, err := CreateToken(context.Background(), client, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully created token" {
		t.Errorf("expected success message, got %v", parsed["message"])
	}
}

func TestGetToken_ReturnsToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/token/auth/token-uuid-789" {
			t.Errorf("expected path /token/auth/token-uuid-789, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"guid":   "token-uuid-789",
			"name":   "ci-token",
			"active": true,
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := GetToken(context.Background(), client, "token-uuid-789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["name"] != "ci-token" {
		t.Errorf("expected name 'ci-token', got %v", parsed["name"])
	}
}

func TestDeleteToken_SendsCorrectDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/token/auth/token-uuid-789" {
			t.Errorf("expected path /token/auth/token-uuid-789, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"data":    "token-uuid-789",
			"message": "successfully deleted token",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	result, err := DeleteToken(context.Background(), client, "token-uuid-789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var parsed map[string]string
	if err := json.Unmarshal(result, &parsed); err != nil {
		t.Fatalf("failed to parse result: %v", err)
	}

	if parsed["message"] != "successfully deleted token" {
		t.Errorf("expected success message, got %s", parsed["message"])
	}
}

func TestList_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "unauthorized",
			"message": "invalid token",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "bad-token")
	_, err := List(context.Background(), client)
	if err == nil {
		t.Fatal("expected error for unauthorized request")
	}
}

func TestDelete_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "not found",
			"message": "user not found",
		})
	}))
	defer server.Close()

	client := api.NewClientWith(server.URL, "test-token")
	_, err := Delete(context.Background(), client, "nonexistent-uuid")
	if err == nil {
		t.Fatal("expected error for not found")
	}
}
