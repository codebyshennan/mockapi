import {
  Box,
  Card,
  CircularProgress,
  Dialog,
  DialogContent,
  DialogTitle,
  Grid
} from '@mui/material'
import { CredentialResponse, GoogleLogin } from '@react-oauth/google'
import { useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { toast } from 'react-toastify'
import AuthApi from '../../api/AuthApi'
import { useAuth } from '../../context/AuthContext'

function LoginPage() {
  const loaded = useRef()
  const [openLoadingDialog, setOpenLoadingDialog] = useState(false)

  const navigate = useNavigate()
  const { setToken } = useAuth()

  const onGoogleLoginError = () => {
    toast.error('An unexpected error has occured, please try again later')
  }

  const onGoogleLoginSuccess = async (response: CredentialResponse) => {
    setOpenLoadingDialog(true)
    AuthApi.googleLogin(response)
      .then((res) => {
        setToken(res.data.token)
      })
      .catch((err) => {
        console.log({ err })
        onGoogleLoginError()
      })
      .finally(() => {
        setOpenLoadingDialog(false)
      })
  }

  return (
    <Box
      className='background'
      sx={{ width: '100vw', height: '100vh', overflowX: 'hidden' }}
    >
      <Grid
        container
        direction='row'
        justifyContent='space-evenly'
        alignItems='center'
        style={{ height: '100vh', margin: '0 auto', padding: '1em' }}
      >
        <Grid item xs md={6}>
          <Box
            style={{
              color: '#FFFFFF',
              fontStyle: 'roboto',
              fontWeight: 900,
              fontSize: '5em',
              margin: '0 auto'
            }}
          >
            API Sandbox
            <br />
            <Box style={{ fontSize: '50%' }}>Mock API Server</Box>
          </Box>
        </Grid>
        <Grid item xs={12} md={6}>
          <Card
            style={{
              padding: '36px 2rem',
              maxWidth: 'fit-content',
              margin: '0 auto'
            }}
          >
            <Grid container direction='column' spacing={2}>
              <Grid item>
                <GoogleLogin
                  onSuccess={onGoogleLoginSuccess}
                  onError={onGoogleLoginError}
                />
              </Grid>
            </Grid>
          </Card>
        </Grid>
      </Grid>
      {/* Loading spinner dialog */}
      <Dialog
        open={openLoadingDialog}
        onClose={() => setOpenLoadingDialog(false)}
        aria-labelledby='alert-dialog-title'
        aria-describedby='alert-dialog-description'
      >
        <DialogTitle id='alert-dialog-title'>Please wait..</DialogTitle>
        <DialogContent style={{ textAlign: 'center', marginBottom: 20 }}>
          <CircularProgress />
        </DialogContent>
      </Dialog>
    </Box>
  )
}

export default LoginPage
