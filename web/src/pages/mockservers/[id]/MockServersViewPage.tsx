import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import MockServerApi from '../../../api/MockServerApi'
import { useAuth } from '../../../context/AuthContext'
import Layout from '../../../layout/Layout'
import { MockServerGetData } from '../../../types/api/v1/mockServer.type'
import MockServerForm from '../MockServerForm'

export default function MockServersViewPage() {
  const { id } = useParams()
  const {
    authRes: { token }
  } = useAuth()

  const [mockServer, setMockServer] = useState<MockServerGetData | undefined>()
  useEffect(() => {
    if (!id || !token) return

    MockServerApi.getServerById(id, token).then((res) =>
      setMockServer(res.data)
    )
  }, [])

  return (
    <Layout title='View Mock Server'>
      {mockServer && <MockServerForm mode='view' initialValues={mockServer} />}
    </Layout>
  )
}
