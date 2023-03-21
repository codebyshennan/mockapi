import { Box, Paper, Typography } from '@mui/material'
import React, { ReactNode } from 'react'

interface CardProps {
  title: string | ReactNode
  children: ReactNode
  onClick?: () => void
}

export default function SbCard({ title, children, onClick }: CardProps) {
  return (
    <Paper
      sx={{
        p: 2
      }}
      onClick={onClick}
    >
      {typeof title === 'string' ? (
        <Typography variant='h6'>{title}</Typography>
      ) : (
        <Box>{title}</Box>
      )}
      {children}
    </Paper>
  )
}
