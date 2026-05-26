package transform

import (
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apiserver/pkg/admission/plugin/cel"
)

const (
	matchKinds = `
	!has(params.spec) ? true: (
		!has(params.spec.match) ? true: (
			!has(params.spec.match.kinds) ? true : (
				params.spec.match.kinds.exists(groupskinds,
					(!has(groupskinds.kinds) || size(groupskinds.kinds) == 0 || "*" in groupskinds.kinds || request.kind.kind in groupskinds.kinds) &&
					(!has(groupskinds.apiGroups) || size(groupskinds.apiGroups) == 0 || "*" in groupskinds.apiGroups || request.kind.group in groupskinds.apiGroups)
				)
			)
		)
	)
	`

	// Note that switching the glob to a regex is valid because of how Gatekeeper validates the wildcard matcher
	// (with this regex: "+kubebuilder:validation:Pattern=`^(\*|\*-)?[a-z0-9]([-:a-z0-9]*[a-z0-9])?(\*|-\*)?$`").
	matchNameGlob = `
	!has(params.spec) ? true: (
		!has(params.spec.match) ? true: (
			!has(params.spec.match.name) ? true : (
				[object, oldObject].exists(obj,
					obj != null && (
						(has(obj.metadata.generateName) && obj.metadata.generateName != "" && params.spec.match.name.endsWith("*") && string(obj.metadata.generateName).matches("^" + string(params.spec.match.name).replace("*", ".*") + "$")) ||
						(has(obj.metadata.name) && string(obj.metadata.name).matches("^" + string(params.spec.match.name).replace("*", ".*") + "$"))
					)
				)
			)
		)
	)
	`

	// Note that switching the glob to a regex is valid because of how Gatekeeper validates the wildcard matcher
	// (with this regex: "+kubebuilder:validation:Pattern=`^(\*|\*-)?[a-z0-9]([-:a-z0-9]*[a-z0-9])?(\*|-\*)?$`").
	// TODO: consider using the `namespaceObject` field provided by ValidatingAdmissionPolicy.
	matchNamespacesGlob = `
	!has(params.spec) ? true: (
		!has(params.spec.match) ? true: (
			!has(params.spec.match.namespaces) ? true : (
				[object, oldObject].exists(obj,
					obj != null && (
						// cluster-scoped objects always match
						!has(obj.metadata.namespace) || obj.metadata.namespace == "" ? true : (
							params.spec.match.namespaces.exists(nsMatcher,
								(string(obj.metadata.namespace).matches("^" + string(nsMatcher).replace("*", ".*") + "$"))
							)
						)
					)
				)
			)
		)
	)
	`

	// Note that switching the glob to a regex is valid because of how Gatekeeper validates the wildcard matcher
	// (with this regex: "+kubebuilder:validation:Pattern=`^(\*|\*-)?[a-z0-9]([-:a-z0-9]*[a-z0-9])?(\*|-\*)?$`").
	// TODO: consider using the `namespaceObject` field provided by ValidatingAdmissionPolicy.
	matchExcludedNamespacesGlob = `
	!has(params.spec) ? true: (
		!has(params.spec.match) ? true: (
			!has(params.spec.match.excludedNamespaces) ? true : (
				[object, oldObject].exists(obj,
					obj != null && (
						// cluster-scoped objects always match
						!has(obj.metadata.namespace) || obj.metadata.namespace == "" ? true : (
							!params.spec.match.excludedNamespaces.exists(nsMatcher,
								(string(obj.metadata.namespace).matches("^" + string(nsMatcher).replace("*", ".*") + "$"))
							)
						)
					)
				)
			)
		)
	)
	`

	// Expression to exclude objects in globally excluded namespaces and namespace objects themselves if they're in the exclusion list from Config resource.
	matchGlobalExcludedNamespacesGlob = `
	[object, oldObject].exists(obj,
		obj != null && (
			// For namespace objects, check if the namespace name itself is in the exclusion list
			(has(obj.kind) && obj.kind == "Namespace" && has(obj.metadata.name)) ? (
				![%s].exists(nsMatcher,
					(string(obj.metadata.name).matches("^" + string(nsMatcher).replace("*", ".*") + "$"))
				)
			) : (
				// cluster-scoped objects (non-namespace) always match
				!has(obj.metadata.namespace) || obj.metadata.namespace == "" ? true : (
					![%s].exists(nsMatcher,
						(string(obj.metadata.namespace).matches("^" + string(nsMatcher).replace("*", ".*") + "$"))
					)
				)
			)
		)
	)
	`
	// Expression to exempt objects in globally exempted namespaces and namespace objects themselves if they're in the exemption list from exempt namespace flags.
	matchGlobalExemptedNamespacesGlob = `
	[object, oldObject].exists(obj,
		obj != null && (
			// For namespace objects, check if the namespace name itself is in the exemption list
			(has(obj.kind) && obj.kind == "Namespace" && has(obj.metadata.name)) ? (
				![%s].exists(nsMatcher,
					(string(obj.metadata.name).matches("^" + string(nsMatcher).replace("*", ".*") + "$")) &&
					has(obj.metadata.labels) &&
					("admission.gatekeeper.sh/ignore" in obj.metadata.labels)
				)
			) : (
				// cluster-scoped objects (non-namespace) always match
				!has(obj.metadata.namespace) || obj.metadata.namespace == "" ? true : (
					![%s].exists(nsMatcher,
						(string(obj.metadata.namespace).matches("^" + string(nsMatcher).replace("*", ".*") + "$"))
					)
				)
			)
		)
	)
	`
)

func MatchExcludedNamespacesGlobV1Beta1() admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchGlobalExcludedNamespacesGlobV1Beta1(excludedNamespaces string) admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchGlobalExemptedNamespacesGlobV1Beta1(exemptedNamespaces string) admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchExcludedNamespacesGlobCEL() []cel.ExpressionAccessor {
	_ = "STUB: not implemented"
	return nil
}

func MatchNamespacesGlobV1Beta1() admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchNamespacesGlobCEL() []cel.ExpressionAccessor { _ = "STUB: not implemented"; return nil }

func MatchNameGlobV1Beta1() admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchNameGlobCEL() []cel.ExpressionAccessor { _ = "STUB: not implemented"; return nil }

func MatchKindsV1Beta1() admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.MatchCondition)
}

func MatchKindsCEL() []cel.ExpressionAccessor { _ = "STUB: not implemented"; return nil }

func BindParamsV1Beta1() admissionregistrationv1beta1.Variable {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.Variable)
}

func BindParamsCEL() cel.NamedExpressionAccessor {
	_ = "STUB: not implemented"
	return *new(cel.NamedExpressionAccessor)
}

func BindObjectV1Beta1() admissionregistrationv1beta1.Variable {
	_ = "STUB: not implemented"
	return *new(admissionregistrationv1beta1.Variable)
}

func AllMatchersV1Beta1() []admissionregistrationv1beta1.MatchCondition {
	_ = "STUB: not implemented"
	return nil
}

func AllVariablesCEL() []cel.NamedExpressionAccessor { _ = "STUB: not implemented"; return nil }

func AllVariablesV1Beta1() []admissionregistrationv1beta1.Variable {
	_ = "STUB: not implemented"
	return nil
}
