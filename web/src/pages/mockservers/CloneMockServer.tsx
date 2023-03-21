import {
  SelectChangeEvent,
  Box,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button
} from '@mui/material'
import React, { useEffect, useState } from 'react'
import { toast } from 'react-toastify'
import MockServerApi from '../../api/MockServerApi'
import { useAuth } from '../../context/AuthContext'
import { MockServerGetData } from '../../types/api/v1/mockServer.type'

interface CloneMockServerProps {
  onConfirm: (d: MockServerGetData) => void
}

export default function CloneMockServer(props: CloneMockServerProps) {
  const { onConfirm } = props

  const [mockServers, setMockServers] = useState<MockServerGetData[]>([])
  const [selected, setSelected] = useState<string>('')
  const [selectedMockServer, setMockServer] = useState<
    MockServerGetData | undefined
  >()

  const {
    authRes: { token }
  } = useAuth()

  useEffect(() => {
    MockServerApi.getServers(token)
      .then((res) => setMockServers(res.data))
      .catch((err) => toast.error(err))
  }, [])

  const handleSelect = (event: SelectChangeEvent) => {
    setSelected(event.target.value)

    const swagger = mockServers.find((ele) => ele.id === event.target.value)

    if (!swagger) return
    setMockServer(swagger)
  }

  return (
    <Box width={500} display='flex' flexDirection='column' rowGap={2}>
      <FormControl fullWidth>
        <InputLabel id='clone-mockservers-select'>
          Select Mock Server
        </InputLabel>
        <Select
          labelId='clone-mockservers-select'
          id='clone-mockservers-select'
          value={selected}
          onChange={handleSelect}
          label='Select Swagger to import'
        >
          <MenuItem value='' sx={{ display: 'none' }}>
            {null}
          </MenuItem>

          {mockServers.map((element) => (
            <MenuItem value={element.id}>
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
        disabled={!selectedMockServer}
        fullWidth
        variant='contained'
        onClick={() => onConfirm(selectedMockServer!)}
      >
        Import endpoints from Swagger
      </Button>
    </Box>
  )
}
