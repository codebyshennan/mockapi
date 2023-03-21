import {
  Button,
  Input,
  InputAdornment,
  Paper,
  TextField,
  Typography
} from '@mui/material'
import React, { SyntheticEvent, useCallback, useState } from 'react'
import LinkIcon from '@mui/icons-material/Link'
import SwaggerApi from '../../../api/SwaggerApi'
import { useAuth } from '../../../context/AuthContext'
import { LoadingButton } from '@mui/lab'
import { toast } from 'react-toastify'
import { useNavigate } from 'react-router-dom'
import { parseUrlContent } from './utils'

interface UploadSwaggerWithUrlProps {
  serviceId: string
}

export default function UploadSwaggerWithUrl(props: UploadSwaggerWithUrlProps) {
  const navigate = useNavigate()

  const [url, setUrl] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(false)
  const [isDirty, setIsDirty] = useState<boolean>(false)
  const {
    authRes: { token }
  } = useAuth()

  const onUrlChange = useCallback(
    (e: {
      target: {
        value: string
      }
    }) => {
      setIsDirty(true)
      setUrl(e.target.value)
    },
    []
  )

  const onSubmitUrl = async () => {
    if (!url) return

    setLoading(true)

    const spec = await parseUrlContent(url)
    if (!spec) {
      setLoading(false)
      return
    }

    console.log(spec)
    return
    SwaggerApi.createSwagger(token, {
      sourceType: 'doc',
      source: spec,
      version: '-', // TODO
      serviceRef: props.serviceId
    })
      .then(() => {
        toast.success('Swagger uploaded successfully')
        navigate(`/services/${props.serviceId}`)
      })
      .catch((err) => toast.error(err))
      .finally(() => setLoading(false))
  }

  return (
    <>
      <Paper sx={{ p: 2, mb: 2 }}>
        <Typography variant='subtitle1' sx={{ mb: 2 }}>
          Provide a link to your Swagger Doc
        </Typography>

        <TextField
          InputProps={{
            startAdornment: (
              <InputAdornment position='start'>
                <LinkIcon />
              </InputAdornment>
            )
          }}
          autoFocus
          color={isDirty && !url ? 'error' : undefined}
          helperText={isDirty && !url ? 'Field cannot be empty' : undefined}
          placeholder='Enter URL'
          value={url}
          onChange={onUrlChange}
          fullWidth
        />
      </Paper>
      <LoadingButton
        type='button'
        variant='contained'
        color='success'
        onClick={onSubmitUrl}
        loading={loading}
      >
        Upload Swagger
      </LoadingButton>
    </>
  )
}
