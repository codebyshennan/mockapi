export interface MockServerCreateData {
  name: string
  responses: MockEndPtCreateData[]
}

export interface MockServerGetData {
  id: string
  name: string
  ownerRef: string
  createdAt: string
  updatedAt: string
  endpts: MockEndPtGetData[]
}

export interface MockServerPatchData {
  name?: string
  responses: {
    create: MockEndPtCreateData[]
    update: MockEndPtUpdateData[]
  }
}

export interface MockEndPtCreateData {
  method: 'GET' | 'POST' | 'PATCH' | 'PUT' | 'DELETE'
  endptRegex: string
  resCode: number
  resBody: string | null,
  timeout: number | null,
  writes: {
    data: string
    dest: string
  } | null
}

export interface MockEndPtGetData {
  id: string
  method: 'GET' | 'POST' | 'PATCH' | 'PUT' | 'DELETE'
  endptRegex: string
  resCode: number
  resBody: string | null,
  timeout: number | null,
  writes: {
    data: string,
    dest: string
  } | null
}

export interface MockEndPtUpdateData extends MockEndPtCreateData {
  id: string
}
