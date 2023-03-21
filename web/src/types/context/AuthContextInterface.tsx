import { AuthRes } from '../api/v1/auth.type'
import { UserGetData } from '../api/v1/user.type'

export interface OptionalAuthContextInterface {
  authRes: {
    token: string | undefined
    user: UserGetData | undefined
  }
  setToken: (d: string) => void
}

export interface AuthContextInterface {
  authRes: {
    token: string
    user: UserGetData
  }
  setToken: (d: string) => void
}
