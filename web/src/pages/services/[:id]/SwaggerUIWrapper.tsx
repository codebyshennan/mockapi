import { ThemeProvider } from '@emotion/react'
import { Box, createTheme } from '@mui/material'
import React from 'react'
import SwaggerUI from 'swagger-ui-react'

export default function SwaggerUIWrapper(props: { swagger: object }) {
  return (
    <Box
      sx={{
        backgroundColor: 'white'
      }}
    >
      <SwaggerUI spec={props.swagger} />
    </Box>
  )
}
