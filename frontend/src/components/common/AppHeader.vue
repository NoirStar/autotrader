<template>
  <div>
    <v-app-bar color="teal lighten-3" dense dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>Page title</v-toolbar-title>

      <v-spacer></v-spacer>
      <template v-if="isLogin">
        <span>{{ $store.state.username }}</span>
        <a href="javascript:;" @click="logoutUser">로그아웃</a>
      </template>
      <template v-else>
        <router-link to="/login">로그인</router-link> |
        <router-link to="/signup">회원가입</router-link>
      </template>
      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-menu left bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon v-bind="attrs" v-on="on">
            <v-icon>mdi-dots-vertical</v-icon>
          </v-btn>
        </template>

        <v-list>
          <v-list-item v-for="n in 5" :key="n" @click="() => {}">
            <v-list-item-title>Option {{ n }}</v-list-item-title>
          </v-list-item>
        </v-list>
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

<style></style>
