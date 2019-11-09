package pmstr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"strings"
)
//put parameter as string
func (p *Pmstr) PutString(name, value string) *pmstrPutParameterInput {
	input := &ssm.PutParameterInput{
		Name:  aws.String(name),
		Type:  aws.String(ssm.ParameterTypeString),
		Value: aws.String(value),
	}
	return newPutParameterInput(p, input)
}
//put parameter as string list
func (p *Pmstr) PutStringList(name string, value []string) *pmstrPutParameterInput {
	input := &ssm.PutParameterInput{
		Name:  aws.String(name),
		Type:  aws.String(ssm.ParameterTypeStringList),
		Value: aws.String(strings.Join(value, ",")),
	}
	return newPutParameterInput(p, input)
}
//put parameter as secure string
func (p *Pmstr) PutSecureString(name, value string) *pmstrPutParameterInput {
	input := &ssm.PutParameterInput{
		Name:  aws.String(name),
		Type:  aws.String(ssm.ParameterTypeSecureString),
		Value: aws.String(value),
	}
	return newPutParameterInput(p, input)
}

type pmstrPutParameterInput struct {
	p *Pmstr
	*ssm.PutParameterInput
}

func newPutParameterInput(p *Pmstr, input *ssm.PutParameterInput) *pmstrPutParameterInput {
	return &pmstrPutParameterInput{
		p:     p,
		PutParameterInput: input,
	}
}

func (pp *pmstrPutParameterInput) Description(description string) *pmstrPutParameterInput {
	pp.SetDescription(description)
	return pp
}

func (pp *pmstrPutParameterInput) AllowedPattern(allowedPattern string) *pmstrPutParameterInput {
	pp.SetAllowedPattern(allowedPattern)
	return pp
}

func (pp *pmstrPutParameterInput) KeyId(keyId string) *pmstrPutParameterInput {
	pp.SetKeyId(keyId)
	return pp
}

func (pp *pmstrPutParameterInput) Name(name string) *pmstrPutParameterInput {
	pp.SetName(name)
	return pp
}

func (pp *pmstrPutParameterInput) Value(value string) *pmstrPutParameterInput {
	pp.SetValue(value)
	return pp
}

func (pp *pmstrPutParameterInput) Type(t string) *pmstrPutParameterInput {
	pp.SetType(t)
	return pp
}

func (pp *pmstrPutParameterInput) Tier(tier string) *pmstrPutParameterInput {
	pp.SetTier(tier)
	return pp
}

func (pp *pmstrPutParameterInput) Tags(tags []*ssm.Tag) *pmstrPutParameterInput {
	pp.SetTags(tags)
	return pp
}

func (pp *pmstrPutParameterInput) Policies(policies string) *pmstrPutParameterInput {
	pp.SetPolicies(policies)
	return pp
}

func (pp *pmstrPutParameterInput) Overwrite(b bool) *pmstrPutParameterInput {
	pp.SetOverwrite(b)
	return pp
}

//execute ssm.PutParameter
func (pp *pmstrPutParameterInput) Do() (*ssm.PutParameterOutput, error) {
	if err := pp.Validate(); err != nil {
		return nil, err
	}
	return pp.p.ssmClient.PutParameter(pp.PutParameterInput)
}