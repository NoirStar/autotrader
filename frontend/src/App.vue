<template>
  <v-app>
    <AppHeader></AppHeader>
    <router-view></router-view>
  </v-app>
</template>

<script>
import AppHeader from '@/components/common/AppHeader.vue';
export default {
  components: {
    AppHeader,
  },
};
</script>

<style></style>
