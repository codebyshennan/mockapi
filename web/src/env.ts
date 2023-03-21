const CONFIG = {
  GOOGLE_CLIENT_ID: import.meta.env.VITE_GOOGLE_CLIENT_ID,
  NODE_ENV: import.meta.env.DEV ? 'development' : 'production',
  SB_API: import.meta.env.VITE_SB_API
}

export default CONFIG
