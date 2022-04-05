// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata, such as the path information about an Amazon EFS location.
func (c *Client) DescribeLocationEfs(ctx context.Context, params *DescribeLocationEfsInput, optFns ...func(*Options)) (*DescribeLocationEfsOutput, error) {
	if params == nil {
		params = &DescribeLocationEfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationEfs", params, optFns, c.addOperationDescribeLocationEfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationEfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationEfsRequest
type DescribeLocationEfsInput struct {

	// The Amazon Resource Name (ARN) of the EFS location to describe.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

// DescribeLocationEfsResponse
type DescribeLocationEfsOutput struct {

	// The time that the EFS location was created.
	CreationTime *time.Time

	// The subnet that DataSync uses to access target EFS file system. The subnet must
	// have at least one mount target for that file system. The security group that you
	// provide needs to be able to communicate with the security group on the mount
	// target in the subnet specified.
	Ec2Config *types.Ec2Config

	// The Amazon Resource Name (ARN) of the EFS location that was described.
	LocationArn *string

	// The URL of the EFS location that was described.
	LocationUri *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationEfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationEfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeLocationEfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationEfs(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeLocationEfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationEfs",
	}
}
