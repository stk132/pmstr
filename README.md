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

cfg := aws.NewConfig()
	
client := pmstr.New(ssm.New(cfg))
//Get StringType Paramter
param, err := client.Get("parameter_name").AsString()

//Get StringListType Parameter
params, err := client.Get("string_list_parameter_name").AsStringList()

//Get SecureStringType Parameter
secureParam, err := client.Get("secure_parameter_name")).AsSecureString()
```


## supported API

- ssm.GetParameter