import React, { useCallback, useState } from 'react'
import {
  FieldArray,
  FieldArrayRenderProps,
  Form,
  Formik,
  FormikHelpers,
  FormikProps
} from 'formik'
import * as yup from 'yup'
import {
  TextField,
  Button,
  Box,
  Typography,
  Paper,
  Drawer
} from '@mui/material'
import {
  MockEndPtGetData,
  MockServerGetData
} from '../../types/api/v1/mockServer.type'
import MoreHorizIcon from '@mui/icons-material/MoreHoriz'
import MockEndptForm from './MockEndptForm'
import DeleteIcon from '@mui/icons-material/Delete'
import { getMockServerCreatePayload, getMockServerUpdatePayload } from './utils'
import { useAuth } from '../../context/AuthContext'
import MockServerApi from '../../api/MockServerApi'
import { useNavigate, useParams } from 'react-router-dom'
import { toast } from 'react-toastify'
import ImportEndPoints from './ImportEndPoints'
import { SwaggerGetData } from '../../types/api/v1/swagger.type'
import CloneMockServer from './CloneMockServer'

const validationSchema = yup.object({
  name: yup.string().required('Mock Server name is required')
})

export interface MockServerFormValues {
  name: string
  endpts: MockEndPtFields[]
}

export type MockEndPtFields = Omit<MockEndPtGetData, 'id'> & {
  id?: string
}

export interface MockServerFormProps {
  mode: 'new' | 'edit' | 'view'
  initialValues?: MockServerFormValues
  useClone?: boolean
}

