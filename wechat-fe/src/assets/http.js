import axios from 'axios';

const http = axios.create({
//   headers: {
//     'Content-type': 'application/x-www-form-urlencoded'
//   },
//   baseURL: 'http://127.0.0.1:8989',
  timeout: 8000
});
http.interceptors.request.use(
  config => {
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

export default http;
