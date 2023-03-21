export interface PathItemObject {
  get: OperationObject,
  post: OperationObject,
  delete: OperationObject,
  patch: OperationObject,
  put: OperationObject,
}

export interface OperationObject {
  requestBody?: unknown,
  responses: Record<string, {
    content: Record<string, unknown> // application-json: <spec>
  }>
}

export type PathsObject = Record<string, PathItemObject>

export interface OpenApiV3 {
  paths: PathsObject
}