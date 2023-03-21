package models

// Represents the OpenApi 3.0 Specification Swagger object
// See https://swagger.io/specification under OpenAPI Object
// for list of fields
type OpenApiV3 struct {
	Openapi    string                      `json:"openapi" bson:"openapi"`
	Info       v3infoObject                `json:"info" bson:"info"`
	Servers    []map[string]interface{}    `json:"servers" bson:"servers"`
	Paths      map[string]v3pathItemObject `json:"paths" bson:"paths"`
	Components map[string]interface{}      `json:"components" bson:"components"`
	Security   []map[string]interface{}    `json:"security" bson:"security"`
	Tags       []map[string]interface{}    `json:"tags" bson:"tags"`
}

type v3infoObject struct {
	Title          string `json:"title" bson:"title"`
	Description    string `json:"description" bson:"description"`
	TermsOfService string `json:"termsOfService" bson:"termsOfService"`
	Version        string `json:"version" bson:"version"`
}

type v3pathItemObject struct {
	Ref    string             `json:"$ref,omitempty" bson:"swagger_ref,omitempty"`
	Get    *v3OperationObject `json:"get,omitempty" bson:"get,omitempty"`
	Post   *v3OperationObject `json:"post,omitempty" bson:"post,omitempty"`
	Put    *v3OperationObject `json:"put,omitempty" bson:"put,omitempty"`
	Patch  *v3OperationObject `json:"patch,omitempty" bson:"patch,omitempty"`
	Delete *v3OperationObject `json:"delete,omitempty" bson:"delete,omitempty"`
}

func (o v3pathItemObject) AsMap() map[string]v3OperationObject {
	result := make(map[string]v3OperationObject)

	if o.Get != nil && (*o.Get).Responses != nil {
		result["GET"] = *o.Get
	}

	if o.Post != nil && (*o.Post).Responses != nil {
		result["POST"] = *o.Post
	}

	if o.Put != nil && (*o.Put).Responses != nil {
		result["PUT"] = *o.Put
	}

	if o.Patch != nil && (*o.Patch).Responses != nil {
		result["PATCH"] = *o.Patch
	}

	if o.Delete != nil && (*o.Delete).Responses != nil {
		result["DELETE"] = *o.Delete
	}

	return result
}

type v3OperationObject struct {
	RequestBody map[string]any `json:"requestBody,omitempty" bson:"requestBody,omitempty"`
	Responses   map[string]any `json:"responses,omitempty" bson:"responses,omitempty"`
}
