import Vue from 'vue';
import Vuex from 'vuex';
import { loginUser } from '@/api/index';
import { getCoinInfo } from '@/api/coin';
import { connectWebSocket } from '@/api/websocket';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    username: '',
    socketConn: null,
    coinInfo: [],
  },
  getters: {
    isLogin(state) {
      return state.username !== '';
    },
  },
  mutations: {
    setUsername(state, username) {
      state.username = username;
    },
    clearUsername(state) {
      state.username = '';
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
      commit('setUsername', data);
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
