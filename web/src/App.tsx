import CONFIG from './env'
import { GoogleOAuthProvider } from '@react-oauth/google'
import UnauthenticatedApp from './apps/UnauthenticatedApp'
import { useOptionalAuth } from './context/AuthContext'
import AuthenticatedApp from './apps/AuthenticatedApp'
import circlesTheme from './context/CirclesPaletteContext'
import { ThemeProvider } from '@mui/material'

function App() {
  const {
    authRes: { token }
  } = useOptionalAuth()

  if (!token) {
    return (
      <GoogleOAuthProvider clientId={CONFIG.GOOGLE_CLIENT_ID}>
        <UnauthenticatedApp />
      </GoogleOAuthProvider>
    )
  }

  return (
    <ThemeProvider theme={circlesTheme}>
      <AuthenticatedApp />
    </ThemeProvider>
  )
}

export default App
