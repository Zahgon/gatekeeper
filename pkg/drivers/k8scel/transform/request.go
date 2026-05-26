package transform

import (
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/admission"
	auditinternal "k8s.io/apiserver/pkg/apis/audit"
	"k8s.io/apiserver/pkg/authentication/user"
)

func RequestToVersionedAttributes(request *admissionv1.AdmissionRequest) (*admission.VersionedAttributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FRICTION this wrapper class is excessive. Validator code should define an interface that only requires the methods it needs.
type RequestWrapper struct {
	ar               *admissionv1.AdmissionRequest
	object           runtime.Object
	oldObject        runtime.Object
	operationOptions runtime.Object
}

func NewWrapper(req *admissionv1.AdmissionRequest) (*RequestWrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this may be unnecessary, since GetOptions() may not be used by downstream
// code, but is better than doing this lazily and needing to panic if GetOptions()
// fails.

func (w *RequestWrapper) GetName() string { _ = "STUB: not implemented"; return "" }

func (w *RequestWrapper) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (w *RequestWrapper) GetResource() schema.GroupVersionResource {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionResource)
}

func (w *RequestWrapper) GetSubresource() string { _ = "STUB: not implemented"; return "" }

var opMap = map[admissionv1.Operation]admission.Operation{
	admissionv1.Create:  admission.Create,
	admissionv1.Update:  admission.Update,
	admissionv1.Delete:  admission.Delete,
	admissionv1.Connect: admission.Connect,
}

func (w *RequestWrapper) GetOperation() admission.Operation {
	_ = "STUB: not implemented"
	return *new(admission.Operation)
}

func (w *RequestWrapper) GetOperationOptions() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (w *RequestWrapper) IsDryRun() bool { _ = "STUB: not implemented"; return false }

func (w *RequestWrapper) GetObject() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (w *RequestWrapper) GetOldObject() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (w *RequestWrapper) GetKind() schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind)
}

func (w *RequestWrapper) GetUserInfo() user.Info { _ = "STUB: not implemented"; return *new(user.Info) }

func (w *RequestWrapper) AddAnnotation(_, _ string) error { _ = "STUB: not implemented"; return nil }

func (w *RequestWrapper) AddAnnotationWithLevel(_, _ string, _ auditinternal.Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *RequestWrapper) GetReinvocationContext() admission.ReinvocationContext {
	_ = "STUB: not implemented"
	return *new(admission.ReinvocationContext)
}
