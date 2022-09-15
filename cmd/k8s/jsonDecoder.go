package k8s

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/json"
)

func Decode(data string) {

	var unstruct unstructured.Unstructured

	err := json.Unmarshal([]byte(data), &unstruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", unstruct)
	switch unstruct.GetKind() {
	case "List":
		fmt.Println(unstruct.GetKind())
		err := unstruct.EachListItem(func(object runtime.Object) error {
			switch object.GetObjectKind().GroupVersionKind().Kind {
			case "Deployment":
				var obj v1.Deployment
				err := runtime.DefaultUnstructuredConverter.FromUnstructured(object.(*unstructured.Unstructured).UnstructuredContent(), &obj)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("obj: ", obj)
				fmt.Printf("%#v\n", obj.Status)
			case "Service":
				var obj v12.Service
				err := runtime.DefaultUnstructuredConverter.FromUnstructured(object.(*unstructured.Unstructured).UnstructuredContent(), &obj)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("obj: ", obj)

			default:
				fmt.Println(object.GetObjectKind().GroupVersionKind().Kind)

			}
			return err
		})
		if err != nil {
			fmt.Printf("EachListItem err: %s\n", err)
		}

	default:

		fmt.Println(unstruct.GetKind())
	}

}
