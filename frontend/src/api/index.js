import axios from 'axios';

const commonInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
});

const credentialInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  withCredentials: true,
});

function registerUser(userData) {
  return commonInst.post('signup', userData);
}

function loginUser(userData) {
  return credentialInst.post('login', userData);
}

function checkDuplicate(data) {
  return commonInst.post('check', data);
}

export { registerUser, loginUser, checkDuplicate };
