package util

import (
	"errors"

	apiconstraints "github.com/open-policy-agent/frameworks/constraint/pkg/apis/constraints"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// EnforcementAction is the response we take to violations.
type EnforcementAction string

// The set of possible responses to policy violations.
const (
	Deny         EnforcementAction = "deny"
	Dryrun       EnforcementAction = "dryrun"
	Warn         EnforcementAction = "warn"
	Scoped       EnforcementAction = "scoped"
	Unrecognized EnforcementAction = "unrecognized"
)

const (
	// WebhookEnforcementPoint is the enforcement point for admission.
	WebhookEnforcementPoint = "validation.gatekeeper.sh"

	// AuditEnforcementPoint is the enforcement point for audit.
	AuditEnforcementPoint = "audit.gatekeeper.sh"

	// GatorEnforcementPoint is the enforcement point for gator cli.
	GatorEnforcementPoint = "gator.gatekeeper.sh"

	// VAP enforcement point for ValidatingAdmissionPolicy.
	VAPEnforcementPoint = "vap.k8s.io"

	// AllEnforcementPoints indicates all enforcement points.
	AllEnforcementPoints = "*"
)

var supportedEnforcementPoints = []string{WebhookEnforcementPoint, AuditEnforcementPoint, GatorEnforcementPoint, VAPEnforcementPoint}

var supportedEnforcementActions = []EnforcementAction{Deny, Dryrun, Warn, Scoped}

var supportedScopedActions = []EnforcementAction{Deny, Dryrun, Warn}

// KnownEnforcementActions are all defined EnforcementActions.
var KnownEnforcementActions = []EnforcementAction{Deny, Dryrun, Warn, Scoped, Unrecognized}

// ErrEnforcementAction indicates the passed EnforcementAction is not valid.
var ErrEnforcementAction = errors.New("unrecognized enforcementAction")

// ErrInvalidSpecEnforcementAction indicates that we were unable to parse the
// spec.enforcementAction field as it was not a string.
var ErrInvalidSpecEnforcementAction = errors.New("spec.enforcementAction must be a string")

var ErrUnrecognizedEnforcementPoint = errors.New("unrecognized enforcement points")

var ErrInvalidSpecScopedEnforcementAction = errors.New("spec.scopedEnforcementAction must be in the format of []{action: string, enforcementPoints: []{name: string}}")

func ValidateEnforcementAction(input EnforcementAction, item map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func ValidateScopedEnforcementAction(item map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// validating scopedEnforcementActions

func GetScopedEnforcementAction(item map[string]interface{}) (*[]apiconstraints.ScopedEnforcementAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertToScopedEnforcementActions(object interface{}) (*[]apiconstraints.ScopedEnforcementAction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetEnforcementAction(item map[string]interface{}) (EnforcementAction, error) {
	_ = "STUB: not implemented"
	return *new(EnforcementAction), nil
}

// default enforcementAction is deny

// validating enforcement action - if it is not deny or dryrun or scoped, we are classifying as unrecognized

func ScopedActionForEP(enforcementPoint string, u *unstructured.Unstructured) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func enforcementPointEnabled(scopedEnforcementAction apiconstraints.ScopedEnforcementAction, enforcementPoint string) bool {
	_ = "STUB: not implemented"
	return false
}
