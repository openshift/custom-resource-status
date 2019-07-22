package conditions

import "testing"

func TestSetStatusCondition(t *testing.T) {
	testCases := []struct {
		name               string
		testCondition      Condition
		startConditions    *[]Condition
		expectedConditions *[]Condition
	}{
		{
			name: "add when empty",
			testCondition: Condition{
				Type:    ConditionAvailable,
				Status:  "True",
				Reason:  "Testing",
				Message: "Basic message",
			},
			startConditions: &[]Condition{},
			expectedConditions: &[]Condition{
				Condition{
					Type:    ConditionAvailable,
					Status:  "True",
					Reason:  "Testing",
					Message: "Basic message",
				},
			},
		},
		{
			name: "add to conditions",
			testCondition: Condition{
				Type:    ConditionAvailable,
				Status:  "True",
				Reason:  "TestingAvailableTrue",
				Message: "Available condition true",
			},
			startConditions: &[]Condition{
				Condition{
					Type:    ConditionDegraded,
					Status:  "False",
					Reason:  "TestingDegradedFalse",
					Message: "Degraded condition false",
				},
			},
			expectedConditions: &[]Condition{
				Condition{
					Type:    ConditionAvailable,
					Status:  "True",
					Reason:  "TestingAvailableTrue",
					Message: "Available condition true",
				},
				Condition{
					Type:    ConditionDegraded,
					Status:  "False",
					Reason:  "TestingDegradedFalse",
					Message: "Degraded condition false",
				},
			},
		},
		{
			name: "replace condition",
			testCondition: Condition{
				Type:    ConditionDegraded,
				Status:  "True",
				Reason:  "TestingDegradedTrue",
				Message: "Degraded condition true",
			},
			startConditions: &[]Condition{
				Condition{
					Type:    ConditionDegraded,
					Status:  "False",
					Reason:  "TestingDegradedFalse",
					Message: "Degraded condition false",
				},
			},
			expectedConditions: &[]Condition{
				Condition{
					Type:    ConditionDegraded,
					Status:  "True",
					Reason:  "TestingDegradedTrue",
					Message: "Degraded condition true",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SetStatusCondition(tc.startConditions, tc.testCondition)
			compareConditions(t, tc.startConditions, tc.expectedConditions)
		})
	}

	return
}

func TestRemoveStatusCondition(t *testing.T) {
	testCases := []struct {
		name               string
		testConditionType  ConditionType
		startConditions    *[]Condition
		expectedConditions *[]Condition
	}{
		{
			name:               "remove when empty",
			testConditionType:  ConditionAvailable,
			startConditions:    &[]Condition{},
			expectedConditions: &[]Condition{},
		},
		{
			name:              "basic remove",
			testConditionType: ConditionAvailable,
			startConditions: &[]Condition{
				Condition{
					Type:    ConditionAvailable,
					Status:  "True",
					Reason:  "TestingAvailableTrue",
					Message: "Available condition true",
				},
				Condition{
					Type:    ConditionDegraded,
					Status:  "False",
					Reason:  "TestingDegradedFalse",
					Message: "Degraded condition false",
				},
			},
			expectedConditions: &[]Condition{
				Condition{
					Type:    ConditionDegraded,
					Status:  "False",
					Reason:  "TestingDegradedFalse",
					Message: "Degraded condition false",
				},
			},
		},
		{
			name:              "remove last condition",
			testConditionType: ConditionAvailable,
			startConditions: &[]Condition{
				Condition{
					Type:    ConditionAvailable,
					Status:  "True",
					Reason:  "TestingAvailableTrue",
					Message: "Available condition true",
				},
			},
			expectedConditions: &[]Condition{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			RemoveStatusCondition(tc.startConditions, tc.testConditionType)
			compareConditions(t, tc.startConditions, tc.expectedConditions)
		})
	}

	return
}

func compareConditions(t *testing.T, gotConditions *[]Condition, expectedConditions *[]Condition) {
	for _, expectedCondition := range *expectedConditions {
		testCondition := FindStatusCondition(*gotConditions, expectedCondition.Type)
		if testCondition == nil {
			t.Errorf("Condition type '%v' not found in '%v'", expectedCondition.Type, *gotConditions)
		}
		if testCondition.Status != expectedCondition.Status {
			t.Errorf("Unexpected status '%v', expected '%v'", testCondition.Status, expectedCondition.Status)
		}
		if testCondition.Message != expectedCondition.Message {
			t.Errorf("Unexpected message '%v', expected '%v'", testCondition.Message, expectedCondition.Message)
		}
	}
}
