import React from 'react'
import Layout from '../../../layout/Layout'
import ServiceForm, { ServicesFormProps } from '../ServiceForm'
import * as yup from 'yup'
import { Form, Formik, FormikHelpers, FormikProps } from 'formik'
import { Box, Paper, TextField, Typography } from '@mui/material'

export default function ServicesNewPage() {
  return (
    <Layout title='Mock Servers'>
      <ServiceForm mode='new' />
    </Layout>
  )
}
