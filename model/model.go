package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Company
// Rol ->
const (
	TimeLayout string = "2006-01-02T15:04:05"

	RolOwner string = "OWNER"
	RolStaff string = "STAFF"

	SourceAutomatic string = "AUTOMATIC"
	SourceManual    string = "MANUAL"

	StatusPending   string = "PENDING_CONFIRMATION"
	StatusConfirmed string = "CONFIRMED"
	StatusActive    string = "ACTIVE"
	StatusInactive  string = "INACTIVE"

	DocTypeStaff   string = "Staff"
	DocTypeService string = "Service"
	DocTypeCompany string = "Company"
	DocTypeClass   string = "Class"

	SexM string = "M"
	SexW string = "W"
	SexA string = "A" //Any
)

type Company struct {
	ID          string            `json:"CompanyID"`
	UserSub     string            `json:"UserSub"`
	Name        string            `json:"Name,omitempty"`
	Description string            `json:"Description,omitempty"`
	Timestamp   string            `json:"Timestamp,omitempty"`
	Staff       []*Staff          `json:"Staff,omitempty"`
	Services    []*CompanyService `json:"Services,omitempty"`
	Rol         string            `json:"Rol,omitempty"`
	Status      string            `json:"Status,omitempty"` // ACTIVE, INACTIVE
}

type Staff struct {
	ID        string `json:"StaffID"`
	UserSub   string `json:"UserSub"`
	CompanyID string `json:"CompanyID,omitempty"`
	Name      string `json:"Name,omitempty"`
	Rol       string `json:"Rol,omitempty"`
	Status    string `json:"Status,omitempty"` // PENDING_CONFIRMATION, CONFIRMED
	CreatedBy string `json:"CreatedBy,omitempty"`
}

type CompanyService struct {
	ID                 string              `json:"ServiceID"`
	UserSub            string              `json:"UserSub"`
	CompanyID          string              `json:"CompanyID,omitempty"`
	Name               string              `json:"Name,omitempty"`
	Description        string              `json:"Description,omitempty"`
	ServiceType        string              `json:"ServiceType,omitempty"`
	Duration           int64               `json:"Duration,omitempty"`
	Price              int64               `json:"Price,omitempty"`
	CurrencyID         string              `json:"CurrencyID,omitempty"`
	Status             string              `json:"Status,omitempty"` // ACTIVE, INACTIVE
	Sex                string              `json:"Sex"`              // Men, Women, Any
	MinAge             int                 `json:"MinAge"`
	MaxAge             int                 `json:"MaxAge"`
	Tags               []string            `json:"Tags,omitempty"` // For instance: HARD, SOFT, etc.
	CancellationPolicy *CancellationPolicy `json:"CancellationPolicy,omitempty"`
}

type AutomaticSchedule struct {
	Periodicity string `json:"Periodicity"` // Dailie, Weekly, Monthly, Yearly
	Monday      *bool  `json:"Monday,omitempty"`
	Tuesday     *bool  `json:"Tuesday,omitempty"`
	Wednesday   *bool  `json:"Wednesday,omitempty"`
	Thursday    *bool  `json:"Thursday,omitempty"`
	Friday      *bool  `json:"Friday,omitempty"`
	Saturday    *bool  `json:"Saturday,omitempty"`
	Sunday      *bool  `json:"Sunday,omitempty"`
	StartTime   string `json:"StartTime,omitempty"` // Inherited from service if not specified
	Duration    int64  `json:"Duration,omitempty"`  // Minutes
	ClassID     string `json:"ClassID"`
}

type Class struct {
	ID                 string              `json:"ClassID"`
	ServiceID          string              `json:"ServiceID"`
	UserSub            string              `json:"UserSub"`
	CompanyID          string              `json:"CompanyID,omitempty"` // Repeating attribute for direct access
	MinBookings        int                 `json:"MinBookings,omitempty"`
	MaxBookings        int                 `json:"MaxBookings,omitempty"`
	StartTime          string              `json:"StartTime,omitempty"` // Inherited from service if not specified
	Duration           time.Duration       `json:"Duration,omitempty"`  // Minutes
	Place              *Place              `json:"Place,omitempty"`
	Price              int64               `json:"Price,omitempty"`              // Inherited from service if not specified
	CancellationPolicy *CancellationPolicy `json:"CancellationPolicy,omitempty"` // Inherited from service if not specified
	Status             string              `json:"Status,omitempty"`             // CREATED -> PENDING_CONFIRMATION ->
	Source             string              `json:"Source,omitempty"`             // AUTOMATIC, MANUAL
}

func NewClass(userSub string) Class {
	c := Class{}
	c.UserSub = userSub
	c.MinBookings = 0
	c.MaxBookings = 99
	c.Price = 0
	c.Status = StatusPending
	c.Source = SourceManual
	c.Duration = time.Minute * 60
	c.StartTime = time.Now().Add(time.Minute * 60).Format(TimeLayout)
	return c
}

type CancellationPolicy struct {
	ID             string  `json:"CancellationPolicyID"`
	Name           string  `json:"Name"`
	CostPercentage float64 `json:"CostPercentage"`      // 0% means free_cancellation
	TimeLimit      int64   `json:"TimeLimit,omitempty"` // Minutes before e.g: 60 min before start
}

type Place struct {
	Name      string  `json:"Name"`
	Direction string  `json:"Direction"`
	Latitud   float64 `json:"Latitud"`
	Longitud  float64 `json:"Longitud"`
}

func NewCompanyService(companyID string, userSub string) CompanyService {
	c := CompanyService{}
	uid, _ := uuid.NewV4()
	c.ID = uid.String()
	c.MinAge = -1
	c.MaxAge = 9999
	c.CompanyID = companyID
	c.Status = StatusActive
	c.Sex = SexA
	c.UserSub = userSub
	return c
}

func NewCompany(userSub string) Company {
	c := Company{}
	uid, _ := uuid.NewV4()
	c.ID = uid.String()
	c.UserSub = userSub
	c.Rol = RolOwner
	c.Status = StatusActive
	c.Timestamp = time.Now().Format(TimeLayout)
	return c
}
