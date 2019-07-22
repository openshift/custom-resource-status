Custom Resource Status
======================

The purpose of this project is to provide some level of standardization and
best-practices with respect to managing the status of custom resources. This project
steals, err draws from:

* [Cluster Version Operator (CVO)](https://github.com/openshift/cluster-version-operator)
  that manages essential OpenShift operators.
* [ClusterOperator Custom Resource](https://github.com/openshift/cluster-version-operator/blob/master/docs/dev/clusteroperator.md#what-should-an-operator-report-with-clusteroperator-custom-resource)
  that exists for operators managed by CVO to communicate their status.
* [openshift/library-go ClusterOperator status helpers](https://github.com/openshift/library-go/blob/master/pkg/config/clusteroperator/v1helpers/status.go)
  that makes it easy to manage the status on a ClusterOperator resource.

The goal here is to prescribe, without mandate, how to meaningfully populate the
status of the Custom Resources your operator manages. Types, constants, and
functions are provided for the following:

* [Conditions](#Conditions)
* [Object References](#Object-References)

# Conditions

Provides:

* `Condition` type as specified in the [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
* `ConditionType` and generally useful constants for this type (ie. "Available",
    "Progressing", "Degraded", and "Upgradeable")
* Functions for setting, removing, finding, and evaluating conditions.

To use, simply add `Conditions` to your Custom Resource Status struct like:

```
// ExampleAppStatus defines the observed state of ExampleApp
type ExampleAppStatus struct {
  ...
  // conditions describes the state of the operator's reconciliation functionality.
  // +patchMergeKey=type
  // +patchStrategy=merge
  // +optional
  // Conditions is a list of conditions related to operator reconciliation
  Conditions []conditions.Condition `json:"conditions,omitempty"  patchStrategy:"merge" patchMergeKey:"type"`
}
```

Then, as appropriate in your Reconcile function, use
`conditions.SetStatusConditions` like:

```
instance := &examplev1alpha1.ExampleApp{}
err := r.client.Get(context.TODO(), request.NamespacedName, instance)
...handle err

conditions.SetStatusCondition(&instance.Status.Conditions, conditions.Condition{
  Type:   conditions.ConditionAvailable,
  Status: corev1.ConditionFalse,
  Reason: "ReconcileStarted",
  Message: "Reconciling resource"
})

// Update the status
err = r.client.Status().Update(context.TODO(), instance)
...handle err
```


# Object References

The `ObjectReference` type is provided by Kubernetes Core API
`"k8s.io/api/core/v1"` but the functions to set and find an `ObjectReference`
are provided. This is useful if you would like to include in the Status of your
Custom Resource a list of objects that are managed by your operator (ie.
Deployments, Services, other Custom Resources, etc.).

For example, we can add `RelatedObjects` to our Custom Resource Status like:

```
// ExampleAppStatus defines the observed state of ExampleApp
type ExampleAppStatus struct {
  ...
  // RelatedObjects is a list of objects that are "interesting" or related to this operator.
  RelatedObjects []corev1.ObjectReference `json:"relatedObjects,omitempty"`
}
```

Then, through Reconcile, when an object we manage has been found we can add it to
the `RelatedObjects` slice.

```
found := &someAPI.SomeObject{}
err := r.client.Get(context.TODO(), types.NamespacedName{Name: object.Name, Namespace: object.Namespace}, found)
...handle err

objectreferences.SetObjectReference(&instance.Status.RelatedObjects, corev1.ObjectReference{
  APIVersion: found.TypeMeta.APIVersion,
  Kind:       found.TypeMeta.Kind,
  Namespace:  found.Namespace,
  Name:       found.Name,
})

// Update the status
err = r.client.Status().Update(context.TODO(), instance)
...handle err
```

# TODO

* Add unit tests
* Add support for adding reconciled version to the status. This would be useful
    for communicating what version has been reconciled (ie. etcd cluster is at
    version v3.3.13)
