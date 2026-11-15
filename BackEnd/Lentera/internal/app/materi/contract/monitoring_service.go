package contract

import (
	"context"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
)

type MonitoringService interface {
	Dashboard(ctx context.Context, userID, cohortID uint) (*dto.DashboardResponse, error)
	Matrix(ctx context.Context, userID, cohortID uint, q dto.MatrixQuery) (*dto.MatrixResponse, error)
	MemberProgress(ctx context.Context, userID, cohortID, memberID uint) (*dto.MemberProgressResponse, error)
}
