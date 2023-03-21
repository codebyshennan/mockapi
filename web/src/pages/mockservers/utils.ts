import { MockEndPtCreateData, MockEndPtUpdateData, MockServerCreateData, MockServerPatchData } from '../../types/api/v1/mockServer.type'
import { OpenApiV3 } from '../../types/openApi'
import { MockEndPtFields, MockServerFormValues } from './MockServerForm'

/**
 * Gets the request body for mock server update
 */
export function getMockServerUpdatePayload(
  newValues: MockServerFormValues,
  prev?: MockServerFormValues
): MockServerPatchData | undefined {
  if (!prev) return
  const res: MockServerPatchData = {
    responses: {
      create: [],
      update: []
    }
  }

  if (prev.name !== newValues.name) {
    res.name = newValues.name
  }

  for (const endpt of newValues.endpts) {
    if (!endpt.id) {
      if (endpt.endptRegex && endpt.method && endpt.resCode) {
        const t = {
          ...endpt,
          method: endpt.method,
          resCode: parseInt(endpt.resCode.toString()),
          resBody: endpt.resBody === '' ? null : minifyJson(endpt.resBody ?? '')
        }

        if (endpt.timeout) {
          t.timeout = parseInt(endpt.timeout?.toString() ?? '0')
          if (t.timeout === 0) {
            t.timeout = null
          }
        }

        if (endpt.writes && endpt.writes.dest !== '' && minifyJson(endpt.writes.data) !== '') {
          t.writes = {
            dest: endpt.writes.dest.toString(),
            data: minifyJson(endpt.writes.data) ?? ''
          }
        } else {
          t.writes = null
        }


        res.responses.create.push(t)
      }
      continue
    }

    const idx = prev.endpts.findIndex(element => element.id === endpt.id)
    if (idx === -1) {
      console.error('Endpoint dropped')
      continue
    } else {
      const t = {
        ...endpt,
        id: endpt.id,
        method: endpt.method,
        resCode: parseInt(endpt.resCode.toString()),
        resBody: endpt.resBody === '' ? null : minifyJson(endpt.resBody ?? '')
      }

      if (endpt.timeout) {
        t.timeout = parseInt(endpt.timeout?.toString())
        if (t.timeout === 0) {
          t.timeout = null
        }
      }

      if (endpt.writes && endpt.writes.dest !== '' && minifyJson(endpt.writes.data) !== '') {
        t.writes = {
          dest: endpt.writes.dest,
          data: minifyJson(endpt.writes.data) ?? ''
        }
      } else {
        t.writes = null
      }

      res.responses.update.push(t)
    }
  }

  return res
}

/**
 * Gets the request body for mock server create
 */
export function getMockServerCreatePayload(values: MockServerFormValues): MockServerCreateData {
  return {
    name: values.name,
    responses: values.endpts.map(element => {
      let body: string | null = !element.resBody ? null : minifyJson(element.resBody)
      let writeData: string | null = !element.writes?.data ? '' : minifyJson(element.writes.data)

      const t: MockEndPtCreateData = {
        ...element,
        resBody: body,
        resCode: parseInt(element.resCode.toString()),
        writes: null
      }

      if (writeData !== null && writeData !== '') {
        t.writes = {
          dest: element.writes?.dest ?? '',
          data: writeData
        }
      }

      if (t.timeout === 0 || !t.timeout) {
        t.timeout = null
      } else {
        t.timeout = parseInt(t.timeout.toString())
      }

      return t
    })
  }
}

export function importEndpoints(spec: OpenApiV3): MockEndPtFields[] {
  try {
    const endpts: MockEndPtFields[] = []

    for (const endpt in spec.paths) {
      for (const method in spec.paths[endpt]) {
        endpts.push({
          method: method.toUpperCase() as any,
          endptRegex: endpt,
          resCode: 200,
          resBody: null,
          writes: null,
          timeout: null
        })
      }
    }

    return endpts
  } catch (error) {
    console.error(error)
    return []
  }
}

function minifyJson(json: string): string | null {
  try {
    return JSON.stringify(JSON.parse(json.trim()))
  } catch {
    return json?.trim() ?? null
  }
}