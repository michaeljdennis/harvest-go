package endpoint

// Users ...
type Users struct {
	*Endpoint
}

// NewUsers ...
func NewUsers() *Users {
	return &Users{
		Endpoint: NewEndpoint(
			"users",
		),
	}
}

// UsersModel ...
type UsersModel struct {
	BaseModel
	Users []UserModel `json:"users"`
}

// Me ...
type Me struct {
	*Endpoint
}

// NewMe returns new Me
func NewMe() *Me {
	return &Me{
		Endpoint: NewEndpoint(
			"users/me",
		),
	}
}

// UserModel ...
type UserModel struct {
	// BaseModel
	ID                           int      `json:"id"`
	FirstName                    string   `json:"first_name"`
	LastName                     string   `json:"last_name"`
	Email                        string   `json:"email"`
	Telephone                    string   `json:"telephone"`
	Timezone                     string   `json:"timezone"`
	HasAccessToAllFutureProjects bool     `json:"has_access_to_all_future_projects"`
	IsContractor                 bool     `json:"is_contractor"`
	IsAdmin                      bool     `json:"is_admin"`
	IsProjectManager             bool     `json:"is_project_manager"`
	CanSeeRates                  bool     `json:"can_see_rates"`
	CanCreateProjects            bool     `json:"can_create_projects"`
	CanCreateInvoices            bool     `json:"can_create_invoices"`
	IsActive                     bool     `json:"is_active"`
	CreatedAt                    string   `json:"created_at"`
	UpdatedAt                    string   `json:"updated_at"`
	WeeklyCapacity               int      `json:"weekly_capacity"`
	DefaultHourlyRate            float32  `json:"default_hourly_rate"`
	CostRate                     float32  `json:"cost_rate"`
	Roles                        []string `json:"roles"`
	AvatarURL                    string   `json:"avatar_url"`
}
