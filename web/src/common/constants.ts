export function getAuthHeader(token: string) {
  return {
    Authorization: `Bearer ${token}`
  }
}

export const COOKIE_NAME = 'csb_t'
