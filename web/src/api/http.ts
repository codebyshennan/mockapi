import axios, { AxiosRequestHeaders } from 'axios'

export const httpGet = async (
  url: string,
  headers: AxiosRequestHeaders,
  params?: any
) =>
  axios({
    method: 'get',
    url,
    headers,
    params
  })

export const httpPost = async (
  url: string,
  headers: AxiosRequestHeaders,
  data: any
) =>
  axios({
    method: 'post',
    url,
    headers,
    data
  })

export const httpPut = async (
  url: string,
  headers: AxiosRequestHeaders,
  data: any
) =>
  axios({
    method: 'put',
    url,
    headers,
    data
  })

export const httpPatch = async (
  url: string,
  headers: AxiosRequestHeaders,
  data: any
) =>
  axios({
    method: 'patch',
    url,
    headers,
    data
  })

export const httpDelete = async (url: string, headers: AxiosRequestHeaders) =>
  axios({
    method: 'delete',
    url,
    headers
  })
