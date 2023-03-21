import { Box, Paper, Typography } from '@mui/material'
import format from 'date-fns/format'
import parseJSON from 'date-fns/parseJSON'
import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import SwaggerUI from 'swagger-ui-react'
import ServiceApi from '../../../api/ServiceApi'
import { useAuth } from '../../../context/AuthContext'
import Layout from '../../../layout/Layout'
import { ServiceGetData } from '../../../types/api/v1/service.type'
import 'swagger-ui-react/swagger-ui.css'
import SwaggerApi from '../../../api/SwaggerApi'
import SwaggerUIWrapper from './SwaggerUIWrapper'
import NoSwaggerRef from './NoSwaggerRef'
import { SwaggerGetData } from '../../../types/api/v1/swagger.type'

export default function ServicesViewPage() {
  const { id } = useParams()
  const {
    authRes: { token }
  } = useAuth()

  const [svc, setSvc] = useState<ServiceGetData | undefined>()
  useEffect(() => {
    if (!id || !token) return

    ServiceApi.getServiceById(id, token).then((res) => setSvc(res.data))
  }, [])

  const [swaggerModel, setSwaggerModel] = useState<SwaggerGetData | undefined>()
  useEffect(() => {
    if (!id || !token || !svc || svc.swaggerRefs.length === 0) return

    SwaggerApi.getSwaggerById(svc.swaggerRefs[0], token).then((res) =>
      setSwaggerModel(res.data)
    )
  }, [svc])

  if (!id) {
    return null
  }

  return (
    <Layout title='View Service'>
      <Box display='flex' flexDirection='column' rowGap={2}>
        <>
          {svc && (
            <Paper sx={{ p: 2 }}>
              <Typography variant='body1'>
                API Service Name: {svc.name}
              </Typography>
              <Typography variant='body1'>
                Created On: {format(parseJSON(svc.createdAt), 'dd/MM/yyyy')}
              </Typography>
              {svc.desc && <p>{svc.desc}</p>}
            </Paper>
          )}

          {svc && !swaggerModel && <NoSwaggerRef serviceId={id} />}
          {svc && swaggerModel && (
            <SwaggerUIWrapper swagger={swaggerModel.swaggerSpec as object} />
          )}
        </>
      </Box>
    </Layout>
  )
}
