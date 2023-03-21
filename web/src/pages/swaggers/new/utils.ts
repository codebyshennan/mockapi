import { httpGet } from '../../../api/http'
import { wrap } from '../../../common/utils'

/**
 * Parses the url into a swagger document.
 */
export async function parseUrlContent(url: string): Promise<object | null> {
  // Retrieve from /swagger-ui-init.js
  const swaggerInit = await getFromSwaggerUiInit(url)
  if (swaggerInit) {
    return swaggerInit
  }

  return swaggerInit

  // If not, retrieve swagger spec from url through GET request to /doc.json
  const docJson = await getFromDocJson(url)
  return docJson
}

async function getFromSwaggerUiInit(givenUrl: string): Promise<object | null> {
  const { data: res, err } = await wrap(
    httpGet(givenUrl + '/swagger-ui-init.js', {})
  )
  if (err) {
    console.error(err)
    return null
  }

  try {
    const resStr = JSON.stringify(res.data)
    console.log(resStr)

    // Swagger spec json string starts after `"swaggerDoc{": ` in the response
    const swaggerSpecStrStartIndex =
      resStr.indexOf(`"swaggerDoc": {`) + '"swaggerDoc": {'.length

    // Swagger spec json string ends 4 characters before `"customOptions"` in the response
    // There's some trailing white space to minus off, that .trim() can't seem to remove
    const swaggerSpecStrEndIndex = resStr.indexOf(`"customOptions"`) - 5

    if (swaggerSpecStrStartIndex === -1 || swaggerSpecStrEndIndex === -1) {
      return null
    }

    const spec =
      '{' +
      resStr
        .substring(swaggerSpecStrStartIndex, swaggerSpecStrEndIndex)
        .trim() +
      '}'

    return JSON.parse(spec)
  } catch (error) {
    console.error(error)
    return null
  }
}

async function getFromDocJson(givenUrl: string): Promise<object | null> {
  const { data: res, err } = await wrap(httpGet(givenUrl + '/doc.json', {}))
  if (err) {
    // first way works
    return null
  }
  return res.data
}
