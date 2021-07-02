<template>
  <v-app>
    <AppHeader></AppHeader>
    <AppNevigator></AppNevigator>
    <router-view></router-view>
  </v-app>
</template>

<script>
import AppHeader from '@/components/common/AppHeader.vue';
import AppNevigator from '@/components/common/AppNevigator.vue';
export default {
  components: {
    AppHeader,
    AppNevigator,
  },
};
</script>

<style></style>
