package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/entity"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.CohortRepository {
	return &repository{db: db}
}

func (r *repository) CreateWithOwner(ctx context.Context, cohort *entity.Cohort, owner *entity.Membership) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(cohort).Error; err != nil {
			return err
		}
		owner.CohortID = cohort.ID
		return tx.Create(owner).Error
	})
}

func (r *repository) FindUserCohorts(ctx context.Context, userID uint) ([]dto.CohortResponse, error) {
	type row struct {
		ID          uint
		Name        string
		Description string
		StartDate   time.Time
		EndDate     time.Time
		Role        entity.CohortRole
		MemberCount int64
		JoinedAt    time.Time
		CreatedAt   time.Time
	}

	var rows []row
	err := r.db.WithContext(ctx).
		Table("memberships AS m").
		Select(`
			c.id,
			c.name,
			c.description,
			c.start_date,
			c.end_date,
			m.role,
			m.joined_at,
			c.created_at,
			(SELECT COUNT(*) FROM memberships WHERE cohort_id = c.id AND status = ?) AS member_count
		`, entity.MembershipActive).
		Joins("JOIN cohorts AS c ON c.id = m.cohort_id").
		Where("m.user_id = ? AND m.status = ?", userID, entity.MembershipActive).
		Order("m.joined_at DESC").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}

	results := make([]dto.CohortResponse, 0, len(rows))
	for _, item := range rows {
		results = append(results, dto.CohortResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			StartDate:   item.StartDate.Format("2006-01-02"),
			EndDate:     item.EndDate.Format("2006-01-02"),
			Role:        item.Role,
			MemberCount: item.MemberCount,
			JoinedAt:    item.JoinedAt,
			CreatedAt:   item.CreatedAt,
		})
	}
	return results, nil
}

func (r *repository) FindCohortByID(ctx context.Context, cohortID uint) (*entity.Cohort, error) {
	var c entity.Cohort
	if err := r.db.WithContext(ctx).First(&c, cohortID).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *repository) FindMembership(ctx context.Context, cohortID, userID uint) (*entity.Membership, error) {
	var m entity.Membership
	err := r.db.WithContext(ctx).
		Where("cohort_id = ? AND user_id = ?", cohortID, userID).
		First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *repository) FindUserByID(ctx context.Context, userID uint) (*contract.UserBasic, error) {
	var u contract.UserBasic
	err := r.db.WithContext(ctx).Table("users").Select("id, name, email").Where("id = ?", userID).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*contract.UserBasic, error) {
	var u contract.UserBasic
	err := r.db.WithContext(ctx).Table("users").Select("id, name, email").Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repository) CreateInvitation(ctx context.Context, inv *entity.Invitation) error {
	return r.db.WithContext(ctx).Create(inv).Error
}

