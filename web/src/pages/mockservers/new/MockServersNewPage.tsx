import React from 'react'
import { useLocation } from 'react-router-dom'
import Layout from '../../../layout/Layout'
import MockServerForm from '../MockServerForm'

type PAGE_STATE = {
  useClone: boolean
}

export default function MockServersNewPage() {
  const pageState = useLocation().state as PAGE_STATE | null

  return (
    <Layout title='Mock Servers'>
      <MockServerForm mode='new' useClone={pageState?.useClone ?? false} />
    </Layout>
  )
}
