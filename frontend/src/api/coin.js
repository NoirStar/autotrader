import axios from 'axios';

const commonInst = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
});

function getCoinInfo() {
  return commonInst.get('coins');
}

export { getCoinInfo };
