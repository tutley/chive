import axios from 'axios'

export const HTTP = axios.create({
  // TODO: Find a way to have a dev vs production setting for this
  baseURL: `http://localhost:3333/api/`
  // headers: {
  //   Authorization: 'Bearer {token}'
  // }
})
