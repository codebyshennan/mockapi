type Ok<T> = { data: T; err: undefined }
type Err = { data: undefined; err: Error }

export async function wrap<T>(promise: Promise<T>): Promise<Err | Ok<T>> {
  return await promise
    .then((data) => ({ data, err: undefined }))
    .catch((err) => ({ data: undefined, err }))
}
