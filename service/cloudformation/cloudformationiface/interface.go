// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudformationiface provides an interface for the AWS CloudFormation.
package cloudformationiface

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// CloudFormationAPI is the interface type for cloudformation.CloudFormation.
type CloudFormationAPI interface {
	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
}
