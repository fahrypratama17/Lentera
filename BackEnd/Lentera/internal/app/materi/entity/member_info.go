package entity

// MemberInfo = data anggota hasil join memberships + users.
// BUKAN tabel (tidak di-AutoMigrate), hanya wadah hasil query.
type MemberInfo struct {
	UserID uint
	Name   string
	Role   string // mentor | member
	Status string // active | left
}