export default function MockServerForm(props: MockServerFormProps) {
  const { initialValues, mode, useClone = false } = props
  const formTitle = getFormTitle(mode)
  const values: MockServerFormValues = initialValues || {
    name: '',
    endpts: []
  }

  const { id: serverId } = useParams()
  const navigate = useNavigate()
  const [openDrawer, setOpenDrawer] = useState<boolean>(false)
  const [importDrawer, setImportDrawer] = useState<boolean>(false)
  const [cloneDrawer, setCloneDrawer] = useState<boolean>(useClone)
  const [mockResIdx, setMockResIdx] = useState<number>(-1)
  const {
    authRes: { token }
  } = useAuth()

  const setSelectedMockEndPt = useCallback((index: number) => {
    setMockResIdx(index)
    setOpenDrawer(true)
  }, [])

  const onSubmit = useCallback(
    (
      values: MockServerFormValues,
      _actions: FormikHelpers<MockServerFormValues>
    ) => {
      if (mode === 'view' || !token) return

      if (mode === 'new') {
        const payload = getMockServerCreatePayload(values)
        MockServerApi.createServer(token, payload)
          .then(() => toast.success('Done'))
          .catch((err) => {
            toast.error(err)
          })
        return
      }

      if (mode === 'edit' && serverId) {
        const payload = getMockServerUpdatePayload(values, initialValues)
        if (!payload) return

        MockServerApi.updateServer(serverId, token, payload)
          .then(() => toast.success('Done'))
          .catch((err) => {
            toast.error(err)
          })
      }
    },
    [token, serverId, mode]
  )

  const isDisabled = mode === 'view'

  return (
    <Formik
      initialValues={values}
      onSubmit={onSubmit}
      validationSchema={validationSchema}
    >
      {(formProps: FormikProps<MockServerFormValues>) => {
        return (
          <Form>
            <Paper elevation={1} sx={{ p: 2 }}>
              <Typography variant='subtitle1' sx={{ mb: 2 }}>
                {formTitle}
              </Typography>

              <Box mb={2}>
                <Typography variant='subtitle1' pb={0.5}>
                  Overall
                </Typography>
                <TextField
                  fullWidth
                  disabled={isDisabled}
                  name='name'
                  label='Mock Server Name'
                  value={formProps.values.name}
                  onChange={formProps.handleChange}
                />
              </Box>

              <Typography variant='subtitle1' pb={2}>
                {!isDisabled ? 'Edit' : ''} Mock Server Endpoints
              </Typography>
              <FieldArray
                name='endpts'
                render={(arrayHelpers) => (
                  <Box>
                    {formProps.values && formProps.values.endpts.length > 0 ? (
                      formProps.values.endpts.map((_endpt, index) => (
                        <Box key={`${index}-endpt`}>
                          <Box display='flex' columnGap={2} mb={2}>
                            <MockEndptForm
                              disabled={isDisabled}
                              index={index}
                              mode='row'
                              formProps={formProps}
                            />

                            <Button
                              type='button'
                              onClick={() => setSelectedMockEndPt(index)}
                              variant='contained'
                              color='secondary'
                            >
                              <MoreHorizIcon />
                            </Button>

                            {!isDisabled && (
                              <Button
                                disabled={isDisabled}
                                type='button'
                                onClick={() => arrayHelpers.remove(index)}
                                variant='contained'
                                color='error'
                              >
                                <DeleteIcon />
                              </Button>
                            )}
                          </Box>
                          {index === formProps.values.endpts.length - 1 &&
                            !isDisabled && (
                              <Box display='flex' columnGap={2}>
                                <Button
                                  disabled={isDisabled}
                                  type='button'
                                  fullWidth
                                  variant='contained'
                                  onClick={() =>
                                    arrayHelpers.push({
                                      method: '',
                                      endptRegex: '',
                                      resCode: '',
                                      resBody: '',
                                      timeout: 0,
                                      writes: {
                                        data: '',
                                        ttl: 0
                                      }
                                    })
                                  }
                                >
                                  Add a mock endpoint
                                </Button>

                                <Button
                                  disabled={isDisabled}
                                  type='button'
                                  fullWidth
                                  variant='outlined'
                                  onClick={() => setImportDrawer(true)}
                                >
                                  Import endpoints from Swagger
                                </Button>
                              </Box>
                            )}
                        </Box>
                      ))
                    ) : (
                      <Box display='flex' columnGap={2}>
                        <Button
                          disabled={isDisabled}
                          type='button'
                          fullWidth
                          variant='contained'
                          onClick={() =>
                            arrayHelpers.push({
                              method: '',
                              endptRegex: '',
                              resCode: '',
                              resBody: '',
                              timeout: 0,
                              writes: {
                                data: '',
                                ttl: 0
                              }
                            })
                          }
                        >
                          Add a mock endpoint
                        </Button>

                        <Button
                          disabled={isDisabled}
                          type='button'
                          fullWidth
                          variant='outlined'
                          onClick={() => setImportDrawer(true)}
                        >
                          Import endpoints from Swagger
                        </Button>
                      </Box>
                    )}

                    {/* Drawer for editing endpoints. */}
                    <Drawer
                      anchor='right'
                      open={openDrawer}
                      onClose={() => setOpenDrawer(false)}
                    >
                      {mockResIdx >= 0 && (
                        <Box display='flex' flexDirection='column' p={3}>
                          <Typography variant='h6' pb={2}>
                            {!isDisabled ? 'Edit' : ''} Endpoint
                          </Typography>
                          <MockEndptForm
                            disabled={isDisabled}
                            index={mockResIdx}
                            formProps={formProps}
                            mode='column'
                          />
                        </Box>
                      )}
                    </Drawer>

                    {/* Drawer for importing endpoints. */}
                    <Drawer
                      anchor='right'
                      open={importDrawer}
                      onClose={() => setImportDrawer(false)}
                    >
                      <Box display='flex' flexDirection='column' p={3}>
                        <Typography variant='h6' pb={2}>
                          Import Endpoints
                        </Typography>
                        <ImportEndPoints onConfirm={onImport(arrayHelpers)} />
                      </Box>
                    </Drawer>

                    {/* Drawer for cloning an existing mock server. */}
                    <Drawer
                      anchor='right'
                      open={cloneDrawer}
                      onClose={() => setCloneDrawer(false)}
                    >
                      <Box display='flex' flexDirection='column' p={3}>
                        <Typography variant='h6' pb={2}>
                          Clone an existing Mock Server
                        </Typography>
                        <CloneMockServer onConfirm={onClone(arrayHelpers)} />
                      </Box>
                    </Drawer>
                  </Box>
                )}
              />
            </Paper>
            {!isDisabled && (
              <Box display='flex' columnGap={2} mt={2}>
                <Button type='submit' variant='contained' color='success'>
                  {formTitle}
                </Button>
                <Button type='reset' variant='outlined' color='warning'>
                  Reset
                </Button>
              </Box>
            )}
            {isDisabled && (
              <Box display='flex' columnGap={2} mt={2}>
                <Button
                  type='button'
                  variant='contained'
                  color='success'
                  onClick={() => navigate(`/mockservers/${serverId}/update`)}
                >
                  Edit Mock Server
                </Button>
              </Box>
            )}
          </Form>
        )
      }}
    </Formik>
  )
}

/* Returns the form title. */
function getFormTitle(mode: MockServerFormProps['mode']) {
  switch (mode) {
    case 'edit':
      return 'Edit Mock Server'
    case 'new':
      return 'Create Mock Server'
    case 'view':
      return 'View Mock Server'
    default:
      // intentially return something that looks weird
      // to detect the bug
      return 'mOck S3rv3r'
  }
}

/**
 * Add things into the form's array field on import.
 */
const onImport =
  (arrayHelpers: FieldArrayRenderProps) => (d: SwaggerGetData) => {
    d.endpts
      .map((ele) => ({
        ...ele,
        method: ele.method,
        endptRegex: ele.endptRegex,
        resCode: '200',
        resBody: '',
        timeout: 0,
        writes: {
          data: '',
          ttl: 0
        }
      }))
      .forEach((ele) => arrayHelpers.push(ele))
  }

/**
 * Add things into the form's array field on clone.
 */
const onClone =
  (arrayHelpers: FieldArrayRenderProps) => (d: MockServerGetData) => {
    d.endpts
      .map((ele) => ({
        ...ele,
        method: ele.method,
        endptRegex: ele.endptRegex,
        resCode: '200',
        resBody: '',
        timeout: 0,
        writes: {
          data: '',
          ttl: 0
        }
      }))
      .forEach((ele) => arrayHelpers.push(ele))
  }
