import { Box, Typography } from '@mui/material'
import React, { ReactNode, useEffect, useState } from 'react'
import { DRAWER_WIDTH } from '../context/CirclesPaletteContext'
import Navbar from './Navbar'

export default function Layout(props: { children: ReactNode; title: string }) {
  const [width, setWidth] = useState<number>(window.innerWidth)

  function handleWindowSizeChange() {
    setWidth(window.innerWidth)
  }

  useEffect(() => {
    window.addEventListener('resize', handleWindowSizeChange)
    return () => {
      window.removeEventListener('resize', handleWindowSizeChange)
    }
  }, [])

  return (
    <>
      <Navbar />
      <Box
        component='main'
        flexGrow={1}
        p={3}
        marginLeft={`${DRAWER_WIDTH}px`}
        width={`${(width - DRAWER_WIDTH - 2 * 24) * 0.7}px`}
      >
        <Typography variant='h5' pb={3}>
          {props.title}
        </Typography>
        <>{props.children}</>
      </Box>
    </>
  )
}
