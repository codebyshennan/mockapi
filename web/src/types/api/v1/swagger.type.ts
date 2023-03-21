export interface SwaggerCreateData {
  version: string
  serviceRef: string
  sourceType: 'doc' | 'url'
  source: unknown
}

export interface SwaggerGetData {
  createdAt: string
  id: string
  ownerRef: string
  serviceRef: string
  swaggerSpec: unknown
  updatedAt: string
  version: string
  endpts: {
    endptRegex: string,
    method: string
  }[]
}