func (r *repository) FindPendingInvitation(ctx context.Context, cohortID uint, email string) (*entity.Invitation, error) {
	var inv entity.Invitation
	err := r.db.WithContext(ctx).
		Where("cohort_id = ? AND email = ? AND status = ?", cohortID, email, entity.InvitationPending).
		First(&inv).Error
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

func (r *repository) FindInvitationByID(ctx context.Context, invitationID uint) (*entity.Invitation, error) {
	var inv entity.Invitation
	if err := r.db.WithContext(ctx).First(&inv, invitationID).Error; err != nil {
		return nil, err
	}
	return &inv, nil
}

func (r *repository) FindUserPendingInvitations(ctx context.Context, email string) ([]dto.InvitationResponse, error) {
	var res []dto.InvitationResponse
	err := r.db.WithContext(ctx).
		Table("invitations AS i").
		Select(`
			i.id,
			i.cohort_id,
			c.name AS cohort_name,
			i.inviter_id,
			u.name AS inviter_name,
			i.email,
			i.role,
			i.status,
			i.created_at,
			i.responded_at
		`).
		Joins("JOIN cohorts AS c ON c.id = i.cohort_id").
		Joins("JOIN users AS u ON u.id = i.inviter_id").
		Where("i.email = ? AND i.status = ?", email, entity.InvitationPending).
		Order("i.created_at DESC").
		Scan(&res).Error
	return res, err
}

func (r *repository) FindCohortInvitations(ctx context.Context, cohortID uint) ([]dto.InvitationResponse, error) {
	var res []dto.InvitationResponse
	err := r.db.WithContext(ctx).
		Table("invitations AS i").
		Select(`
			i.id,
			i.cohort_id,
			c.name AS cohort_name,
			i.inviter_id,
			u.name AS inviter_name,
			i.email,
			i.role,
			i.status,
			i.created_at,
			i.responded_at
		`).
		Joins("JOIN cohorts AS c ON c.id = i.cohort_id").
		Joins("JOIN users AS u ON u.id = i.inviter_id").
		Where("i.cohort_id = ?", cohortID).
		Order("i.created_at DESC").
		Scan(&res).Error
	return res, err
}

func (r *repository) UpdateInvitationStatus(ctx context.Context, invitationID uint, status entity.InvitationStatus, respondedAt *time.Time) error {
	updates := map[string]any{"status": status}
	if respondedAt != nil {
		updates["responded_at"] = respondedAt
	}
	return r.db.WithContext(ctx).Model(&entity.Invitation{}).Where("id = ?", invitationID).Updates(updates).Error
}

func (r *repository) AcceptInvitationTx(ctx context.Context, inv *entity.Invitation, userID uint, now time.Time) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Invitation{}).
			Where("id = ?", inv.ID).
			Updates(map[string]any{
				"status":       entity.InvitationAccepted,
				"responded_at": now,
			}).Error; err != nil {
			return err
		}

		var existing entity.Membership
		err := tx.Where("cohort_id = ? AND user_id = ?", inv.CohortID, userID).First(&existing).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx.Create(&entity.Membership{
				CohortID: inv.CohortID,
				UserID:   userID,
				Role:     inv.Role,
				Status:   entity.MembershipActive,
				JoinedAt: now,
			}).Error
		}
		if err != nil {
			return err
		}

		return tx.Model(&entity.Membership{}).
			Where("id = ?", existing.ID).
			Updates(map[string]any{
				"role":      inv.Role,
				"status":    entity.MembershipActive,
				"joined_at": now,
				"left_at":   nil,
			}).Error
	})
}

func (r *repository) UpdateCohort(ctx context.Context, cohort *entity.Cohort) error {
	return r.db.WithContext(ctx).Save(cohort).Error
}

func (r *repository) DeleteCohort(ctx context.Context, cohortID uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Cohort{}, cohortID).Error
}

func (r *repository) ListActiveMembers(ctx context.Context, cohortID uint) ([]dto.MemberResponse, error) {
	var members []dto.MemberResponse
	err := r.db.WithContext(ctx).
		Table("memberships AS m").
		Select(`
			m.user_id,
			u.name,
			u.email,
			m.role,
			m.joined_at
		`).
		Joins("JOIN users AS u ON u.id = m.user_id").
		Where("m.cohort_id = ? AND m.status = ?", cohortID, entity.MembershipActive).
		Order("CASE WHEN m.role = 'mentor' THEN 0 ELSE 1 END, m.joined_at ASC").
		Scan(&members).Error
	return members, err
}

func (r *repository) CountActiveMentors(ctx context.Context, cohortID uint) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.Membership{}).
		Where("cohort_id = ? AND role = ? AND status = ?", cohortID, entity.RoleMentor, entity.MembershipActive).
		Count(&count).Error
	return count, err
}

func (r *repository) UpdateMemberRole(ctx context.Context, cohortID, userID uint, role entity.CohortRole) error {
	return r.db.WithContext(ctx).
		Model(&entity.Membership{}).
		Where("cohort_id = ? AND user_id = ? AND status = ?", cohortID, userID, entity.MembershipActive).
		Update("role", role).Error
}

func (r *repository) LeaveMembership(ctx context.Context, cohortID, userID uint, leftAt time.Time) error {
	return r.db.WithContext(ctx).
		Model(&entity.Membership{}).
		Where("cohort_id = ? AND user_id = ? AND status = ?", cohortID, userID, entity.MembershipActive).
		Updates(map[string]any{
			"status":  entity.MembershipLeft,
			"left_at": leftAt,
		}).Error
}