import { Button } from '@mui/material'
import Paper from '@mui/material/Paper'
import React from 'react'
import { useNavigate } from 'react-router-dom'

export default function NoSwaggerRef(props: { serviceId: string }) {
  const navigate = useNavigate()

  return (
    <Paper
      sx={{
        p: 2,
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',
        rowGap: 1
      }}
    >
      This API service does not have a Swagger document.
      <Button
        type='button'
        onClick={() =>
          navigate('/swaggers/new', {
            state: {
              serviceId: props.serviceId
            }
          })
        }
        variant='contained'
      >
        Create One
      </Button>
    </Paper>
  )
}
