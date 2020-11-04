// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a SageMaker image.
func (c *Client) DescribeImage(ctx context.Context, params *DescribeImageInput, optFns ...func(*Options)) (*DescribeImageOutput, error) {
	if params == nil {
		params = &DescribeImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImage", params, optFns, addOperationDescribeImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageInput struct {

	// The name of the image to describe.
	//
	// This member is required.
	ImageName *string
}

type DescribeImageOutput struct {

	// When the image was created.
	CreationTime *time.Time

	// The description of the image.
	Description *string

	// The name of the image as displayed.
	DisplayName *string

	// When a create, update, or delete operation fails, the reason for the failure.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the image.
	ImageArn *string

	// The name of the image.
	ImageName *string

	// The status of the image.
	ImageStatus types.ImageStatus

	// When the image was last modified.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the IAM role that enables Amazon SageMaker to
	// perform tasks on your behalf.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImage{}, middleware.After)
	if err != nil {
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpDescribeImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeImage",
	}
}