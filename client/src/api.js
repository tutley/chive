import axios from 'axios'

// IF you use JWT Auth, this is how you would add the token to api requests
// let isLoggedIn = localStorage.getItem('token') != null
// let authString = ''

// if (isLoggedIn) {
//   authString = 'Bearer ' + localStorage.getItem('token')
// }

var apiURL = process.env.API_ROOT + '/api'

const config = {
  baseURL: apiURL
  // headers: {
  //   ...(isLoggedIn ? {Authorization: authString} : {})
  // }
}

export const HTTP = axios.create(config)
