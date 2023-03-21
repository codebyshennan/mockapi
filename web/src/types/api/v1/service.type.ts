export interface ServiceCreateData {
  name: string
  desc: string
}

export interface ServiceGetData {
  id: string
  name: string
  desc: string
  swaggerRefs: string[]
  createdAt: string
  updatedAt: string
}
