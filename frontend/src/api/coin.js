import axios from 'axios';

const commonInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
});

const upbitInst = axios.create({
  baseURL: process.env.VUE_APP_UPBIT_URL,
});

function getCoinPrice(coinData) {
  return upbitInst.post('', coinData);
}

function getCoinInfo() {
  return commonInst.get('coins');
}

export { getCoinPrice, getCoinInfo };
