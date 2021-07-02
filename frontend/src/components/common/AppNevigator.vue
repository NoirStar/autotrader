<template>
  <v-navigation-drawer class="pt-5" app clipped fixed mini-variant>
    <v-avatar
      v-for="n in 6"
      :key="n"
      :color="`grey ${n === 1 ? 'darken' : 'lighten'}-1`"
      :size="n === 1 ? 36 : 20"
      class="d-block text-center mx-auto mb-9"
    ></v-avatar>
  </v-navigation-drawer>
</template>

<script>
export default {};
</script>

<style></style>
