import { AxiosPromise } from 'axios'
import { getAuthHeader } from '../common/constants'
import CONFIG from '../env'
import { SwaggerCreateData, SwaggerGetData } from '../types/api/v1/swagger.type'
import { EmptyObj } from '../types/util.type'
import { httpGet, httpPost } from './http'

const BASE_URL = `${CONFIG.SB_API}/api/v1/swaggers`

class SwaggerApi {
  createSwagger(
    token: string,
    data: SwaggerCreateData
  ): AxiosPromise<EmptyObj> {
    return httpPost(BASE_URL, getAuthHeader(token), data)
  }

  getSwaggers(token: string): AxiosPromise<SwaggerGetData[]> {
    return httpGet(BASE_URL, getAuthHeader(token))
  }

  getSwaggerById(id: string, token: string): AxiosPromise<SwaggerGetData> {
    return httpGet(`${BASE_URL}/${id}`, getAuthHeader(token))
  }
}

export default new SwaggerApi()
