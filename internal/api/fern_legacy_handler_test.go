	r := gin.Default()
	r := gin.Default()
	r := gin.Default()
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	projectsApp "github.com/guidewire-oss/fern-platform/internal/domains/projects/application"
	projectsDomain "github.com/guidewire-oss/fern-platform/internal/domains/projects/domain"
	// ...existing code...
	"github.com/guidewire-oss/fern-platform/internal/domains/testing/domain"
	"github.com/guidewire-oss/fern-platform/pkg/logging"
	"github.com/guidewire-oss/fern-platform/pkg/config"
)

// Mocks

type mockTestRunService struct {
	mock.Mock
}

func (m *mockTestRunService) CreateTestRun(ctx interface{}, tr *domain.TestRun) error {
	args := m.Called(ctx, tr)
	return args.Error(0)
}
func (m *mockTestRunService) GetTestRunByRunID(ctx interface{}, runID string) (*domain.TestRun, error) {
	args := m.Called(ctx, runID)
	return args.Get(0).(*domain.TestRun), args.Error(1)
}
func (m *mockTestRunService) ListTestRuns(ctx interface{}, projectUUID string, limit, offset int) ([]*domain.TestRun, int, error) {
	args := m.Called(ctx, projectUUID, limit, offset)
	return args.Get(0).([]*domain.TestRun), args.Int(1), args.Error(2)
}
func (m *mockTestRunService) CreateSuiteRun(ctx interface{}, sr *domain.SuiteRun) error { return nil }
func (m *mockTestRunService) CreateSpecRun(ctx interface{}, sr *domain.SpecRun) error { return nil }

// ...other methods as needed...

type mockProjectService struct {
	mock.Mock
}
func (m *mockProjectService) CreateProject(ctx interface{}, id projectsDomain.ProjectID, name string, team projectsDomain.Team, source string) (*projectsDomain.Project, error) {
	args := m.Called(ctx, id, name, team, source)
	return args.Get(0).(*projectsDomain.Project), args.Error(1)
}
func (m *mockProjectService) GetProject(ctx interface{}, id projectsDomain.ProjectID) (*projectsDomain.Project, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*projectsDomain.Project), args.Error(1)
}
func (m *mockProjectService) UpdateProject(ctx interface{}, id projectsDomain.ProjectID, req projectsApp.UpdateProjectRequest) error { return nil }
func (m *mockProjectService) ListProjects(ctx interface{}, limit, offset int) ([]*projectsDomain.Project, int, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*projectsDomain.Project), args.Int(1), args.Error(2)
}

func TestCreateFernTestReport(t *testing.T) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	mockTestRunSvc := new(mockTestRunService)
	mockProjectSvc := new(mockProjectService)
	cfg := &config.LoggingConfig{Level: "debug", Format: "text", Output: "stdout", Structured: false}
	logger, _ := logging.NewLogger(cfg)
	handler := NewFernLegacyHandler(mockTestRunSvc, mockProjectSvc, logger)
	 handler.RegisterRoutes(r.Group("/api"))

	project := &projectsDomain.Project{}
	mockProjectSvc.On("GetProject", mock.Anything, projectsDomain.ProjectID("proj-123")).Return(project, nil)
	mockTestRunSvc.On("GetTestRunByRunID", mock.Anything, mock.Anything).Return((*domain.TestRun)(nil), nil)
	mockTestRunSvc.On("CreateTestRun", mock.Anything, mock.Anything).Return(nil)

	body := map[string]interface{}{
		"test_project_id": "proj-123",
		"test_project_name": "TestProj",
		"test_seed": 42,
		"start_time": time.Now().Format(time.RFC3339),
		"suite_runs": []interface{}{},
	}
	b, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/reports/testrun", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestListFernTestReports(t *testing.T) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	 r := gin.Default()
	mockTestRunSvc := new(mockTestRunService)
	mockProjectSvc := new(mockProjectService)
	cfg := &config.LoggingConfig{Level: "debug", Format: "text", Output: "stdout", Structured: false}
	logger, _ := logging.NewLogger(cfg)
	handler := NewFernLegacyHandler(mockTestRunSvc, mockProjectSvc, logger)
	 handler.RegisterRoutes(r.Group("/api"))
	 handler.RegisterRoutes(r.Group("/api"))

	tr := &domain.TestRun{RunID: "run-1", ProjectID: "proj-1", GitBranch: "main", GitCommit: "abc123", Status: "completed", StartTime: time.Now()}
	mockTestRunSvc.On("ListTestRuns", mock.Anything, "proj-1", 20, 0).Return([]*domain.TestRun{tr}, 1, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/reports/testruns?project_uuid=proj-1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetFernTestReport_NotFound(t *testing.T) {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	 r := gin.Default()
	mockTestRunSvc := new(mockTestRunService)
	mockProjectSvc := new(mockProjectService)
	cfg := &config.LoggingConfig{Level: "debug", Format: "text", Output: "stdout", Structured: false}
	logger, _ := logging.NewLogger(cfg)
	handler := NewFernLegacyHandler(mockTestRunSvc, mockProjectSvc, logger)
	 handler.RegisterRoutes(r.Group("/api"))
	r := gin.Default()
	handler.RegisterRoutes(r.Group("/api"))

	mockTestRunSvc.On("GetTestRunByRunID", mock.Anything, "run-404").Return((*domain.TestRun)(nil), assert.AnError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/reports/testrun/run-404", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
