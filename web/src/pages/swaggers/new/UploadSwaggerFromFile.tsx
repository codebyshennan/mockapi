import {
  Paper,
  Typography,
  TextField,
  InputAdornment,
  Button
} from '@mui/material'
import React, { useCallback, useState } from 'react'
import AttachFileIcon from '@mui/icons-material/AttachFile'
import SwaggerApi from '../../../api/SwaggerApi'
import { useAuth } from '../../../context/AuthContext'
import { useNavigate } from 'react-router-dom'
import { toast } from 'react-toastify'
import LoadingButton from '@mui/lab/LoadingButton'

export default function UploadSwaggerFromFile(props: { serviceId: string }) {
  const navigate = useNavigate()
  const [file, setFile] = useState<string>('')
  const [fileName, setFilename] = useState<string>('')
  const [loading, setLoading] = useState<boolean>(false)
  const {
    authRes: { token }
  } = useAuth()

  const onSubmitFile = useCallback(() => {
    if (!file) return

    try {
      const fileAsObj = JSON.parse(file)
      setLoading(true)
      SwaggerApi.createSwagger(token, {
        sourceType: 'doc',
        source: fileAsObj,
        version: '-', // TODO
        serviceRef: props.serviceId
      })
        .then(() => {
          toast.success('Swagger uploaded successfully')
          navigate(`/services/${props.serviceId}`)
        })
        .catch((err) => toast.error(err))
        .finally(() => setLoading(false))
    } catch (error) {
      toast.error(error as string)
    }
  }, [file])

  const onFileUpload = useCallback((files: FileList | null) => {
    if (!files || files?.length === 0) return
    const file = files[0]
    const fileName = file.name
    const reader = new FileReader()

    reader.onload = async (e) => {
      if (!e || !e.target) return

      const text = e.target.result?.toString()
      if (!text) return
      console.log(text)
      setFile(text)
      setFilename(fileName)
    }

    reader.readAsText(file)
  }, [])

  return (
    <>
      <Paper sx={{ p: 2, mb: 2 }}>
        <Typography variant='subtitle1' sx={{ mb: 2 }}>
          Upload a Swagger Doc from your Computer
        </Typography>

        {fileName && (
          <Typography variant='subtitle2'>
            Currently uploaded: {fileName}
          </Typography>
        )}
        <input
          accept='.json'
          style={{ display: 'none' }}
          id='upload-swagger'
          type='file'
          onChange={(e) => onFileUpload(e.target.files)}
        />
        <label htmlFor='upload-swagger'>
          <Button variant='contained' component='span'>
            {!fileName ? 'Upload File' : 'Change upload'}
          </Button>
        </label>
      </Paper>
      <LoadingButton
        type='button'
        variant='contained'
        color='success'
        onClick={onSubmitFile}
        loading={loading}
      >
        Upload Swagger
      </LoadingButton>
    </>
  )
}
