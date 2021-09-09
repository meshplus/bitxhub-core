package governance

type GovernanceStatus string
type EventType string

const (
	REGISTERED = 0
	APPROVED   = 1

	GovernanceRegisting    GovernanceStatus = "registering"
	GovernanceAvailable    GovernanceStatus = "available"
	GovernanceUnavailable  GovernanceStatus = "unavailable"
	GovernanceUpdating     GovernanceStatus = "updating"
	GovernanceFreezing     GovernanceStatus = "freezing"
	GovernanceActivating   GovernanceStatus = "activating"
	GovernanceFrozen       GovernanceStatus = "frozen"
	GovernanceLogouting    GovernanceStatus = "logouting"
	GovernanceBinding      GovernanceStatus = "binding"
	GovernanceUnbinding    GovernanceStatus = "unbinding"
	GovernanceBindable     GovernanceStatus = "bindable"
	GovernanceForbidden    GovernanceStatus = "forbidden"
	GovernanceTransferring GovernanceStatus = "transfering"

	EventRegister EventType = "register"
	EventUpdate   EventType = "update"
	EventFreeze   EventType = "freeze"
	EventActivate EventType = "activate"
	EventLogout   EventType = "logout"
	EventApprove  EventType = "approve"
	EventReject   EventType = "reject"
	EventBind     EventType = "bind"
	EventUnbind   EventType = "unbind"
	EventTransfer EventType = "transfer"
)

type RegisterResult struct {
	IsRegistered bool   `json:"is_registered"`
	ID           string `json:"id"`
}

type GovernanceResult struct {
	ProposalID string `json:"proposal_id"`
	Extra      []byte `json:"extra"`
}
