package pmstr

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)

type pmstrGetParameters struct {
	p *Pmstr
	input *ssm.GetParametersInput
}

func (p *Pmstr) GetParameters() *pmstrGetParameters {
	input := &ssm.GetParametersInput{
		Names:          nil,
		WithDecryption: aws.Bool(false),
	}
	return &pmstrGetParameters{
		p:     p,
		input: input,
	}
}


func (pgps *pmstrGetParameters) SetNames(names []string) *pmstrGetParameters {
	n := make([]*string, len(names))
	for i, v := range names {
		n[i] = aws.String(v)
	}
	pgps.input.SetNames(n)
	return pgps
}

func (pgps *pmstrGetParameters) SetWithDecryption(b bool) *pmstrGetParameters {
	pgps.input.SetWithDecryption(b)
	return pgps
}

func (pgps *pmstrGetParameters) Run() (*pmstrGetParametersOutput, error) {
	if err := pgps.input.Validate(); err != nil {
		return nil, err
	}

	output, err := pgps.p.ssmClient.GetParameters(pgps.input)
	if err != nil {
		return nil, err
	}

	if len(output.InvalidParameters) != 0 {
		errParams := make([]string, len(output.InvalidParameters))
		for i, v := range output.InvalidParameters {
			errParams[i] = aws.StringValue(v)
		}
		return nil, errors.New(fmt.Sprintf("error params: %s", strings.Join(errParams, ",")))
	}

	params := make(map[string]*ssm.Parameter, len(output.Parameters))
	for _, v := range output.Parameters {
		params[aws.StringValue(v.Name)] = v
	}

	return &pmstrGetParametersOutput{params}, nil
}

type pmstrGetParametersOutput struct {
	outputs map[string]*ssm.Parameter
}

func (pgpo *pmstrGetParametersOutput) Get(name string) *pmstrGetParametersOutputValue {
	param, ok := pgpo.outputs[name]
	if !ok {
		return &pmstrGetParametersOutputValue{
			err:       errors.New(fmt.Sprintf("paramter %s not found", name)),
			parameter: nil,
		}
	}

	return &pmstrGetParametersOutputValue{
		err:       nil,
		parameter: param,
	}
}

type pmstrGetParametersOutputValue struct {
	err error
	parameter *ssm.Parameter
}

func (pgpov *pmstrGetParametersOutputValue) errCheck(paramType string) error {
	if pgpov.err != nil {
		return pgpov.err
	}

	if *pgpov.parameter.Type != paramType {
		return errors.New(fmt.Sprintf("parameter type is not %s. actual: %s", paramType, *pgpov.parameter.Type))
	}

	return nil
}

func (pgpov *pmstrGetParametersOutputValue) AsString() (string, error) {
	if err := pgpov.errCheck(ssm.ParameterTypeString); err != nil {
		return "", err
	}

	return aws.StringValue(pgpov.parameter.Value), nil
}

func (pgpov *pmstrGetParametersOutputValue) AsStringList() ([]string, error) {
	if err := pgpov.errCheck(ssm.ParameterTypeStringList); err != nil {
		return nil, err
	}

	val := aws.StringValue(pgpov.parameter.Value)
	return strings.Split(val, ","), nil
}

func (pgpov *pmstrGetParametersOutputValue) AsSecureString() (string, error) {
	if err := pgpov.errCheck(ssm.ParameterTypeSecureString); err != nil {
		return "", err
	}

	return aws.StringValue(pgpov.parameter.Value), nil
}