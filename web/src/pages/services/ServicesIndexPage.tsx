import { Box, Button, Link } from '@mui/material'
import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import ServiceApi from '../../api/ServiceApi'
import SbCard from '../../components/SbCard/SbCard'
import { useAuth } from '../../context/AuthContext'
import Layout from '../../layout/Layout'
import { ServiceGetData } from '../../types/api/v1/service.type'

export default function ServicesIndexPage() {
  const [svcs, setSvcs] = useState<ServiceGetData[]>([])
  const {
    authRes: { token }
  } = useAuth()
  const navigate = useNavigate()

  useEffect(() => {
    ServiceApi.getServices(token).then((res) => setSvcs(res.data))
  }, [])

  return (
    <Layout title='Available API Services'>
      <Box display='flex' flexDirection='column' rowGap={1}>
        <Button
          fullWidth
          variant='contained'
          onClick={() => navigate('/services/new')}
        >
          Add API Service
        </Button>
        {svcs &&
          svcs.map((svc) => (
            <SbCard
              key={svc.id}
              title={
                <Link
                  href={`/services/${svc.id}`}
                  variant='subtitle1'
                  sx={{
                    p: 0
                  }}
                >
                  {svc.name}
                </Link>
              }
            >
              {svc.desc}
            </SbCard>
          ))}
      </Box>
    </Layout>
  )
}
