import Vue from 'vue';
import VueCookies from 'vue-cookies';

Vue.use(VueCookies);

function getCookie(key) {
  return Vue.$cookies.get(key);
}

function setCookie(key, value) {
  Vue.$cookies.set(key, value);
}

function removeCookie(key) {
  Vue.$cookies.remove(key);
}

export { getCookie, setCookie, removeCookie };
