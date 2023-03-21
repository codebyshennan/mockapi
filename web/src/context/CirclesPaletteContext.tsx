import { createTheme } from '@mui/material'

const theme = createTheme({
  palette: {
    mode: 'dark'
  }
})

export const DRAWER_WIDTH = 240

export default theme

declare module '@mui/material/styles' {
  interface Theme {}
  // allow configuration using `createTheme`
  interface ThemeOptions {}
}
