package main

//
//import (
//	"fmt"
//	appsv1 "k8s.io/api/apps/v1"
//	appsv1beta1 "k8s.io/api/apps/v1beta1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apimachinery/pkg/util/intstr"
//	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
//	aascheme "k8s.io/kubernetes/pkg/api/legacyscheme"
//	"k8s.io/kubernetes/pkg/apis/apps"
//	"k8s.io/kubernetes/pkg/apis/apps/v1"
//	"k8s.io/kubernetes/pkg/apis/apps/v1beta1"
//	"k8s.io/kubernetes/pkg/apis/apps/v1beta2"
//)
//
//func main() {
//	// 第一部分：实例化一个空的 Scheme 资源注册表，将 v1beta1、v1的资源版本及内部版本的 Deployment 资源注册到 Scheme 资源注册表中。
//	//scheme := runtime.NewScheme()
//	//scheme.Scheme.AddKnownTypes(appsv1beta1.SchemeGroupVersion, &appsv1beta1.Deployment{})
//	//scheme.Scheme.AddKnownTypes(appsv1.SchemeGroupVersion, &appsv1.Deployment{})
//	//scheme.Scheme.AddKnownTypes(apps.SchemeGroupVersion, &appsv1.Deployment{})
//	//metav1.AddToGroupVersion(scheme.Scheme, appsv1beta1.SchemeGroupVersion)
//	//metav1.AddToGroupVersion(scheme.Scheme, appsv1.SchemeGroupVersion)
//
//	// 第二部分：实例化 v1beta1Deployment 资源对象，
//	var res int32
//	res = 4
//	v1beta1Deployment := &appsv1beta1.Deployment{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "Deployment",
//			APIVersion: "apps/v1beta1",
//		},
//		Spec: appsv1beta1.DeploymentSpec{
//			Strategy: appsv1beta1.DeploymentStrategy{
//				Type: appsv1beta1.RollingUpdateDeploymentStrategyType,
//				RollingUpdate: &appsv1beta1.RollingUpdateDeployment{
//					MaxUnavailable: func() *intstr.IntOrString { i := intstr.FromInt(0); return &i }(),
//					MaxSurge:       func() *intstr.IntOrString { i := intstr.FromInt(0); return &i }(),
//				},
//			},
//			Replicas:             &res,
//			RevisionHistoryLimit: &res,
//		},
//	}
//
//	// 通过 scheme.ConvertToVersion 将其转换为内部版本
//	// v1beta1 -> __internal
//	fmt.Println("before v1beta1 to internal")
//	objInternal, err := aascheme.Scheme.ConvertToVersion(v1beta1Deployment, apps.SchemeGroupVersion)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("GVK", objInternal.GetObjectKind().GroupVersionKind().String())
//	// output:
//	// GVK: /, Kind =
//
//	// 通过 scheme.ConvertToVersion 转换为目标资源版本，并通过断言的方式来判断是否转换成功。
//	// __internal -> v1
//	objV1, err := aascheme.Scheme.ConvertToVersion(objInternal, appsv1.SchemeGroupVersion)
//	if err != nil {
//		panic(err)
//	}
//	v1Deployment, ok := objV1.(*appsv1.Deployment)
//	if !ok {
//		panic("Got wrong type")
//	}
//	fmt.Println("GVK: ", v1Deployment.GetObjectKind().GroupVersionKind().String())
//	//output:
//	//GVK: apps/v1, Kind = Deployment
//	fmt.Println("v1 is ", v1Deployment)
//
//	//PrintScheme()
//	fmt.Println("------------------------")
//
//}
//
//func PrintScheme() {
//	fmt.Printf("aa")
//	for key, value := range aascheme.Scheme.AllKnownTypes() {
//		fmt.Printf("key is %v, value is %v", key, value)
//	}
//}
//
//func init() {
//	Install(aascheme.Scheme)
//}
//
//func Install(scheme *runtime.Scheme) {
//	//utilruntime.Must(core.AddToScheme(scheme))
//	//utilruntime.Must(v1.AddToScheme(scheme))
//	//utilruntime.Must(scheme.SetVersionPriority(v1.SchemeGroupVersion))
//	utilruntime.Must(apps.AddToScheme(scheme))
//	utilruntime.Must(v1beta1.AddToScheme(scheme))
//	utilruntime.Must(v1beta2.AddToScheme(scheme))
//	utilruntime.Must(v1.AddToScheme(scheme))
//	utilruntime.Must(scheme.SetVersionPriority(v1.SchemeGroupVersion, v1beta2.SchemeGroupVersion, v1beta1.SchemeGroupVersion))
//}
