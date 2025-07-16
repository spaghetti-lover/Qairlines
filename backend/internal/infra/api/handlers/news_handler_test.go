package handlers_test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/entities"
	mocknews "github.com/spaghetti-lover/qairlines/internal/domain/mock/news"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/handlers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetNewsHandler(t *testing.T) {
	news := randomNews()

	testCases := []struct {
		name          string
		newsID        int64
		adminHeader   string
		buildStubs    func(mockUseCase *mocknews.MockIGetNewsUseCase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			newsID:      news.ID,
			adminHeader: "true",
			buildStubs: func(mockUseCase *mocknews.MockIGetNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), news.ID).
					Times(1).
					Return(news, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// Có thể kiểm tra thêm body nếu muốn
			},
		},
		{
			name:        "Unauthorized",
			newsID:      news.ID,
			adminHeader: "false",
			buildStubs: func(mockUseCase *mocknews.MockIGetNewsUseCase) {
				mockUseCase.EXPECT().Execute(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:        "NotFound",
			newsID:      news.ID,
			adminHeader: "true",
			buildStubs: func(mockUseCase *mocknews.MockIGetNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), news.ID).
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InvalidID",
			newsID:      0,
			adminHeader: "true",
			buildStubs: func(mockUseCase *mocknews.MockIGetNewsUseCase) {
				mockUseCase.EXPECT().Execute(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mocknews.NewMockIGetNewsUseCase(ctrl)
			tc.buildStubs(mockUseCase)

			handler := handlers.NewNewsHandler(nil, nil, nil, nil, mockUseCase, nil)
			router := gin.Default()
			router.GET("/api/news/:id", handler.GetNews)

			url := fmt.Sprintf("/api/news/%d", tc.newsID)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("admin", tc.adminHeader)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			tc.checkResponse(t, w)
		})
	}
}

func TestListNewsHandler(t *testing.T) {
	newsList := []entities.News{
		{
			ID:      1,
			Title:   "News 1",
			Content: "Content 1",
			Image:   "https://example.com/image1.jpg",
		},
		{
			ID:      2,
			Title:   "News 2",
			Content: "Content 2",
			Image:   "https://example.com/image2.jpg",
		},
	}

	testCases := []struct {
		name          string
		page          int
		limit         int
		buildStubs    func(mockUseCase *mocknews.MockIListNewsUseCase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			page:  1,
			limit: 10,
			buildStubs: func(mockUseCase *mocknews.MockIListNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), 1, 10).
					Times(1).
					Return(newsList, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:  "BindError",
			page:  0, // invalid page
			limit: 10,
			buildStubs: func(mockUseCase *mocknews.MockIListNewsUseCase) {
				mockUseCase.EXPECT().Execute(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:  "InternalError",
			page:  1,
			limit: 10,
			buildStubs: func(mockUseCase *mocknews.MockIListNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), 1, 10).
					Times(1).
					Return(nil, fmt.Errorf("unexpected error"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mocknews.NewMockIListNewsUseCase(ctrl)
			tc.buildStubs(mockUseCase)

			handler := handlers.NewNewsHandler(mockUseCase, nil, nil, nil, nil, nil)
			router := gin.Default()
			router.GET("/api/news", handler.ListNews)

			url := fmt.Sprintf("/api/news?page=%d&limit=%d", tc.page, tc.limit)
			req, _ := http.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			tc.checkResponse(t, w)
		})
	}
}

func TestDeleteNewsHandler(t *testing.T) {
	testCases := []struct {
		name          string
		newsID        int64
		adminHeader   string
		buildStubs    func(mockUseCase *mocknews.MockIDeleteNewsUseCase)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			newsID:      1,
			adminHeader: "true",
			buildStubs: func(mockUseCase *mocknews.MockIDeleteNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), int64(1)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:        "Unauthorized",
			newsID:      1,
			adminHeader: "false",
			buildStubs: func(mockUseCase *mocknews.MockIDeleteNewsUseCase) {
				mockUseCase.EXPECT().Execute(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:        "NotFound",
			adminHeader: "true",
			newsID:      2,
			buildStubs: func(mockUseCase *mocknews.MockIDeleteNewsUseCase) {
				mockUseCase.EXPECT().
					Execute(gomock.Any(), int64(2)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InvalidID",
			adminHeader: "true",
			newsID:      0,
			buildStubs: func(mockUseCase *mocknews.MockIDeleteNewsUseCase) {
				mockUseCase.EXPECT().Execute(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mocknews.NewMockIDeleteNewsUseCase(ctrl)
			tc.buildStubs(mockUseCase)

			handler := handlers.NewNewsHandler(nil, mockUseCase, nil, nil, nil, nil)
			router := gin.Default()
			router.DELETE("/api/news", handler.DeleteNews)

			url := fmt.Sprintf("/api/news?id=%d", tc.newsID)
			req, _ := http.NewRequest("DELETE", url, nil)
			req.Header.Set("admin", tc.adminHeader)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)
			tc.checkResponse(t, w)
		})
	}
}

func randomNews() *dto.GetNewsResponse {
	return &dto.GetNewsResponse{
		ID:      utils.RandomInt(1, 1000),
		Title:   utils.RandomString(10),
		Content: utils.RandomString(10),
		Image:   "https://example.com/image.jpg",
	}
}
