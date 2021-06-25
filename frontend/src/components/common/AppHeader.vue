<template>
  <div>
    <v-app-bar color="grey darken-4" dense dark>
      <router-link class="header-link" to="/">
        <v-toolbar-title>Tree Autotrader </v-toolbar-title>
      </router-link>

      <v-spacer></v-spacer>
      <template v-if="isLogin">
        <span class="mr-3">{{ $store.state.username }} 님</span>
        <a href="javascript:;" class="header-link" @click="logoutUser">
          로그아웃
        </a>
      </template>
      <template v-else>
        <router-link class="header-link mr-3" to="/login"> 로그인 </router-link>
        <router-link class="header-link" to="/signup">회원가입</router-link>
      </template>

      <v-menu left bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>
      </v-menu>
    </v-app-bar>
  </div>
</template>

<script>
export default {
  computed: {
    isLogin() {
      return this.$store.getters.isLogin;
    },
  },
  methods: {
    logoutUser() {
      this.$store.commit('clearUsername');
      this.$router.push('/login');
    },
  },
};
</script>

<style scoped>
.header-link {
  text-decoration: none;
  color: #b2c0cc;
  font-weight: 500;
  transition-delay: initial;
  transition-duration: 0.08s;
  transition-property: all;
  transition-timing-function: ease-in-out;
}

.header-link:hover {
  color: white;
}
</style>
