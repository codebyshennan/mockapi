import React, { useCallback, useState } from 'react'
import * as yup from 'yup'
import { Form, Formik, FormikHelpers, FormikProps } from 'formik'
import { Box, Button, Paper, TextField, Typography } from '@mui/material'
import ServiceApi from '../../api/ServiceApi'
import { useAuth } from '../../context/AuthContext'
import { toast } from 'react-toastify'
import { useNavigate } from 'react-router-dom'

const validationSchema = yup.object({
  name: yup.string().required('Mock Server name is required')
})

interface ServicesFormValues {
  name: string
  desc: string
}

export interface ServicesFormProps {
  mode: 'new' | 'edit' | 'view'
  initialValues?: ServicesFormValues
}

function getFormTitle(mode: ServicesFormProps['mode']) {
  switch (mode) {
    case 'edit':
      return 'Edit Service'
    case 'new':
      return 'Create Service'
    case 'view':
      return 'View Service'
    default:
      // intentially return something that looks weird
      // to detect the bug
      return 's3rvic3'
  }
}

export default function ServiceForm(props: ServicesFormProps) {
  const { initialValues, mode } = props
  const values =
    mode === 'new' || !initialValues ? { name: '', desc: '' } : initialValues
  const formTitle = getFormTitle(mode)

  const [loading, setLoading] = useState<boolean>(false)
  const {
    authRes: { token }
  } = useAuth()
  const navigate = useNavigate()

  const onSubmit = useCallback(
    (
      values: ServicesFormValues,
      actions: FormikHelpers<ServicesFormValues>
    ) => {
      setLoading(true)
      ServiceApi.createService(token, values)
        .then((res) => {
          toast.success('API Service Created')
          navigate('../')
        })
        .catch((err) => toast.error(err))
        .finally(() => setLoading(false))
    },
    []
  )

  return (
    <Formik
      initialValues={values}
      onSubmit={onSubmit}
      validationSchema={validationSchema}
    >
      {(formProps: FormikProps<ServicesFormValues>) => {
        return (
          <Form>
            <Paper
              elevation={1}
              sx={{
                p: 2
              }}
            >
              <Typography variant='subtitle1' sx={{ mb: 2 }}>
                {formTitle}
              </Typography>

              <Box mb={2}>
                <TextField
                  fullWidth
                  name='name'
                  label='Service Name'
                  value={formProps.values.name}
                  onChange={formProps.handleChange}
                />

                <TextField
                  fullWidth
                  multiline
                  minRows={10}
                  name='desc'
                  label='Description'
                  value={formProps.values.desc}
                  onChange={formProps.handleChange}
                />
              </Box>
            </Paper>
            <Box display='flex' columnGap={2} mt={2}>
              <Button type='submit' variant='contained' color='success'>
                {formTitle}
              </Button>
              <Button type='reset' variant='outlined' color='warning'>
                Reset
              </Button>
            </Box>
          </Form>
        )
      }}
    </Formik>
  )
}
