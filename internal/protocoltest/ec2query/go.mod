module github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.10.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.0.7
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.0-00010101000000-000000000000
	github.com/aws/smithy-go v1.8.2-0.20211102001011-cc8d2d3f851f
	github.com/google/go-cmp v0.5.6
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../../internal/endpoints/v2/
