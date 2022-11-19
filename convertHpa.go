package main

import (
	"fmt"
	hv1 "k8s.io/api/autoscaling/v1"
	appsv1beta2 "k8s.io/api/autoscaling/v2beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	aascheme "k8s.io/kubernetes/pkg/api/legacyscheme"

	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/apis/autoscaling/v1"
	"k8s.io/kubernetes/pkg/apis/autoscaling/v2"
	"k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1"
	"k8s.io/kubernetes/pkg/apis/autoscaling/v2beta2"
)

func main() {

	// 第二部分：实例化 v1beta1Deployment 资源对象，
	v1beta2Hpa := &appsv1beta2.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "test",
			Name:      "test",
		},
	}
	fmt.Println("v2 is ", v1beta2Hpa)
	// 通过 scheme.ConvertToVersion 将其转换为内部版本
	// v1beta1 -> __internal
	fmt.Println("before v1beta1 to internal")
	objInternal, err := aascheme.Scheme.ConvertToVersion(v1beta2Hpa, autoscaling.SchemeGroupVersion)
	if err != nil {
		panic(err)
	}
	fmt.Println("GVK", objInternal.GetObjectKind().GroupVersionKind().String())
	// output:
	// GVK: /, Kind =

	// 通过 scheme.ConvertToVersion 转换为目标资源版本，并通过断言的方式来判断是否转换成功。
	// __internal -> v1
	objV1, err := aascheme.Scheme.ConvertToVersion(objInternal, v1.SchemeGroupVersion)
	if err != nil {
		panic(err)
	}
	v1Hpa, ok := objV1.(*hv1.HorizontalPodAutoscaler)
	if !ok {
		panic("Got wrong type")
	}
	fmt.Println("GVK: ", v1Hpa.GetObjectKind().GroupVersionKind().String())
	//output:
	//GVK: apps/v1, Kind = Deployment
	fmt.Println("v1 is ", v1Hpa)

	//PrintScheme()
	fmt.Println("------------------------")

}

func PrintScheme() {
	fmt.Printf("aa")
	for key, value := range aascheme.Scheme.AllKnownTypes() {
		fmt.Printf("key is %v, value is %v", key, value)
	}
}

func init() {
	Install(aascheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(autoscaling.AddToScheme(scheme))
	utilruntime.Must(v2beta2.AddToScheme(scheme))
	utilruntime.Must(v2.AddToScheme(scheme))
	utilruntime.Must(v2beta1.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))
	// TODO: move v2 to the front of the list in 1.24
	utilruntime.Must(scheme.SetVersionPriority(v1.SchemeGroupVersion, v2.SchemeGroupVersion, v2beta1.SchemeGroupVersion, v2beta2.SchemeGroupVersion))
}
