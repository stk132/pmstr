# pmstr

## what is pmstr?

aws-sdk-go ssm Parameter Store API wrapper

## usage

```go
import (
    "github.com/stk132/pmstr"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ssm"
)

sess := session.New(aws.NewConfig())
	
client := pmstr.New(ssm.New(sess))
//Get StringType Paramter
param, err := client.Get("parameter_name").AsString()

//Get StringListType Parameter
params, err := client.Get("string_list_parameter_name").AsStringList()

//Get SecureStringType Parameter
secureParam, err := client.Get("secure_parameter_name").AsSecureString()

//call GetParameters
param,s err := client.GetParameters().setNames([]string{"one", "two", "three"}).Run()

oneValue, err := params["one"].AsString() // or AsStringList, AsSecureString
twoValue, err := params["two"].AsString()
threeValue, err := params["threee"]


//Put StringType Parameter
output, err := client.PutString("parameter_name", "parameter_value").Do()

//Put StringListType Parameter
output, err := client.PutStringList("parameter_name", []string{"value1", "value2"}).Do()

//Put SecureStringType Parameter
output, err := client.PutSecureString("parameter_name", "parameter_value").Do()

//set parameter by method chain
output, err := client.PutString("parameter_name", "parameter_value")
    .Description("description")
    .Overwrite(true)
    .Policies("policies")
    .Do()
```


## supported API

- ssm.GetParameter
- ssm.GetParameters
- ssm.PutParameter