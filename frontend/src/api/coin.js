import axios from 'axios';
import { setInterceptors } from './common/interceptors';

// 엑시오스 초기화 함수
function createInstance() {
  const instance = axios.create({
    baseURL: process.env.VUE_APP_API_URL,
  });

  return setInterceptors(instance);
}

const instance = createInstance();

// 코인 리스트정보 API
function getCoinInfo() {
  return instance.get('coins');
}

// 코인 마켓리스트정보 API
function getMarketInfo(minute) {
  return instance.post('market', minute);
}

export { getCoinInfo, getMarketInfo };
