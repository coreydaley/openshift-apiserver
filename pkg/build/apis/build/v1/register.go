package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	corev1conversions "k8s.io/kubernetes/pkg/apis/core/v1"

	v1 "github.com/openshift/api/build/v1"
	"github.com/openshift/openshift-apiserver/pkg/build/apis/build"
)

var (
	localSchemeBuilder = runtime.NewSchemeBuilder(
		build.Install,
		v1.Install,
		corev1conversions.AddToScheme,
		AddFieldSelectorKeyConversions,
		RegisterDefaults,
	)
	Install = localSchemeBuilder.AddToScheme
)
