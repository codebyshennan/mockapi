import React, {
  createContext,
  ReactNode,
  useCallback,
  useContext,
  useEffect,
  useState
} from 'react'
import Cookies from 'js-cookie'
import {
  AuthContextInterface,
  OptionalAuthContextInterface
} from '../types/context/AuthContextInterface'
import AuthApi from '../api/AuthApi'
import { UserGetData } from '../types/api/v1/user.type'
import UserApi from '../api/UserApi'

const OptionalAuthContext = createContext<OptionalAuthContextInterface | null>(
  null
)

export function OptionalAuthContextProvider(props: { children: ReactNode }) {
  const [token, _setToken] = useState<string | undefined>(() => {
    const token = Cookies.get('csb_t')

    if (!token) {
      return undefined
    }

    return token
  })

  const [user, setUser] = useState<UserGetData | undefined>()
  useEffect(() => {
    if (!token) return

    UserApi.getSelf(token)
      .then((res) => setUser(res.data))
      .catch((_err) => {
        // do nothing
      })
  }, [token])

  const setToken = useCallback((d: string) => {
    Cookies.set('csb_t', d)
    _setToken(d)
  }, [])

  return (
    <OptionalAuthContext.Provider
      value={{
        authRes: {
          token,
          user
        },
        setToken
      }}
      {...props}
    />
  )
}

export function useOptionalAuth() {
  const ctx = useContext(OptionalAuthContext)
  if (!ctx) {
    throw new Error(
      'useOptionalAuth has to be used in OptionalAuthContextProvider'
    )
  }
  return ctx
}

/**
 * A wrapper around OptionalAuthContext to prevent the need
 * for additional optional chaining like
 * `useOptionalAuth().authRes?.`
 */
export function useAuth(): AuthContextInterface {
  const ctx = useContext(OptionalAuthContext)
  if (!ctx) {
    throw new Error('useAuth has to be used in OptionalAuthContextProvider')
  }

  return ctx as AuthContextInterface
}
