import {
  Box,
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectChangeEvent
} from '@mui/material'
import { FieldArrayRenderProps, FormikProps } from 'formik'
import React, { useCallback, useEffect, useState } from 'react'
import { toast } from 'react-toastify'
import SwaggerApi from '../../api/SwaggerApi'
import { useAuth } from '../../context/AuthContext'
import { SwaggerGetData } from '../../types/api/v1/swagger.type'
import { MockServerFormValues } from './MockServerForm'

interface ImportEndPointsProps {
  onConfirm: (d: SwaggerGetData) => void
}

export default function ImportEndPoints(props: ImportEndPointsProps) {
  const { onConfirm } = props

  const [swaggers, setSwaggers] = useState<SwaggerGetData[]>([])
  const [selected, setSelected] = useState<string>('')
  const [selectedSwagger, setSelectedSwagger] = useState<
    SwaggerGetData | undefined
  >()

  const {
    authRes: { token }
  } = useAuth()

  useEffect(() => {
    SwaggerApi.getSwaggers(token)
      .then((res) => setSwaggers(res.data))
      .catch((err) => toast.error(err))
  }, [])

  const handleSelect = (event: SelectChangeEvent) => {
    setSelected(event.target.value)

    const swagger = swaggers.find((ele) => ele.id === event.target.value)

    if (!swagger) return
    setSelectedSwagger(swagger)
  }

  return (
    <Box width={500} display='flex' flexDirection='column' rowGap={2}>
      <FormControl fullWidth>
        <InputLabel id='import-endpts-select'>Select Swagger</InputLabel>
        <Select
          labelId='import-endpts-select'
          id='import-endpts-select'
          value={selected}
          onChange={handleSelect}
          label='Select Swagger to import'
        >
          <MenuItem value='' sx={{ display: 'none' }}>
            {null}
          </MenuItem>

          {swaggers.map((element) => (
            <MenuItem value={element.id} key={element.id}>
              {/* Read the swagger spec for swagger title 
                  @ts-ignore */}
              {element.swaggerSpec?.info?.title}
              {`  (${element.endpts?.length} endpoints)`}
            </MenuItem>
          ))}
        </Select>
      </FormControl>

      <Button
        type='button'
        disabled={!selectedSwagger}
        fullWidth
        variant='contained'
        onClick={() => onConfirm(selectedSwagger!)}
      >
        Import endpoints from Swagger
      </Button>
    </Box>
  )
}
