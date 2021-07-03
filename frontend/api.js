import axios from 'axios';
import constants from './constants';

const api = axios.create({
  baseURL: constants.BACKEND_HOST,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

export default api;
