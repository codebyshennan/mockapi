import {
  Box,
  CssBaseline,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemText,
  Typography
} from '@mui/material'
import React, { ReactNode, useState } from 'react'
import MailIcon from '@mui/icons-material/Mail'
import { useNavigate } from 'react-router-dom'
import SandboxToastContainer from './SandboxToastContainer'
import { DRAWER_WIDTH } from '../context/CirclesPaletteContext'

const NAV_ITEMS = [
  {
    icon: <MailIcon />,
    name: 'Services',
    dst: '/services'
  },
  {
    icon: <MailIcon />,
    name: 'Mock Servers',
    dst: '/mockservers'
  },
  {
    icon: <MailIcon />,
    name: 'Account',
    dst: '/account'
  }
]

function NavBar() {
  const navigate = useNavigate()
  return (
    <List>
      {NAV_ITEMS.map((item) => (
        <ListItem key={item.name} disablePadding>
          <ListItemButton onClick={() => navigate(item.dst)}>
            {item.icon}
            <ListItemText primary={item.name} />
          </ListItemButton>
        </ListItem>
      ))}
    </List>
  )
}

export default function Navbar() {
  return (
    <>
      <CssBaseline />
      <Box sx={{ display: 'flex' }}>
        <Drawer
          variant='permanent'
          sx={{
            '& .MuiDrawer-paper': {
              boxSizing: 'border-box',
              width: DRAWER_WIDTH
            }
          }}
          open
        >
          <NavBar />
        </Drawer>
      </Box>

      <SandboxToastContainer />
    </>
  )
}
