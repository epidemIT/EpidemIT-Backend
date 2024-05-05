package entity

type Mentor struct {
	FullName  string `json:"full_name"`
	Company   string `json:"company"`
	Specialty string `json:"specialty"`
	Bio       string `json:"bio"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
