import axios from 'axios';

const commonInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
});

const upbitInst = axios.create({
  baseURL: process.env.VUE_APP_UPBIT_URL,
});

function registerUser(userData) {
  return commonInst.post('signup', userData);
}

function loginUser(userData) {
  return commonInst.post('login', userData);
}

function getCoinPrice(coinData) {
  return upbitInst.post('', coinData);
}

export { registerUser, loginUser, getCoinPrice };
