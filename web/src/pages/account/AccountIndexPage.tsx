import React, { SyntheticEvent, useCallback } from 'react'
import { useAuth, useOptionalAuth } from '../../context/AuthContext'
import Layout from '../../layout/Layout'
import AuthApi from '../../api/AuthApi'
import { Box, Button } from '@mui/material'
import { useNavigate } from 'react-router-dom'
import jwt_decode from 'jwt-decode'
import format from 'date-fns/format'
import formatDistance from 'date-fns/formatDistance'

function onLogout(e: SyntheticEvent) {
  e.preventDefault()
  AuthApi.logout()
  window.location.reload()
}

function getExpDateOfJwt(token: string): string {
  const { exp } = jwt_decode<{ exp: number }>(token)
  const date = new Date(0)
  date.setUTCSeconds(exp)
  return `${format(date, 'yyyy/MM/dd HH:mm:ss')} (${formatDistance(
    date,
    new Date(),
    {
      addSuffix: true
    }
  )})`
}

export default function AccountIndexPage() {
  const {
    authRes: { token, user }
  } = useOptionalAuth()

  if (!user || !token) {
    return <Layout title='My Account'>Loading...</Layout>
  }

  return (
    <Layout title='My Account'>
      <Box
        display='flex'
        flexDirection='column'
        rowGap={2}
        alignItems='flex-start'
      >
        <Box>
          Your authentication token is:
          <Box
            sx={{
              wordBreak: 'break-word',
              fontFamily: ['Monaco', 'monospace']
            }}
          >
            {token}
          </Box>
        </Box>
        <Box>
          Your authentication token expires on: {getExpDateOfJwt(token)} {}
        </Box>
        <Box>You are logged in using: {user.email}</Box>
        <Button variant='contained' onClick={onLogout}>
          Logout
        </Button>
      </Box>
    </Layout>
  )
}
