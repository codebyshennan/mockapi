import { AxiosPromise } from 'axios'
import { getAuthHeader } from '../common/constants'
import CONFIG from '../env'
import { UserGetData } from '../types/api/v1/user.type'
import { httpGet } from './http'

const BASE_URL = `${CONFIG.SB_API}/api/v1/users`

class UserApi {
  getUserById(id: string, token: string): AxiosPromise<UserGetData> {
    return httpGet(`${BASE_URL}/${id}`, getAuthHeader(token))
  }

  getSelf(token: string): AxiosPromise<UserGetData> {
    return httpGet(`${BASE_URL}/self`, getAuthHeader(token))
  }
}

export default new UserApi()
