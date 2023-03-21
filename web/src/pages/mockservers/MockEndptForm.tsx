import {
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Box,
  BoxProps,
  Button,
  Typography,
  Accordion,
  AccordionDetails,
  AccordionSummary
} from '@mui/material'
import { FormikProps } from 'formik'
import React, { useCallback, useState } from 'react'
import { MockServerFormValues } from './MockServerForm'
import AceEditor from 'react-ace'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'

import 'ace-builds/src-noconflict/mode-json'
import 'ace-builds/src-noconflict/theme-terminal'
import 'ace-builds/src-noconflict/ext-language_tools'

interface MockEndptFormProps {
  disabled: boolean
  index: number
  formProps: FormikProps<MockServerFormValues>
  mode: 'column' | 'row'
}

const rowStyles: BoxProps = {
  columnGap: 2
}

const colStyles: BoxProps = {
  flexDirection: 'column',
  rowGap: 2,
  width: 500
}

export default function MockEndptForm(props: MockEndptFormProps) {
  const { index, formProps, mode = 'row', disabled: isDisabled } = props

  const [editorOpen, setEditorOpen] = useState<boolean>(false)
  const [timeoutOpen, setTimoutOpen] = useState<boolean>(false)
  const [writePolicy, setWritePolicy] = useState<boolean>(false)

  const onResBodyChange = useCallback((value: string) => {
    formProps.setValues((prev) => {
      prev.endpts[index].resBody = value
      return prev
    })
  }, [])

  const onWriteDataChange = useCallback((value: string) => {
    formProps.setValues((prev) => {
      if (prev.endpts[index].writes !== null) {
        // ts bug, the type guard is not working as expected
        // @ts-ignore
        prev.endpts[index].writes.data = value
      } else {
        prev.endpts[index].writes = {
          data: value,
          dest: ''
        }
      }

      return prev
    })
  }, [])

  return (
    <Box
      display='flex'
      flexGrow={1}
      {...(mode === 'row' && rowStyles)}
      {...(mode === 'column' && colStyles)}
    >
      <Box display='flex'></Box>
      <FormControl sx={{ flexBasis: mode === 'row' ? '20%' : undefined }}>
        <InputLabel id={`fieldarr-endpoint-${index}`}>Method</InputLabel>
        <Select
          disabled={isDisabled}
          name={`endpts[${index}].method`}
          labelId={`fieldarr-endpoint-${index}`}
          value={formProps.values.endpts[index].method}
          onChange={formProps.handleChange}
        >
          <MenuItem value='GET'>GET</MenuItem>
          <MenuItem value='POST'>POST</MenuItem>
          <MenuItem value='PUT'>PUT</MenuItem>
          <MenuItem value='PATCH'>PATCH</MenuItem>
          <MenuItem value='DELETE'>DELETE</MenuItem>
        </Select>
      </FormControl>

      <TextField
        disabled={isDisabled}
        sx={{ flexGrow: mode === 'row' ? 1 : undefined }}
        name={`endpts[${index}].endptRegex`}
        label='Endpoint'
        value={formProps.values.endpts[index].endptRegex}
        onChange={formProps.handleChange}
      />

      <TextField
        disabled={isDisabled}
        sx={{ flexBasis: mode === 'row' ? '15%' : undefined }}
        name={`endpts[${index}].resCode`}
        label='Status Code'
        value={formProps.values.endpts[index].resCode}
        onChange={formProps.handleChange}
      />

      {mode === 'column' && (
        <Box>
          <Accordion
            expanded={editorOpen}
            onChange={() => setEditorOpen(!editorOpen)}
            TransitionProps={{ unmountOnExit: true }}
            disableGutters
          >
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant='subtitle2'>
                {isDisabled ? 'View' : 'Edit'} Response Body
              </Typography>
            </AccordionSummary>
            <AccordionDetails>
              <AceEditor
                style={{
                  width: '100%'
                }}
                readOnly={isDisabled}
                placeholder='Enter Response Body here'
                mode='json'
                theme='monokai'
                name={`endpts[${index}].resBody`}
                fontSize={14}
                showPrintMargin={true}
                showGutter={true}
                value={formProps.values.endpts[index].resBody ?? ''}
                debounceChangePeriod={500}
                onChange={onResBodyChange}
                setOptions={{
                  enableBasicAutocompletion: false,
                  enableLiveAutocompletion: false,
                  enableSnippets: false,
                  showLineNumbers: true,
                  tabSize: 2
                }}
              />
            </AccordionDetails>
          </Accordion>
        </Box>
      )}

      {mode === 'column' && (
        <Box>
          <Accordion
            expanded={timeoutOpen}
            onChange={() => setTimoutOpen(!timeoutOpen)}
            TransitionProps={{ unmountOnExit: true }}
            disableGutters
          >
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant='subtitle2'>
                {isDisabled ? 'View' : 'Edit'} Timeout
              </Typography>
            </AccordionSummary>
            <AccordionDetails>
              <TextField
                disabled={isDisabled}
                name={`endpts[${index}].timeout`}
                label='Timeout duration'
                value={formProps.values.endpts[index].timeout ?? 0}
                onChange={formProps.handleChange}
              />
            </AccordionDetails>
          </Accordion>
        </Box>
      )}

      {mode === 'column' && (
        <Box>
          <Accordion
            expanded={writePolicy}
            onChange={() => setWritePolicy(!writePolicy)}
            TransitionProps={{ unmountOnExit: true }}
            disableGutters
          >
            <AccordionSummary expandIcon={<ExpandMoreIcon />}>
              <Typography variant='subtitle2'>
                {isDisabled ? 'View' : 'Edit'} Write Policy
              </Typography>
            </AccordionSummary>
            <AccordionDetails>
              <TextField
                fullWidth
                disabled={isDisabled}
                name={`endpts[${index}].writes.dest`}
                label='Data destination'
                value={formProps.values.endpts[index].writes?.dest ?? ''}
                onChange={formProps.handleChange}
              />
              <AceEditor
                style={{
                  width: '100%'
                }}
                readOnly={isDisabled}
                placeholder='Enter data to insert here'
                mode='json'
                theme='monokai'
                name={`endpts[${index}].writes.data`}
                fontSize={14}
                showPrintMargin={true}
                showGutter={true}
                value={formProps.values.endpts[index].writes?.data ?? ''}
                debounceChangePeriod={500}
                onChange={onWriteDataChange}
                setOptions={{
                  enableBasicAutocompletion: false,
                  enableLiveAutocompletion: false,
                  enableSnippets: false,
                  showLineNumbers: true,
                  tabSize: 2
                }}
              />
            </AccordionDetails>
          </Accordion>
        </Box>
      )}
    </Box>
  )
}
