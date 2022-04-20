// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An error returned by the BatchPutProperty action.
type BatchPutPropertyError struct {

	// An object that contains information about errors returned by the
	// BatchPutProperty action.
	//
	// This member is required.
	Entry *PropertyValueEntry

	// The error code.
	//
	// This member is required.
	ErrorCode *string

	// The error message.
	//
	// This member is required.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// An object that contains information about errors returned by the
// BatchPutProperty action.
type BatchPutPropertyErrorEntry struct {

	// A list of objects that contain information about errors returned by the
	// BatchPutProperty action.
	//
	// This member is required.
	Errors []BatchPutPropertyError

	noSmithyDocumentSerde
}

// An object that sets information about a component type create or update request.
type ComponentRequest struct {

	// The ID of the component type.
	ComponentTypeId *string

	// The description of the component request.
	Description *string

	// An object that maps strings to the properties to set in the component type. Each
	// string in the mapping must be unique to this object.
	Properties map[string]PropertyRequest

	noSmithyDocumentSerde
}

// An object that returns information about a component type create or update
// request.
type ComponentResponse struct {

	// The name of the component.
	ComponentName *string

	// The ID of the component type.
	ComponentTypeId *string

	// The name of the property definition set in the request.
	DefinedIn *string

	// The description of the component type.
	Description *string

	// An object that maps strings to the properties to set in the component type. Each
	// string in the mapping must be unique to this object.
	Properties map[string]PropertyResponse

	// The status of the component type.
	Status *Status

	noSmithyDocumentSerde
}

// An object that contains information about a component type.
type ComponentTypeSummary struct {

	// The ARN of the component type.
	//
	// This member is required.
	Arn *string

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The date and time when the component type was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The date and time when the component type was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The description of the component type.
	Description *string

	// The current status of the component type.
	Status *Status

	noSmithyDocumentSerde
}

// The component update request.
type ComponentUpdateRequest struct {

	// The ID of the component type.
	ComponentTypeId *string

	// The description of the component type.
	Description *string

	// An object that maps strings to the properties to set in the component type
	// update. Each string in the mapping must be unique to this object.
	PropertyUpdates map[string]PropertyRequest

	// The update type of the component update request.
	UpdateType ComponentUpdateType

	noSmithyDocumentSerde
}

// The data connector.
type DataConnector struct {

	// A Boolean value that specifies whether the data connector is native to IoT
	// TwinMaker.
	IsNative *bool

	// The Lambda function associated with this data connector.
	Lambda *LambdaFunction

	noSmithyDocumentSerde
}

// An object that specifies the data type of a property.
type DataType struct {

	// The underlying type of the data type.
	//
	// This member is required.
	Type Type

	// The allowed values for this data type.
	AllowedValues []DataValue

	// The nested type in the data type.
	NestedType *DataType

	// A relationship that associates a component with another component.
	Relationship *Relationship

	// The unit of measure used in this data type.
	UnitOfMeasure *string

	noSmithyDocumentSerde
}

// An object that specifies a value for a property.
type DataValue struct {

	// A Boolean value.
	BooleanValue *bool

	// A double value.
	DoubleValue *float64

	// An expression that produces the value.
	Expression *string

	// An integer value.
	IntegerValue *int32

	// A list of multiple values.
	ListValue []DataValue

	// A long value.
	LongValue *int64

	// An object that maps strings to multiple DataValue objects.
	MapValue map[string]DataValue

	// A value that relates a component to another component.
	RelationshipValue *RelationshipValue

	// A string value.
	StringValue *string

	noSmithyDocumentSerde
}

// An object that uniquely identifies an entity property.
type EntityPropertyReference struct {

	// The name of the property.
	//
	// This member is required.
	PropertyName *string

	// The name of the component.
	ComponentName *string

	// The ID of the entity.
	EntityId *string

	// A mapping of external IDs to property names. External IDs uniquely identify
	// properties from external data stores.
	ExternalIdProperty map[string]string

	noSmithyDocumentSerde
}

// An object that contains information about an entity.
type EntitySummary struct {

	// The ARN of the entity.
	//
	// This member is required.
	Arn *string

	// The date and time when the entity was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The ID of the entity.
	//
	// This member is required.
	EntityId *string

	// The name of the entity.
	//
	// This member is required.
	EntityName *string

	// The current status of the entity.
	//
	// This member is required.
	Status *Status

	// The last date and time when the entity was updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The description of the entity.
	Description *string

	// A Boolean value that specifies whether the entity has child entities or not.
	HasChildEntities *bool

	// The ID of the parent entity.
	ParentEntityId *string

	noSmithyDocumentSerde
}

// The error details.
type ErrorDetails struct {

	// The error code.
	Code ErrorCode

	// The error message.
	Message *string

	noSmithyDocumentSerde
}

// The function request body.
type FunctionRequest struct {

	// The data connector.
	ImplementedBy *DataConnector

	// The required properties of the function.
	RequiredProperties []string

	// The scope of the function.
	Scope Scope

	noSmithyDocumentSerde
}

// The function response.
type FunctionResponse struct {

	// The data connector.
	ImplementedBy *DataConnector

	// Indicates whether this function is inherited.
	IsInherited *bool

	// The required properties of the function.
	RequiredProperties []string

	// The scope of the function.
	Scope Scope

	noSmithyDocumentSerde
}

// An object that specifies how to interpolate data in a list.
type InterpolationParameters struct {

	// The interpolation type.
	InterpolationType InterpolationType

	// The interpolation time interval in seconds.
	IntervalInSeconds *int64

	noSmithyDocumentSerde
}

// The Lambda function.
type LambdaFunction struct {

	// The ARN of the Lambda function.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

// An object that filters items in a list of component types.
//
// The following types satisfy this interface:
//  ListComponentTypesFilterMemberExtendsFrom
//  ListComponentTypesFilterMemberIsAbstract
//  ListComponentTypesFilterMemberNamespace
type ListComponentTypesFilter interface {
	isListComponentTypesFilter()
}

// The component type that the component types in the list extend.
type ListComponentTypesFilterMemberExtendsFrom struct {
	Value string

	noSmithyDocumentSerde
}

func (*ListComponentTypesFilterMemberExtendsFrom) isListComponentTypesFilter() {}

// A Boolean value that specifies whether the component types in the list are
// abstract.
type ListComponentTypesFilterMemberIsAbstract struct {
	Value bool

	noSmithyDocumentSerde
}

func (*ListComponentTypesFilterMemberIsAbstract) isListComponentTypesFilter() {}

// The namespace to which the component types in the list belong.
type ListComponentTypesFilterMemberNamespace struct {
	Value string

	noSmithyDocumentSerde
}

func (*ListComponentTypesFilterMemberNamespace) isListComponentTypesFilter() {}

// An object that filters items in a list of entities.
//
// The following types satisfy this interface:
//  ListEntitiesFilterMemberComponentTypeId
//  ListEntitiesFilterMemberExternalId
//  ListEntitiesFilterMemberParentEntityId
type ListEntitiesFilter interface {
	isListEntitiesFilter()
}

// The ID of the component type in the entities in the list.
type ListEntitiesFilterMemberComponentTypeId struct {
	Value string

	noSmithyDocumentSerde
}

func (*ListEntitiesFilterMemberComponentTypeId) isListEntitiesFilter() {}

// The external-Id property of a component. The external-Id property is the primary
// key of an external storage system.
type ListEntitiesFilterMemberExternalId struct {
	Value string

	noSmithyDocumentSerde
}

func (*ListEntitiesFilterMemberExternalId) isListEntitiesFilter() {}

// The parent of the entities in the list.
type ListEntitiesFilterMemberParentEntityId struct {
	Value string

	noSmithyDocumentSerde
}

func (*ListEntitiesFilterMemberParentEntityId) isListEntitiesFilter() {}

// The parent entity update request.
type ParentEntityUpdateRequest struct {

	// The type of the update.
	//
	// This member is required.
	UpdateType ParentEntityUpdateType

	// The ID of the parent entity.
	ParentEntityId *string

	noSmithyDocumentSerde
}

// An object that sets information about a property.
type PropertyDefinitionRequest struct {

	// A mapping that specifies configuration information about the property. Use this
	// field to specify information that you read from and write to an external source.
	Configuration map[string]string

	// An object that contains information about the data type.
	DataType *DataType

	// An object that contains the default value.
	DefaultValue *DataValue

	// A Boolean value that specifies whether the property ID comes from an external
	// data store.
	IsExternalId *bool

	// A Boolean value that specifies whether the property is required.
	IsRequiredInEntity *bool

	// A Boolean value that specifies whether the property is stored externally.
	IsStoredExternally *bool

	// A Boolean value that specifies whether the property consists of time series
	// data.
	IsTimeSeries *bool

	noSmithyDocumentSerde
}

// An object that contains response data from a property definition request.
type PropertyDefinitionResponse struct {

	// An object that contains information about the data type.
	//
	// This member is required.
	DataType *DataType

	// A Boolean value that specifies whether the property ID comes from an external
	// data store.
	//
	// This member is required.
	IsExternalId *bool

	// A Boolean value that specifies whether the property definition can be updated.
	//
	// This member is required.
	IsFinal *bool

	// A Boolean value that specifies whether the property definition is imported from
	// an external data store.
	//
	// This member is required.
	IsImported *bool

	// A Boolean value that specifies whether the property definition is inherited from
	// a parent entity.
	//
	// This member is required.
	IsInherited *bool

	// A Boolean value that specifies whether the property is required in an entity.
	//
	// This member is required.
	IsRequiredInEntity *bool

	// A Boolean value that specifies whether the property is stored externally.
	//
	// This member is required.
	IsStoredExternally *bool

	// A Boolean value that specifies whether the property consists of time series
	// data.
	//
	// This member is required.
	IsTimeSeries *bool

	// A mapping that specifies configuration information about the property.
	Configuration map[string]string

	// An object that contains the default value.
	DefaultValue *DataValue

	noSmithyDocumentSerde
}

// An object that filters items returned by a property request.
type PropertyFilter struct {

	// The operator associated with this property filter.
	Operator *string

	// The property name associated with this property filter.
	PropertyName *string

	// The value associated with this property filter.
	Value *DataValue

	noSmithyDocumentSerde
}

// The latest value of the property.
type PropertyLatestValue struct {

	// An object that specifies information about a property.>
	//
	// This member is required.
	PropertyReference *EntityPropertyReference

	// The value of the property.
	PropertyValue *DataValue

	noSmithyDocumentSerde
}

// An object that sets information about a property.
type PropertyRequest struct {

	// An object that specifies information about a property.
	Definition *PropertyDefinitionRequest

	// The update type of the update property request.
	UpdateType PropertyUpdateType

	// The value of the property.
	Value *DataValue

	noSmithyDocumentSerde
}

// An object that contains information about a property response.
type PropertyResponse struct {

	// An object that specifies information about a property.
	Definition *PropertyDefinitionResponse

	// The value of the property.
	Value *DataValue

	noSmithyDocumentSerde
}

// An object that contains information about a value for a time series property.
type PropertyValue struct {

	// An object that specifies a value for a time series property.
	//
	// This member is required.
	Value *DataValue

	// Timestamp represented in ISO 8601 format
	Time *string

	// The timestamp of a value for a time series property.
	//
	// Deprecated: This field is deprecated and will throw an error in the future. Use
	// time instead.
	Timestamp *time.Time

	noSmithyDocumentSerde
}

// An object that specifies information about time series property values.
type PropertyValueEntry struct {

	// An object that contains information about the entity that has the property.
	//
	// This member is required.
	EntityPropertyReference *EntityPropertyReference

	// A list of objects that specify time series property values.
	PropertyValues []PropertyValue

	noSmithyDocumentSerde
}

// The history of values for a time series property.
type PropertyValueHistory struct {

	// An object that uniquely identifies an entity property.
	//
	// This member is required.
	EntityPropertyReference *EntityPropertyReference

	// A list of objects that contain information about the values in the history of a
	// time series property.
	Values []PropertyValue

	noSmithyDocumentSerde
}

// An object that specifies a relationship with another component type.
type Relationship struct {

	// The type of the relationship.
	RelationshipType *string

	// The ID of the target component type associated with this relationship.
	TargetComponentTypeId *string

	noSmithyDocumentSerde
}

// A value that associates a component and an entity.
type RelationshipValue struct {

	// The name of the target component associated with the relationship value.
	TargetComponentName *string

	// The ID of the target entity associated with this relationship value.
	TargetEntityId *string

	noSmithyDocumentSerde
}

// An object that contains information about a scene.
type SceneSummary struct {

	// The ARN of the scene.
	//
	// This member is required.
	Arn *string

	// The relative path that specifies the location of the content definition file.
	//
	// This member is required.
	ContentLocation *string

	// The date and time when the scene was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The ID of the scene.
	//
	// This member is required.
	SceneId *string

	// The date and time when the scene was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The scene description.
	Description *string

	noSmithyDocumentSerde
}

// An object that represents the status of an entity, component, component type, or
// workspace.
type Status struct {

	// The error message.
	Error *ErrorDetails

	// The current state of the entity, component, component type, or workspace.
	State State

	noSmithyDocumentSerde
}

// An object that contains information about a workspace.
type WorkspaceSummary struct {

	// The ARN of the workspace.
	//
	// This member is required.
	Arn *string

	// The date and time when the workspace was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The date and time when the workspace was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// The description of the workspace.
	Description *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isListComponentTypesFilter() {}
func (*UnknownUnionMember) isListEntitiesFilter()       {}