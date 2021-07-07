<template>
  <v-main>
    <v-container>
      <MarketTable></MarketTable>
      <CostTable></CostTable>
    </v-container>
  </v-main>
</template>

<script>
import CostTable from '@/components/CostTable.vue';
import MarketTable from '@/components/MarketTable.vue';
export default {
  components: { CostTable, MarketTable },
};
</script>

<style></style>
