package fixtures

import (
	"testing"

	expansionunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/expansion/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/match"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type TemplateData struct {
	Name              string
	Apply             []match.ApplyTo
	Source            string
	GenGVK            expansionunversioned.GeneratedGVK
	EnforcementAction string
}

func NewTemplate(data *TemplateData) *expansionunversioned.ExpansionTemplate {
	_ = "STUB: not implemented"
	return nil
}

func LoadFixture(f string, t *testing.T) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

func LoadTemplate(f string, t *testing.T) *expansionunversioned.ExpansionTemplate {
	_ = "STUB: not implemented"
	return nil
}

func LoadAssign(f string, t *testing.T) types.Mutator {
	_ = "STUB: not implemented"
	return *new(types.Mutator)
}

func LoadAssignImage(f string, t *testing.T) types.Mutator {
	_ = "STUB: not implemented"
	return *new(types.Mutator)
}

func LoadAssignMeta(f string, t *testing.T) types.Mutator {
	_ = "STUB: not implemented"
	return *new(types.Mutator)
}

func convertUnstructuredToTyped(u *unstructured.Unstructured, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func TestTemplate(name string, applyID, genID int) *expansionunversioned.ExpansionTemplate {
	_ = "STUB: not implemented"
	return nil
}

func TempMultApply() *expansionunversioned.ExpansionTemplate { _ = "STUB: not implemented"; return nil }
