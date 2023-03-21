package domain

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
	Patch  *v3OperationObject `json:"patch,omitempty" bson:"patch,omitempty"`
	Put    *v3OperationObject `json:"put,omitempty" bson:"put,omitempty"`
	Delete *v3OperationObject `json:"delete,omitempty" bson:"delete,omitempty"`
}

func (o v3pathItemObject) AsMap() map[string]v3OperationObject {
	result := make(map[string]v3OperationObject)

	if o.Get != nil {
		result["GET"] = *o.Get
	}

	if o.Post != nil {
		result["POST"] = *o.Post
	}

	if o.Patch != nil {
		result["PATCH"] = *o.Patch
	}

	if o.Put != nil {
		result["PATCH"] = *o.Put
	}

	if o.Delete != nil {
		result["DELETE"] = *o.Delete
	}

	return result
}

type v3OperationObject struct {
	RequestBody v3RequestBody `json:"requestBody,omitempty" bson:"requestBody,omitempty"`
	Responses   v3Responses   `json:"responses,omitempty" bson:"responses,omitempty"`
}

type v3Responses map[string]v3Response

type v3Response struct {
	Description string                 `json:"description,omitempty" bson:"description,omitempty"`
	Content     map[string]interface{} `json:"content,omitempty" bson:"content,omitempty"`
}

type v3RequestBody struct {
	Description string                 `json:"description,omitempty" bson:"description,omitempty"`
	Content     map[string]interface{} `json:"content,omitempty" bson:"content,omitempty"`
	Required    bool                   `json:"required,omitempty" bson:"required,omitempty"`
}

func (o OpenApiV3) GetEndPts() (res []SwaggerEndPointMetadata, e error) {
	for url, pathItemObj := range o.Paths {
		for method := range pathItemObj.AsMap() {
			res = append(res, SwaggerEndPointMetadata{
				EndpointRegex: url,
				Method:        method,
			})
		}
	}
	return
}
