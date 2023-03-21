import { ThemeProvider } from '@emotion/react'
import React from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import AccountIndexPage from '../pages/account/AccountIndexPage'
import MockServersIndexPage from '../pages/mockservers/MockServerIndexPage'
import ServicesIndexPage from '../pages/services/ServicesIndexPage'
import MockServersNewPage from '../pages/mockservers/new/MockServersNewPage'
import MockServersViewPage from '../pages/mockservers/[id]/MockServersViewPage'
import ServicesViewPage from '../pages/services/[:id]/ServicesViewPage'
import ServicesNewPage from '../pages/services/new/ServicesNewPage'
import SwaggerNewPage from '../pages/swaggers/new/SwaggerNewPage'
import MockServersUpdatePage from '../pages/mockservers/[id]/update/MockServersUpdatePage'

export default function AuthenticatedApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/account' element={<AccountIndexPage />} />

        <Route path='/mockservers' element={<MockServersIndexPage />} />
        <Route path='/mockservers/:id' element={<MockServersViewPage />} />
        <Route
          path='/mockservers/:id/update'
          element={<MockServersUpdatePage />}
        />
        <Route path='/mockservers/new' element={<MockServersNewPage />} />

        <Route path='/services' element={<ServicesIndexPage />} />
        <Route path='/services/:id' element={<ServicesViewPage />} />
        <Route path='/services/new' element={<ServicesNewPage />} />

        <Route path='/swaggers/new' element={<SwaggerNewPage />} />

        <Route path='*' element={<Navigate to='/services' />} />
      </Routes>
    </BrowserRouter>
  )
}
