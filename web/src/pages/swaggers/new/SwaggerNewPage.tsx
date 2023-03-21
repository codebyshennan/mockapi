import { Box, Button } from '@mui/material'
import React, { useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import Layout from '../../../layout/Layout'
import UploadSwaggerFromFile from './UploadSwaggerFromFile'
import UploadSwaggerWithUrl from './UploadSwaggerWithUrl'

export interface PAGE_STATE {
  serviceId: string
}

export default function SwaggerNewPage() {
  const pageState = useLocation().state as PAGE_STATE
  const navigate = useNavigate()

  const [sourceType, setSourceType] = useState<'url' | 'file'>('file')

  // if this page accessed directly through the browser url
  // redirect back
  if (!pageState || !pageState.serviceId) {
    navigate('../')
    return null
  }

  const { serviceId } = pageState

  return (
    <Layout title='Upload Swagger'>
      <Box display='flex' columnGap={2} pb={2}>
        <Button
          sx={{
            flexBasis: '50%'
          }}
          variant={sourceType === 'file' ? 'contained' : 'outlined'}
          onClick={() => setSourceType('file')}
        >
          Using a file
        </Button>
        <Button
          sx={{
            flexBasis: '50%'
          }}
          disabled
          variant={sourceType === 'url' ? 'contained' : 'outlined'}
          onClick={() => setSourceType('url')}
        >
          Using a URL
        </Button>
      </Box>
      {sourceType === 'url' && <UploadSwaggerWithUrl serviceId={serviceId} />}
      {sourceType === 'file' && <UploadSwaggerFromFile serviceId={serviceId} />}
    </Layout>
  )
}
