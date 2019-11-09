package pmstr

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"strings"
)

var (
	ErrNotString        = errors.New("parameter type is not String")
	ErrNotStringList    = errors.New("parameter type is not StringList")
	ErrNotSecureString  = errors.New("parameter type is not SecureString")
	ErrUnknownParamType = errors.New("unknown parameter type")
	ErrMap              = map[string]error{
		ssm.ParameterTypeString:       ErrNotString,
		ssm.ParameterTypeStringList:   ErrNotStringList,
		ssm.ParameterTypeSecureString: ErrNotSecureString,
	}
)

type Pmstr struct {
	ssmClient ssmiface.SSMAPI
}

func New(ssmClient *ssm.SSM) *Pmstr {
	return NewFromSsmiface(ssmClient)
}

func NewFromSsmiface(ssmClient ssmiface.SSMAPI) *Pmstr {
	return &Pmstr{ssmClient}
}

func (p *Pmstr) Get(parameterName string) *pmstrGetParamInput {
	paramInput := &ssm.GetParameterInput{Name: aws.String(parameterName), WithDecryption: aws.Bool(false)}
	return newPmstrGetParamInput(p, paramInput)
}

type pmstrGetParamInput struct {
	p          *Pmstr
	paramInput *ssm.GetParameterInput
}

func newPmstrGetParamInput(p *Pmstr, paramInput *ssm.GetParameterInput) *pmstrGetParamInput {
	return &pmstrGetParamInput{
		p:          p,
		paramInput: paramInput,
	}
}

func (pi *pmstrGetParamInput) AsString() (string, error) {
	output, err := pi.Output()
	if err != nil {
		return "", err
	}
	return output.AsString()
}

func (pi *pmstrGetParamInput) AsStringList() ([]string, error) {
	output, err := pi.Output()
	if err != nil {
		return nil, err
	}
	return output.AsStringList()
}

func (pi *pmstrGetParamInput) AsSecureString() (string, error) {
	pi.paramInput.WithDecryption = aws.Bool(true)
	output, err := pi.Output()
	if err != nil {
		return "", err
	}
	return output.AsSecureString()
}

func (pi *pmstrGetParamInput) Output() (*PmstrGetParameterOutput, error) {
	output, err := pi.p.ssmClient.GetParameter(pi.paramInput)
	if err != nil {
		return nil, err
	}
	return &PmstrGetParameterOutput{output}, nil
}

type PmstrGetParameterOutput struct {
	Output *ssm.GetParameterOutput
}

func (pgp *PmstrGetParameterOutput) AsString() (string, error) {
	if err := paramTypeCheck(*pgp.Output.Parameter.Type, ssm.ParameterTypeString); err != nil {
		return "", err
	}
	return aws.StringValue(pgp.Output.Parameter.Value), nil
}

func (pgp *PmstrGetParameterOutput) AsStringList() ([]string, error) {
	if err := paramTypeCheck(*pgp.Output.Parameter.Type, ssm.ParameterTypeStringList); err != nil {
		return nil, err
	}

	return strings.Split(aws.StringValue(pgp.Output.Parameter.Value), ","), nil
}

func (pgp *PmstrGetParameterOutput) AsSecureString() (string, error) {
	if err := paramTypeCheck(*pgp.Output.Parameter.Type, ssm.ParameterTypeSecureString); err != nil {
		return "", err
	}
	return aws.StringValue(pgp.Output.Parameter.Value), nil
}

func paramTypeCheck(gotParamType, wontParamType string) error {
	if _, ok := ErrMap[wontParamType]; !ok {
		return ErrUnknownParamType
	}

	if gotParamType != wontParamType {
		err, _ := ErrMap[wontParamType]
		return err
	}
	return nil
}
