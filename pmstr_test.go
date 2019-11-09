package pmstr

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"testing"
)

type mockGetParameterOutput struct {
	output *ssm.GetParameterOutput
	err error
}

type mockSSMClient struct {
	ssmiface.SSMAPI
	mockGetParamOutputMap map[string]*mockGetParameterOutput
}

func (m *mockSSMClient) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	o, ok := m.mockGetParamOutputMap[*input.Name]
	if !ok {
		return nil, errors.New("parameter not found")
	}
	return o.output, o.err
}

func TestPmstrGetParamInput_AsString(t *testing.T) {
	mockOutput := &mockGetParameterOutput{
		output: &ssm.GetParameterOutput{Parameter: &ssm.Parameter{
			ARN:              nil,
			LastModifiedDate: nil,
			Name:             aws.String("mock"),
			Selector:         nil,
			SourceResult:     nil,
			Type:             aws.String(ssm.ParameterTypeString),
			Value:            aws.String("mock"),
			Version:          nil,
		}},
		err:    nil,
	}
	paramMap := map[string]*mockGetParameterOutput{
		"mock": mockOutput,
	}
	mock := &mockSSMClient{mockGetParamOutputMap:paramMap}
	client := NewFromSsmiface(mock)
	got, err := client.Get("mock").AsString()
	if err != nil {
		t.Fatal(err)
	}

	if got != "mock" {
		t.Errorf("wont: %v, got: %v", "mock", got)
	}

	if _, err := client.Get("noffound").AsString(); err == nil {
		t.Errorf("should be error")
	}

	if _, err := client.Get("mock").AsSecureString(); err != ErrNotSecureString {
		t.Errorf("wont: %+v, got: %+v", ErrNotSecureString, err)
	}

	if _, err := client.Get("mock").AsStringList(); err != ErrNotStringList {
		t.Errorf("wont: %+v, got: %+v", ErrNotStringList, err)
	}

}
