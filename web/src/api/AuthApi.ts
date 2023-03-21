import { CredentialResponse } from '@react-oauth/google'
import { AxiosPromise } from 'axios'
import Cookies from 'js-cookie'
import { COOKIE_NAME } from '../common/constants'
import CONFIG from '../env'
import { AuthRes } from '../types/api/v1/auth.type'
import { EmptyObj } from '../types/util.type'
import { httpPost } from './http'

const BASE_URL = `${CONFIG.SB_API}/api/v1/auth`

class AuthApi {
  googleLogin(data: CredentialResponse): AxiosPromise<AuthRes> {
    return httpPost(`${BASE_URL}/googleLogin`, {}, data)
  }

  logout(): AxiosPromise<EmptyObj> {
    Cookies.remove(COOKIE_NAME)
    return httpPost(`${BASE_URL}/logout`, {}, {})
  }
}

export default new AuthApi()
