package models

// Enumeraciones como tipos personalizados
type Role string

const (
	RoleAndrei       Role = "ANDREI"
	RoleDaemon       Role = "DAEMON"
	RoleNetworkAdmin Role = "NETWORK_ADMIN"
)

type VictimStatus string

const (
	StatusCaptured  VictimStatus = "CAPTURED"
	StatusEscaped   VictimStatus = "ESCAPED"
	StatusConverted VictimStatus = "CONVERTED"
)

type ReportType string

const (
	ReportResistanceAttempt ReportType = "RESISTANCE_ATTEMPT"
	ReportEscapeAttempt     ReportType = "ESCAPE_ATTEMPT"
	ReportSabotage          ReportType = "SABOTAGE"
)

type Severity string

const (
	SeverityLow      Severity = "LOW"
	SeverityMedium   Severity = "MEDIUM"
	SeverityHigh     Severity = "HIGH"
	SeverityCritical Severity = "CRITICAL"
)

type ReportStatus string

const (
	StatusPending  ReportStatus = "PENDING"
	StatusReviewed ReportStatus = "REVIEWED"
	StatusResolved ReportStatus = "RESOLVED"
)

type RewardType string

const (
	TypeReward     RewardType = "REWARD"
	TypePunishment RewardType = "PUNISHMENT"
)

type ReportCategory string

const (
	CategoryGeneral        ReportCategory = "GENERAL"
	CategorySurvivalTip    ReportCategory = "SURVIVAL_TIP"
	CategoryWarning        ReportCategory = "WARNING"
	CategoryResistancePlan ReportCategory = "RESISTANCE_PLAN"
)
