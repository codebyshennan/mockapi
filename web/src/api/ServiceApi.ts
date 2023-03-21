import { AxiosPromise } from 'axios'
import { getAuthHeader } from '../common/constants'
import CONFIG from '../env'
import { ServiceCreateData, ServiceGetData } from '../types/api/v1/service.type'
import { EmptyObj } from '../types/util.type'
// import { UserGetData } from '../types/api/v1/user.type'
import { httpGet, httpPost } from './http'

const BASE_URL = `${CONFIG.SB_API}/api/v1/services`

class ServiceApi {
  createService(
    token: string,
    data: ServiceCreateData
  ): AxiosPromise<EmptyObj> {
    return httpPost(BASE_URL, getAuthHeader(token), data)
  }

  getServiceById(id: string, token: string): AxiosPromise<ServiceGetData> {
    return httpGet(`${BASE_URL}/${id}`, getAuthHeader(token))
  }

  getServices(token: string): AxiosPromise<ServiceGetData[]> {
    return httpGet(BASE_URL, getAuthHeader(token))
  }
}

export default new ServiceApi()
