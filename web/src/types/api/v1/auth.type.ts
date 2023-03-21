import { UserGetData } from './user.type'

export interface AuthRes {
  token: string
  user: UserGetData
}
