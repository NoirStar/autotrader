import Vue from 'vue';
import Vuex from 'vuex';
import { loginUser } from '@/api/index';
import { getCoinInfo } from '@/api/coin';
import { connectWebSocket } from '@/api/websocket';
import { getCookie, setCookie } from '@/cookies/index';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    nickname: getCookie('nickname') || '',
    token: getCookie('access_token') || '',
    socketConn: null,
    coinInfo: [],
  },
  getters: {
    isLogin(state) {
      return state.token !== '';
    },
  },
  mutations: {
    setNickname(state, nickname) {
      state.nickname = nickname;
    },
    clearNickname(state) {
      state.nickname = '';
    },
    setToken(state, token) {
      state.token = token;
    },
    clearToken(state) {
      state.token = '';
    },
    setCoinInfo(state, data) {
      state.coinInfo = data.filter(e => {
        if (e.market.substr(0, 3) === 'KRW') {
          e.marketShort = e.market.substr(4);
          return e;
        }
      });
    },
    setSocketConn(state, conn) {
      state.socketConn = conn;
    },
  },
  actions: {
    async LOGIN({ commit }, userData) {
      const { data } = await loginUser(userData);
      commit('setToken', getCookie('access_token'));
      commit('setNickname', data);
      setCookie('nickname', data);
      return data;
    },
    async COININFO({ commit }) {
      const { data } = await getCoinInfo();
      commit('setCoinInfo', data);
      return data;
    },
    async CONNECTION({ commit }) {
      const conn = await connectWebSocket();
      commit('setSocketConn', conn);
      return conn;
    },
  },
});
