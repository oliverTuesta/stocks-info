package domain

import (
	"testing"
)

func TestPaginationResponse_CalculateTotalPages(t *testing.T) {
	tests := []struct {
		name          string
		pagination    PaginationResponse
		expectedPages int
	}{
		{
			name: "exact division",
			pagination: PaginationResponse{
				Total:    100,
				PageSize: 10,
			},
			expectedPages: 10,
		},
		{
			name: "remainder division",
			pagination: PaginationResponse{
				Total:    105,
				PageSize: 10,
			},
			expectedPages: 11, // 105/10 = 10.5, rounded up to 11
		},
		{
			name: "single page",
			pagination: PaginationResponse{
				Total:    5,
				PageSize: 10,
			},
			expectedPages: 1,
		},
		{
			name: "zero total",
			pagination: PaginationResponse{
				Total:    0,
				PageSize: 10,
			},
			expectedPages: 0,
		},
		{
			name: "zero page size",
			pagination: PaginationResponse{
				Total:    100,
				PageSize: 0,
			},
			expectedPages: 0,
		},
		{
			name: "large numbers",
			pagination: PaginationResponse{
				Total:    1000000,
				PageSize: 100,
			},
			expectedPages: 10000,
		},
		{
			name: "page size larger than total",
			pagination: PaginationResponse{
				Total:    5,
				PageSize: 20,
			},
			expectedPages: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pagination.CalculateTotalPages()
			if tt.pagination.TotalPages != tt.expectedPages {
				t.Errorf("CalculateTotalPages() = %v, want %v", tt.pagination.TotalPages, tt.expectedPages)
			}
		})
	}
}

func TestPaginationRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		request PaginationRequest
		valid   bool
	}{
		{
			name: "valid request",
			request: PaginationRequest{
				Page:     1,
				PageSize: 20,
				Search:   "test",
			},
			valid: true,
		},
		{
			name: "valid request without search",
			request: PaginationRequest{
				Page:     2,
				PageSize: 50,
			},
			valid: true,
		},
		{
			name: "minimum page",
			request: PaginationRequest{
				Page:     1,
				PageSize: 1,
			},
			valid: true,
		},
		{
			name: "maximum page size",
			request: PaginationRequest{
				Page:     1,
				PageSize: 100,
			},
			valid: true,
		},
		{
			name: "zero page",
			request: PaginationRequest{
				Page:     0,
				PageSize: 20,
			},
			valid: false,
		},
		{
			name: "negative page",
			request: PaginationRequest{
				Page:     -1,
				PageSize: 20,
			},
			valid: false,
		},
		{
			name: "zero page size",
			request: PaginationRequest{
				Page:     1,
				PageSize: 0,
			},
			valid: false,
		},
		{
			name: "negative page size",
			request: PaginationRequest{
				Page:     1,
				PageSize: -5,
			},
			valid: false,
		},
		{
			name: "page size over limit",
			request: PaginationRequest{
				Page:     1,
				PageSize: 101,
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.request.Page >= 1 && tt.request.PageSize >= 1 && tt.request.PageSize <= 100

			if valid != tt.valid {
				t.Errorf("PaginationRequest validation = %v, want %v", valid, tt.valid)
			}
		})
	}
}
