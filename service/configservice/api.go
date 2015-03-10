package configservice

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"time"
)

// DeleteDeliveryChannelRequest generates a request for the DeleteDeliveryChannel operation.
func (c *ConfigService) DeleteDeliveryChannelRequest(input *DeleteDeliveryChannelInput) (req *aws.Request) {
	if opDeleteDeliveryChannel == nil {
		opDeleteDeliveryChannel = &aws.Operation{
			Name:       "DeleteDeliveryChannel",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteDeliveryChannel, input, nil)

	return
}

func (c *ConfigService) DeleteDeliveryChannel(input *DeleteDeliveryChannelInput) (err error) {
	req := c.DeleteDeliveryChannelRequest(input)
	err = req.Send()
	return
}

var opDeleteDeliveryChannel *aws.Operation

// DeliverConfigSnapshotRequest generates a request for the DeliverConfigSnapshot operation.
func (c *ConfigService) DeliverConfigSnapshotRequest(input *DeliverConfigSnapshotInput) (req *aws.Request, output *DeliverConfigSnapshotOutput) {
	if opDeliverConfigSnapshot == nil {
		opDeliverConfigSnapshot = &aws.Operation{
			Name:       "DeliverConfigSnapshot",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeliverConfigSnapshot, input, output)
	output = &DeliverConfigSnapshotOutput{}
	req.Data = output
	return
}

func (c *ConfigService) DeliverConfigSnapshot(input *DeliverConfigSnapshotInput) (output *DeliverConfigSnapshotOutput, err error) {
	req, out := c.DeliverConfigSnapshotRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeliverConfigSnapshot *aws.Operation

// DescribeConfigurationRecorderStatusRequest generates a request for the DescribeConfigurationRecorderStatus operation.
func (c *ConfigService) DescribeConfigurationRecorderStatusRequest(input *DescribeConfigurationRecorderStatusInput) (req *aws.Request, output *DescribeConfigurationRecorderStatusOutput) {
	if opDescribeConfigurationRecorderStatus == nil {
		opDescribeConfigurationRecorderStatus = &aws.Operation{
			Name:       "DescribeConfigurationRecorderStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeConfigurationRecorderStatus, input, output)
	output = &DescribeConfigurationRecorderStatusOutput{}
	req.Data = output
	return
}

func (c *ConfigService) DescribeConfigurationRecorderStatus(input *DescribeConfigurationRecorderStatusInput) (output *DescribeConfigurationRecorderStatusOutput, err error) {
	req, out := c.DescribeConfigurationRecorderStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeConfigurationRecorderStatus *aws.Operation

// DescribeConfigurationRecordersRequest generates a request for the DescribeConfigurationRecorders operation.
func (c *ConfigService) DescribeConfigurationRecordersRequest(input *DescribeConfigurationRecordersInput) (req *aws.Request, output *DescribeConfigurationRecordersOutput) {
	if opDescribeConfigurationRecorders == nil {
		opDescribeConfigurationRecorders = &aws.Operation{
			Name:       "DescribeConfigurationRecorders",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeConfigurationRecorders, input, output)
	output = &DescribeConfigurationRecordersOutput{}
	req.Data = output
	return
}

func (c *ConfigService) DescribeConfigurationRecorders(input *DescribeConfigurationRecordersInput) (output *DescribeConfigurationRecordersOutput, err error) {
	req, out := c.DescribeConfigurationRecordersRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeConfigurationRecorders *aws.Operation

// DescribeDeliveryChannelStatusRequest generates a request for the DescribeDeliveryChannelStatus operation.
func (c *ConfigService) DescribeDeliveryChannelStatusRequest(input *DescribeDeliveryChannelStatusInput) (req *aws.Request, output *DescribeDeliveryChannelStatusOutput) {
	if opDescribeDeliveryChannelStatus == nil {
		opDescribeDeliveryChannelStatus = &aws.Operation{
			Name:       "DescribeDeliveryChannelStatus",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeDeliveryChannelStatus, input, output)
	output = &DescribeDeliveryChannelStatusOutput{}
	req.Data = output
	return
}

func (c *ConfigService) DescribeDeliveryChannelStatus(input *DescribeDeliveryChannelStatusInput) (output *DescribeDeliveryChannelStatusOutput, err error) {
	req, out := c.DescribeDeliveryChannelStatusRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDeliveryChannelStatus *aws.Operation

// DescribeDeliveryChannelsRequest generates a request for the DescribeDeliveryChannels operation.
func (c *ConfigService) DescribeDeliveryChannelsRequest(input *DescribeDeliveryChannelsInput) (req *aws.Request, output *DescribeDeliveryChannelsOutput) {
	if opDescribeDeliveryChannels == nil {
		opDescribeDeliveryChannels = &aws.Operation{
			Name:       "DescribeDeliveryChannels",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeDeliveryChannels, input, output)
	output = &DescribeDeliveryChannelsOutput{}
	req.Data = output
	return
}

func (c *ConfigService) DescribeDeliveryChannels(input *DescribeDeliveryChannelsInput) (output *DescribeDeliveryChannelsOutput, err error) {
	req, out := c.DescribeDeliveryChannelsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeDeliveryChannels *aws.Operation

// GetResourceConfigHistoryRequest generates a request for the GetResourceConfigHistory operation.
func (c *ConfigService) GetResourceConfigHistoryRequest(input *GetResourceConfigHistoryInput) (req *aws.Request, output *GetResourceConfigHistoryOutput) {
	if opGetResourceConfigHistory == nil {
		opGetResourceConfigHistory = &aws.Operation{
			Name:       "GetResourceConfigHistory",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetResourceConfigHistory, input, output)
	output = &GetResourceConfigHistoryOutput{}
	req.Data = output
	return
}

func (c *ConfigService) GetResourceConfigHistory(input *GetResourceConfigHistoryInput) (output *GetResourceConfigHistoryOutput, err error) {
	req, out := c.GetResourceConfigHistoryRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetResourceConfigHistory *aws.Operation

// PutConfigurationRecorderRequest generates a request for the PutConfigurationRecorder operation.
func (c *ConfigService) PutConfigurationRecorderRequest(input *PutConfigurationRecorderInput) (req *aws.Request) {
	if opPutConfigurationRecorder == nil {
		opPutConfigurationRecorder = &aws.Operation{
			Name:       "PutConfigurationRecorder",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutConfigurationRecorder, input, nil)

	return
}

func (c *ConfigService) PutConfigurationRecorder(input *PutConfigurationRecorderInput) (err error) {
	req := c.PutConfigurationRecorderRequest(input)
	err = req.Send()
	return
}

var opPutConfigurationRecorder *aws.Operation

// PutDeliveryChannelRequest generates a request for the PutDeliveryChannel operation.
func (c *ConfigService) PutDeliveryChannelRequest(input *PutDeliveryChannelInput) (req *aws.Request) {
	if opPutDeliveryChannel == nil {
		opPutDeliveryChannel = &aws.Operation{
			Name:       "PutDeliveryChannel",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPutDeliveryChannel, input, nil)

	return
}

func (c *ConfigService) PutDeliveryChannel(input *PutDeliveryChannelInput) (err error) {
	req := c.PutDeliveryChannelRequest(input)
	err = req.Send()
	return
}

var opPutDeliveryChannel *aws.Operation

// StartConfigurationRecorderRequest generates a request for the StartConfigurationRecorder operation.
func (c *ConfigService) StartConfigurationRecorderRequest(input *StartConfigurationRecorderInput) (req *aws.Request) {
	if opStartConfigurationRecorder == nil {
		opStartConfigurationRecorder = &aws.Operation{
			Name:       "StartConfigurationRecorder",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opStartConfigurationRecorder, input, nil)

	return
}

func (c *ConfigService) StartConfigurationRecorder(input *StartConfigurationRecorderInput) (err error) {
	req := c.StartConfigurationRecorderRequest(input)
	err = req.Send()
	return
}

var opStartConfigurationRecorder *aws.Operation

// StopConfigurationRecorderRequest generates a request for the StopConfigurationRecorder operation.
func (c *ConfigService) StopConfigurationRecorderRequest(input *StopConfigurationRecorderInput) (req *aws.Request) {
	if opStopConfigurationRecorder == nil {
		opStopConfigurationRecorder = &aws.Operation{
			Name:       "StopConfigurationRecorder",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opStopConfigurationRecorder, input, nil)

	return
}

func (c *ConfigService) StopConfigurationRecorder(input *StopConfigurationRecorderInput) (err error) {
	req := c.StopConfigurationRecorderRequest(input)
	err = req.Send()
	return
}

var opStopConfigurationRecorder *aws.Operation

type ConfigExportDeliveryInfo struct {
	LastAttemptTime    *time.Time `locationName:"lastAttemptTime" type:"timestamp" timestampFormat:"unix" json:"lastAttemptTime,omitempty"`
	LastErrorCode      *string    `locationName:"lastErrorCode" type:"string" json:"lastErrorCode,omitempty"`
	LastErrorMessage   *string    `locationName:"lastErrorMessage" type:"string" json:"lastErrorMessage,omitempty"`
	LastStatus         *string    `locationName:"lastStatus" type:"string" json:"lastStatus,omitempty"`
	LastSuccessfulTime *time.Time `locationName:"lastSuccessfulTime" type:"timestamp" timestampFormat:"unix" json:"lastSuccessfulTime,omitempty"`

	metadataConfigExportDeliveryInfo `json:"-", xml:"-"`
}

type metadataConfigExportDeliveryInfo struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ConfigStreamDeliveryInfo struct {
	LastErrorCode        *string    `locationName:"lastErrorCode" type:"string" json:"lastErrorCode,omitempty"`
	LastErrorMessage     *string    `locationName:"lastErrorMessage" type:"string" json:"lastErrorMessage,omitempty"`
	LastStatus           *string    `locationName:"lastStatus" type:"string" json:"lastStatus,omitempty"`
	LastStatusChangeTime *time.Time `locationName:"lastStatusChangeTime" type:"timestamp" timestampFormat:"unix" json:"lastStatusChangeTime,omitempty"`

	metadataConfigStreamDeliveryInfo `json:"-", xml:"-"`
}

type metadataConfigStreamDeliveryInfo struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ConfigurationItem struct {
	ARN                          *string             `locationName:"arn" type:"string" json:"arn,omitempty"`
	AccountID                    *string             `locationName:"accountId" type:"string" json:"accountId,omitempty"`
	AvailabilityZone             *string             `locationName:"availabilityZone" type:"string" json:"availabilityZone,omitempty"`
	Configuration                *string             `locationName:"configuration" type:"string" json:"configuration,omitempty"`
	ConfigurationItemCaptureTime *time.Time          `locationName:"configurationItemCaptureTime" type:"timestamp" timestampFormat:"unix" json:"configurationItemCaptureTime,omitempty"`
	ConfigurationItemMD5Hash     *string             `locationName:"configurationItemMD5Hash" type:"string" json:"configurationItemMD5Hash,omitempty"`
	ConfigurationItemStatus      *string             `locationName:"configurationItemStatus" type:"string" json:"configurationItemStatus,omitempty"`
	ConfigurationStateID         *string             `locationName:"configurationStateId" type:"string" json:"configurationStateId,omitempty"`
	RelatedEvents                []*string           `locationName:"relatedEvents" type:"list" json:"relatedEvents,omitempty"`
	Relationships                []*Relationship     `locationName:"relationships" type:"list" json:"relationships,omitempty"`
	ResourceCreationTime         *time.Time          `locationName:"resourceCreationTime" type:"timestamp" timestampFormat:"unix" json:"resourceCreationTime,omitempty"`
	ResourceID                   *string             `locationName:"resourceId" type:"string" json:"resourceId,omitempty"`
	ResourceType                 *string             `locationName:"resourceType" type:"string" json:"resourceType,omitempty"`
	Tags                         *map[string]*string `locationName:"tags" type:"map" json:"tags,omitempty"`
	Version                      *string             `locationName:"version" type:"string" json:"version,omitempty"`

	metadataConfigurationItem `json:"-", xml:"-"`
}

type metadataConfigurationItem struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ConfigurationRecorder struct {
	Name    *string `locationName:"name" type:"string" json:"name,omitempty"`
	RoleARN *string `locationName:"roleARN" type:"string" json:"roleARN,omitempty"`

	metadataConfigurationRecorder `json:"-", xml:"-"`
}

type metadataConfigurationRecorder struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ConfigurationRecorderStatus struct {
	LastErrorCode        *string    `locationName:"lastErrorCode" type:"string" json:"lastErrorCode,omitempty"`
	LastErrorMessage     *string    `locationName:"lastErrorMessage" type:"string" json:"lastErrorMessage,omitempty"`
	LastStartTime        *time.Time `locationName:"lastStartTime" type:"timestamp" timestampFormat:"unix" json:"lastStartTime,omitempty"`
	LastStatus           *string    `locationName:"lastStatus" type:"string" json:"lastStatus,omitempty"`
	LastStatusChangeTime *time.Time `locationName:"lastStatusChangeTime" type:"timestamp" timestampFormat:"unix" json:"lastStatusChangeTime,omitempty"`
	LastStopTime         *time.Time `locationName:"lastStopTime" type:"timestamp" timestampFormat:"unix" json:"lastStopTime,omitempty"`
	Name                 *string    `locationName:"name" type:"string" json:"name,omitempty"`
	Recording            *bool      `locationName:"recording" type:"boolean" json:"recording,omitempty"`

	metadataConfigurationRecorderStatus `json:"-", xml:"-"`
}

type metadataConfigurationRecorderStatus struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteDeliveryChannelInput struct {
	DeliveryChannelName *string `type:"string" json:",omitempty"`

	metadataDeleteDeliveryChannelInput `json:"-", xml:"-"`
}

type metadataDeleteDeliveryChannelInput struct {
	SDKShapeTraits bool `type:"structure" required:"DeliveryChannelName" json:",omitempty"`
}

type DeliverConfigSnapshotInput struct {
	DeliveryChannelName *string `locationName:"deliveryChannelName" type:"string" json:"deliveryChannelName,omitempty"`

	metadataDeliverConfigSnapshotInput `json:"-", xml:"-"`
}

type metadataDeliverConfigSnapshotInput struct {
	SDKShapeTraits bool `type:"structure" required:"deliveryChannelName" json:",omitempty"`
}

type DeliverConfigSnapshotOutput struct {
	ConfigSnapshotID *string `locationName:"configSnapshotId" type:"string" json:"configSnapshotId,omitempty"`

	metadataDeliverConfigSnapshotOutput `json:"-", xml:"-"`
}

type metadataDeliverConfigSnapshotOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeliveryChannel struct {
	Name         *string `locationName:"name" type:"string" json:"name,omitempty"`
	S3BucketName *string `locationName:"s3BucketName" type:"string" json:"s3BucketName,omitempty"`
	S3KeyPrefix  *string `locationName:"s3KeyPrefix" type:"string" json:"s3KeyPrefix,omitempty"`
	SNSTopicARN  *string `locationName:"snsTopicARN" type:"string" json:"snsTopicARN,omitempty"`

	metadataDeliveryChannel `json:"-", xml:"-"`
}

type metadataDeliveryChannel struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeliveryChannelStatus struct {
	ConfigHistoryDeliveryInfo  *ConfigExportDeliveryInfo `locationName:"configHistoryDeliveryInfo" type:"structure" json:"configHistoryDeliveryInfo,omitempty"`
	ConfigSnapshotDeliveryInfo *ConfigExportDeliveryInfo `locationName:"configSnapshotDeliveryInfo" type:"structure" json:"configSnapshotDeliveryInfo,omitempty"`
	ConfigStreamDeliveryInfo   *ConfigStreamDeliveryInfo `locationName:"configStreamDeliveryInfo" type:"structure" json:"configStreamDeliveryInfo,omitempty"`
	Name                       *string                   `locationName:"name" type:"string" json:"name,omitempty"`

	metadataDeliveryChannelStatus `json:"-", xml:"-"`
}

type metadataDeliveryChannelStatus struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeConfigurationRecorderStatusInput struct {
	ConfigurationRecorderNames []*string `type:"list" json:",omitempty"`

	metadataDescribeConfigurationRecorderStatusInput `json:"-", xml:"-"`
}

type metadataDescribeConfigurationRecorderStatusInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeConfigurationRecorderStatusOutput struct {
	ConfigurationRecordersStatus []*ConfigurationRecorderStatus `type:"list" json:",omitempty"`

	metadataDescribeConfigurationRecorderStatusOutput `json:"-", xml:"-"`
}

type metadataDescribeConfigurationRecorderStatusOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeConfigurationRecordersInput struct {
	ConfigurationRecorderNames []*string `type:"list" json:",omitempty"`

	metadataDescribeConfigurationRecordersInput `json:"-", xml:"-"`
}

type metadataDescribeConfigurationRecordersInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeConfigurationRecordersOutput struct {
	ConfigurationRecorders []*ConfigurationRecorder `type:"list" json:",omitempty"`

	metadataDescribeConfigurationRecordersOutput `json:"-", xml:"-"`
}

type metadataDescribeConfigurationRecordersOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDeliveryChannelStatusInput struct {
	DeliveryChannelNames []*string `type:"list" json:",omitempty"`

	metadataDescribeDeliveryChannelStatusInput `json:"-", xml:"-"`
}

type metadataDescribeDeliveryChannelStatusInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDeliveryChannelStatusOutput struct {
	DeliveryChannelsStatus []*DeliveryChannelStatus `type:"list" json:",omitempty"`

	metadataDescribeDeliveryChannelStatusOutput `json:"-", xml:"-"`
}

type metadataDescribeDeliveryChannelStatusOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDeliveryChannelsInput struct {
	DeliveryChannelNames []*string `type:"list" json:",omitempty"`

	metadataDescribeDeliveryChannelsInput `json:"-", xml:"-"`
}

type metadataDescribeDeliveryChannelsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeDeliveryChannelsOutput struct {
	DeliveryChannels []*DeliveryChannel `type:"list" json:",omitempty"`

	metadataDescribeDeliveryChannelsOutput `json:"-", xml:"-"`
}

type metadataDescribeDeliveryChannelsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetResourceConfigHistoryInput struct {
	ChronologicalOrder *string    `locationName:"chronologicalOrder" type:"string" json:"chronologicalOrder,omitempty"`
	EarlierTime        *time.Time `locationName:"earlierTime" type:"timestamp" timestampFormat:"unix" json:"earlierTime,omitempty"`
	LaterTime          *time.Time `locationName:"laterTime" type:"timestamp" timestampFormat:"unix" json:"laterTime,omitempty"`
	Limit              *int       `locationName:"limit" type:"integer" json:"limit,omitempty"`
	NextToken          *string    `locationName:"nextToken" type:"string" json:"nextToken,omitempty"`
	ResourceID         *string    `locationName:"resourceId" type:"string" json:"resourceId,omitempty"`
	ResourceType       *string    `locationName:"resourceType" type:"string" json:"resourceType,omitempty"`

	metadataGetResourceConfigHistoryInput `json:"-", xml:"-"`
}

type metadataGetResourceConfigHistoryInput struct {
	SDKShapeTraits bool `type:"structure" required:"resourceType,resourceId" json:",omitempty"`
}

type GetResourceConfigHistoryOutput struct {
	ConfigurationItems []*ConfigurationItem `locationName:"configurationItems" type:"list" json:"configurationItems,omitempty"`
	NextToken          *string              `locationName:"nextToken" type:"string" json:"nextToken,omitempty"`

	metadataGetResourceConfigHistoryOutput `json:"-", xml:"-"`
}

type metadataGetResourceConfigHistoryOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InsufficientDeliveryPolicyException struct {
	metadataInsufficientDeliveryPolicyException `json:"-", xml:"-"`
}

type metadataInsufficientDeliveryPolicyException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidConfigurationRecorderNameException struct {
	metadataInvalidConfigurationRecorderNameException `json:"-", xml:"-"`
}

type metadataInvalidConfigurationRecorderNameException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidDeliveryChannelNameException struct {
	metadataInvalidDeliveryChannelNameException `json:"-", xml:"-"`
}

type metadataInvalidDeliveryChannelNameException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidLimitException struct {
	metadataInvalidLimitException `json:"-", xml:"-"`
}

type metadataInvalidLimitException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidNextTokenException struct {
	metadataInvalidNextTokenException `json:"-", xml:"-"`
}

type metadataInvalidNextTokenException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidRoleException struct {
	metadataInvalidRoleException `json:"-", xml:"-"`
}

type metadataInvalidRoleException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidS3KeyPrefixException struct {
	metadataInvalidS3KeyPrefixException `json:"-", xml:"-"`
}

type metadataInvalidS3KeyPrefixException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidSNSTopicARNException struct {
	metadataInvalidSNSTopicARNException `json:"-", xml:"-"`
}

type metadataInvalidSNSTopicARNException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvalidTimeRangeException struct {
	metadataInvalidTimeRangeException `json:"-", xml:"-"`
}

type metadataInvalidTimeRangeException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type LastDeliveryChannelDeleteFailedException struct {
	metadataLastDeliveryChannelDeleteFailedException `json:"-", xml:"-"`
}

type metadataLastDeliveryChannelDeleteFailedException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type MaxNumberOfConfigurationRecordersExceededException struct {
	metadataMaxNumberOfConfigurationRecordersExceededException `json:"-", xml:"-"`
}

type metadataMaxNumberOfConfigurationRecordersExceededException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type MaxNumberOfDeliveryChannelsExceededException struct {
	metadataMaxNumberOfDeliveryChannelsExceededException `json:"-", xml:"-"`
}

type metadataMaxNumberOfDeliveryChannelsExceededException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoAvailableConfigurationRecorderException struct {
	metadataNoAvailableConfigurationRecorderException `json:"-", xml:"-"`
}

type metadataNoAvailableConfigurationRecorderException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoAvailableDeliveryChannelException struct {
	metadataNoAvailableDeliveryChannelException `json:"-", xml:"-"`
}

type metadataNoAvailableDeliveryChannelException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoRunningConfigurationRecorderException struct {
	metadataNoRunningConfigurationRecorderException `json:"-", xml:"-"`
}

type metadataNoRunningConfigurationRecorderException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoSuchBucketException struct {
	metadataNoSuchBucketException `json:"-", xml:"-"`
}

type metadataNoSuchBucketException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoSuchConfigurationRecorderException struct {
	metadataNoSuchConfigurationRecorderException `json:"-", xml:"-"`
}

type metadataNoSuchConfigurationRecorderException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type NoSuchDeliveryChannelException struct {
	metadataNoSuchDeliveryChannelException `json:"-", xml:"-"`
}

type metadataNoSuchDeliveryChannelException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type PutConfigurationRecorderInput struct {
	ConfigurationRecorder *ConfigurationRecorder `type:"structure" json:",omitempty"`

	metadataPutConfigurationRecorderInput `json:"-", xml:"-"`
}

type metadataPutConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure" required:"ConfigurationRecorder" json:",omitempty"`
}

type PutDeliveryChannelInput struct {
	DeliveryChannel *DeliveryChannel `type:"structure" json:",omitempty"`

	metadataPutDeliveryChannelInput `json:"-", xml:"-"`
}

type metadataPutDeliveryChannelInput struct {
	SDKShapeTraits bool `type:"structure" required:"DeliveryChannel" json:",omitempty"`
}

type Relationship struct {
	RelationshipName *string `locationName:"relationshipName" type:"string" json:"relationshipName,omitempty"`
	ResourceID       *string `locationName:"resourceId" type:"string" json:"resourceId,omitempty"`
	ResourceType     *string `locationName:"resourceType" type:"string" json:"resourceType,omitempty"`

	metadataRelationship `json:"-", xml:"-"`
}

type metadataRelationship struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ResourceNotDiscoveredException struct {
	metadataResourceNotDiscoveredException `json:"-", xml:"-"`
}

type metadataResourceNotDiscoveredException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type StartConfigurationRecorderInput struct {
	ConfigurationRecorderName *string `type:"string" json:",omitempty"`

	metadataStartConfigurationRecorderInput `json:"-", xml:"-"`
}

type metadataStartConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure" required:"ConfigurationRecorderName" json:",omitempty"`
}

type StopConfigurationRecorderInput struct {
	ConfigurationRecorderName *string `type:"string" json:",omitempty"`

	metadataStopConfigurationRecorderInput `json:"-", xml:"-"`
}

type metadataStopConfigurationRecorderInput struct {
	SDKShapeTraits bool `type:"structure" required:"ConfigurationRecorderName" json:",omitempty"`
}

type ValidationException struct {
	metadataValidationException `json:"-", xml:"-"`
}

type metadataValidationException struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}