package v1

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/equality"

	corev1 "k8s.io/api/core/v1"
)

func TestSetObjectReference(t *testing.T) {
	testCases := []struct {
		name         string
		testRef      corev1.ObjectReference
		startRefs    *[]corev1.ObjectReference
		expectedRefs *[]corev1.ObjectReference
	}{
		{
			name: "add when empty",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{},
			expectedRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
			},
		},
		{
			name: "simple add",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
			},
			expectedRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
			},
		},
		{
			name: "add duplicate reference",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
				UID:        "fooid",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
			},
			expectedRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
					UID:        "fooid",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetObjectReference(tc.startRefs, tc.testRef)
			if !equality.Semantic.DeepEqual(*tc.startRefs, *tc.expectedRefs) {
				t.Errorf("Unexpected object refs '%v', expected '%v'", tc.startRefs, tc.expectedRefs)
			}
		})
	}
	return
}

func TestRemoveObjectReference(t *testing.T) {
	testCases := []struct {
		name         string
		testRef      corev1.ObjectReference
		startRefs    *[]corev1.ObjectReference
		expectedRefs *[]corev1.ObjectReference
	}{
		{
			name: "remove when empty",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs:    &[]corev1.ObjectReference{},
			expectedRefs: &[]corev1.ObjectReference{},
		},
		{
			name: "simple remove",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
			},
			expectedRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
			},
		},
		{
			name: "remove last",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
			},
			expectedRefs: &[]corev1.ObjectReference{},
		},
		{
			// Not sure if this is possible by using SetObjectReference
			// but testing this anyway
			name: "remove matching",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
					UID:        "myuid",
				},
			},
			expectedRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "BarKind",
					Namespace:  "test-namespace",
					Name:       "bar",
					APIVersion: "test.example.io",
				},
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
					UID:        "myuid",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			RemoveObjectReference(tc.startRefs, tc.testRef)
			if !equality.Semantic.DeepEqual(*tc.startRefs, *tc.expectedRefs) {
				t.Errorf("Unexpected object refs '%v', expected '%v'", tc.startRefs, tc.expectedRefs)
			}
		})
	}
	return
}

func TestFindObjectReference(t *testing.T) {
	testCases := []struct {
		name        string
		testRef     corev1.ObjectReference
		startRefs   *[]corev1.ObjectReference
		expectedRef *corev1.ObjectReference
	}{
		{
			name: "simple find",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs: &[]corev1.ObjectReference{
				corev1.ObjectReference{
					Kind:       "FooKind",
					Namespace:  "test-namespace",
					Name:       "foo",
					APIVersion: "test.example.io",
				},
			},
			expectedRef: &corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
		},
		{
			name: "find when empty",
			testRef: corev1.ObjectReference{
				Kind:       "FooKind",
				Namespace:  "test-namespace",
				Name:       "foo",
				APIVersion: "test.example.io",
			},
			startRefs:   &[]corev1.ObjectReference{},
			expectedRef: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			foundRef := FindObjectReference(*tc.startRefs, tc.testRef)
			if !equality.Semantic.DeepEqual(foundRef, tc.expectedRef) {
				t.Errorf("Unexpected object ref '%v', expected '%v'", foundRef, tc.expectedRef)
			}
		})
	}
	return
}
