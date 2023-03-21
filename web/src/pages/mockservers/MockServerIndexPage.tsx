import { Box, Button, Typography, Link } from '@mui/material'
import React, { useEffect, useState } from 'react'
import { NavigateFunction, useNavigate } from 'react-router-dom'
import MockServerApi from '../../api/MockServerApi'
import SbCard from '../../components/SbCard/SbCard'
import { useAuth } from '../../context/AuthContext'
import Layout from '../../layout/Layout'
import { MockServerGetData } from '../../types/api/v1/mockServer.type'

export default function MockServersIndexPage() {
  const [servers, setServers] = useState<MockServerGetData[]>([])
  const {
    authRes: { token }
  } = useAuth()

  useEffect(() => {
    MockServerApi.getServers(token).then((res) => setServers(res.data))
  }, [])

  const navigate = useNavigate()

  return (
    <Layout title='Mock Servers'>
      <Box display='flex' flexDirection='column' rowGap={1}>
        <Box display='flex' columnGap={2}>
          <Button
            sx={{
              flexBasis: '50%'
            }}
            variant='contained'
            onClick={() => navigate('/mockservers/new')}
          >
            Create a Mock Server
          </Button>
          <Button
            sx={{
              flexBasis: '50%'
            }}
            variant='outlined'
            onClick={() =>
              navigate('/mockservers/new', {
                state: {
                  useClone: true
                }
              })
            }
          >
            Clone an existing Mock Server
          </Button>
        </Box>

        {servers.map((server, idx) => (
          <SbCard
            title={
              <Link
                href={`/mockservers/${server.id}`}
                variant='subtitle1'
                sx={{
                  p: 0
                }}
              >
                {server.name}
              </Link>
            }
            key={`${server.id}-${idx}`}
          >
            Number of Mock Endpoints: {server.endpts.length}
          </SbCard>
        ))}
      </Box>
    </Layout>
  )
}
