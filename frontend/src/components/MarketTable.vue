<template>
  <div>market</div>
</template>

<script>
import { getMarketInfo } from '@/api/coin';

export default {
  async created() {
    await getMarketInfo();
  },
};
</script>

<style></style>
