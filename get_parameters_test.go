package pmstr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"reflect"
	"strings"
	"testing"
)

func TestPmstr_GetParameters(t *testing.T) {
	mockClient := &mockSSMClient{
		mockGetParametersOutput: &mockGetParametersOutput{
			output: &ssm.GetParametersOutput{
				InvalidParameters: nil,
				Parameters:        []*ssm.Parameter{
					&ssm.Parameter{
						ARN:              nil,
						LastModifiedDate: nil,
						Name:             aws.String("test"),
						Selector:         nil,
						SourceResult:     nil,
						Type:             aws.String(ssm.ParameterTypeString),
						Value:            aws.String("test"),
						Version:          nil,
					},
				},
			},
			err:    nil,
		},
	}

	client := NewFromSsmiface(mockClient)
	output, err := client.GetParameters().SetNames([]string{"test"}).Run()
	if err != nil {
		t.Fatal(err)
	}

	val, err := output.Get("test").AsString()
	if err != nil {
		t.Fatal(err)
	}

	if val != "test" {
		t.Errorf("wont: test, got: %s", val)
	}

}

func TestPmstrGetParametersOutput(t *testing.T) {
	sut := &pmstrGetParametersOutput{
		outputs: map[string]*ssm.Parameter{
			"test": {
				ARN:              nil,
				LastModifiedDate: nil,
				Name:             aws.String("test"),
				Selector:         nil,
				SourceResult:     nil,
				Type:             aws.String(ssm.ParameterTypeString),
				Value:            aws.String("test"),
				Version:          nil,
			},
		},
	}

	actual := sut.Get("test")
	if actual.err != nil {
		t.Errorf("should not be err")
	}

	if *actual.parameter.Value != "test" {
		t.Errorf("wont: test, got: %s", *actual.parameter.Value)
	}

	actual2 := sut.Get("not found")
	if actual2.err == nil {
		t.Error("should be err")
	}
}


func TestPmstrGetParametersOutputValue_errCheck(t *testing.T) {
	sut := &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             aws.String("test"),
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeString),
			Value:            aws.String("test"),
			Version:          nil,
		},
	}

	tests := []struct{
		typeParam string
	} {
		{
			typeParam: ssm.ParameterTypeSecureString,
		},
		{
			typeParam: ssm.ParameterTypeStringList,
		},
	}

	for _, test := range tests {
		if err := sut.errCheck(test.typeParam); err == nil {
			t.Error("should be err")
		}
	}

	if err := sut.errCheck(ssm.ParameterTypeString); err != nil {
		t.Error("should not be err")
	}


}

func TestPmstrGetParametersOutputValue_AsString(t *testing.T) {
	wont := "test"
	sut := &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             aws.String(wont),
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeString),
			Value:            aws.String(wont),
			Version:          nil,
		},
	}

	actual, err := sut.AsString()
	if err != nil {
		t.Error(err)
	}

	if actual != wont {
		t.Errorf("wont: %s, got: %s", wont, actual)
	}

	if _, err := sut.AsStringList(); err == nil {
		t.Error("should be err")
	}

	if _, err := sut.AsSecureString(); err == nil {
		t.Error("should be err")
	}
}

func TestPmstrGetParametersOutputValue_AsStringList(t *testing.T) {
	wontList := []string{"test1", "test2", "test3"}
	wontValue := strings.Join(wontList, ",")
	sut := &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             nil,
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeStringList),
			Value:            aws.String(wontValue),
			Version:          nil,
		},
	}

	actual, err := sut.AsStringList()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, wontList) {
		t.Errorf("wont: %+v, got: %+v", wontList, actual)
	}

	if _, err := sut.AsString(); err == nil {
		t.Error("should be err")
	}

	if _, err := sut.AsSecureString(); err == nil {
		t.Error("should be err")
	}
}

func TestPmstrGetParametersOutputValue_AsSecureString(t *testing.T) {
	wont := "test"
	sut := &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             nil,
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeSecureString),
			Value:            aws.String(wont),
			Version:          nil,
		},
	}

	actual, err := sut.AsSecureString()
	if err != nil {
		t.Fatal(err)
	}

	if actual != wont {
		t.Errorf("wont: %s, got: %s", wont, actual)
	}

	if _, err := sut.AsString(); err == nil {
		t.Error("should be err")
	}

	if _, err := sut.AsStringList(); err == nil {
		t.Error("should be err")
	}
}

func TestPmstrGetParametersOutputValue_Value(t *testing.T) {
	wont := "test"
	sut := &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             nil,
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeSecureString),
			Value:            aws.String(wont),
			Version:          nil,
		},
	}

	actual, err := sut.Value()
	if err != nil {
		t.Fatal(err)
	}

	if actual != wont {
		t.Errorf("wont: %s, got: %s", wont, actual)
	}

}