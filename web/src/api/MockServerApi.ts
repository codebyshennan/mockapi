import { AxiosPromise } from 'axios'
import { getAuthHeader } from '../common/constants'
import CONFIG from '../env'
import {
  MockServerCreateData,
  MockServerGetData,
  MockServerPatchData
} from '../types/api/v1/mockServer.type'
import { ServiceCreateData, ServiceGetData } from '../types/api/v1/service.type'
import { EmptyObj, IdObj } from '../types/util.type'
import { httpGet, httpPatch, httpPost } from './http'

const BASE_URL = `${CONFIG.SB_API}/api/v1/mockServers`

class MockServerApi {
  createServer(token: string, data: MockServerCreateData): AxiosPromise<IdObj> {
    return httpPost(BASE_URL, getAuthHeader(token), data)
  }

  getServerById(id: string, token: string): AxiosPromise<MockServerGetData> {
    return httpGet(`${BASE_URL}/${id}`, getAuthHeader(token))
  }

  getServers(token: string): AxiosPromise<MockServerGetData[]> {
    return httpGet(BASE_URL, getAuthHeader(token))
  }

  updateServer(
    id: string,
    token: string,
    data: MockServerPatchData
  ): AxiosPromise<EmptyObj> {
    return httpPatch(`${BASE_URL}/${id}`, getAuthHeader(token), data)
  }
}

export default new MockServerApi()
