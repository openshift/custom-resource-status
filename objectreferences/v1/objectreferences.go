package v1

import (
	"k8s.io/apimachinery/pkg/api/equality"

	corev1 "k8s.io/api/core/v1"
)

// SetObjectReference - updates list of object references based on newObject
func SetObjectReference(objects *[]corev1.ObjectReference, newObject corev1.ObjectReference) {
	if objects == nil {
		objects = &[]corev1.ObjectReference{}
	}
	existingObject := FindObjectReference(*objects, newObject)
	if existingObject == nil {
		*objects = append(*objects, newObject)
		return
	}
}

// RemoveObjectReference - updates list of object references to remove rmObject
func RemoveObjectReference(objects *[]corev1.ObjectReference, rmObject corev1.ObjectReference) {
	if objects == nil {
		return
	}
	newObjectReferences := []corev1.ObjectReference{}
	// TODO: this is incredibly inefficient. If the performance hit becomes a
	// problem this should be improved.
	for _, object := range *objects {
		if !equality.Semantic.DeepEqual(object, rmObject) {
			newObjectReferences = append(newObjectReferences, object)
		}
	}

	*objects = newObjectReferences
}

// FindObjectReference - finds the first ObjectReference in a slice of objects
// matching find.
func FindObjectReference(objects []corev1.ObjectReference, find corev1.ObjectReference) *corev1.ObjectReference {
	// TODO: since there is nothing preventing multiple references being added
	// with the same APIVersion/Kind, it may be worthwile in the future to
	// make it possible to get a slice of matching references
	for i := range objects {
		if equality.Semantic.DeepEqual(find, objects[i]) {
			return &objects[i]
		}
	}

	return nil
}
