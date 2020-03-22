// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the CreatePolicy operation.
type CreatePolicyInput struct {
	_ struct{} `type:"structure"`

	// The JSON document that describes the policy. policyDocument must have a minimum
	// length of 1, with a maximum length of 2048, excluding whitespace.
	//
	// PolicyDocument is a required field
	PolicyDocument *string `locationName:"policyDocument" type:"string" required:"true"`

	// The policy name.
	//
	// PolicyName is a required field
	PolicyName *string `location:"uri" locationName:"policyName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePolicyInput"}

	if s.PolicyDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyDocument"))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.PolicyDocument != nil {
		v := *s.PolicyDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the CreatePolicy operation.
type CreatePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The policy ARN.
	PolicyArn *string `locationName:"policyArn" type:"string"`

	// The JSON document that describes the policy.
	PolicyDocument *string `locationName:"policyDocument" type:"string"`

	// The policy name.
	PolicyName *string `locationName:"policyName" min:"1" type:"string"`

	// The policy version ID.
	PolicyVersionId *string `locationName:"policyVersionId" type:"string"`
}

// String returns the string representation
func (s CreatePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.PolicyArn != nil {
		v := *s.PolicyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyDocument != nil {
		v := *s.PolicyDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyName != nil {
		v := *s.PolicyName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PolicyVersionId != nil {
		v := *s.PolicyVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "policyVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreatePolicy = "CreatePolicy"

// CreatePolicyRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates an AWS IoT policy.
//
// The created policy is the default version for the policy. This operation
// creates a policy version with a version identifier of 1 and sets 1 as the
// policy's default version.
//
//    // Example sending a request using CreatePolicyRequest.
//    req := client.CreatePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreatePolicyRequest(input *CreatePolicyInput) CreatePolicyRequest {
	op := &aws.Operation{
		Name:       opCreatePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/policies/{policyName}",
	}

	if input == nil {
		input = &CreatePolicyInput{}
	}

	req := c.newRequest(op, input, &CreatePolicyOutput{})
	return CreatePolicyRequest{Request: req, Input: input, Copy: c.CreatePolicyRequest}
}

// CreatePolicyRequest is the request type for the
// CreatePolicy API operation.
type CreatePolicyRequest struct {
	*aws.Request
	Input *CreatePolicyInput
	Copy  func(*CreatePolicyInput) CreatePolicyRequest
}

// Send marshals and sends the CreatePolicy API request.
func (r CreatePolicyRequest) Send(ctx context.Context) (*CreatePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyResponse{
		CreatePolicyOutput: r.Request.Data.(*CreatePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePolicyResponse is the response type for the
// CreatePolicy API operation.
type CreatePolicyResponse struct {
	*CreatePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePolicy request.
func (r *CreatePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}