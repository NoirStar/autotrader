<template>
  <v-card>
    <v-card-title>
      코인시세
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
    <v-data-table
      :headers="headers"
      :items="$store.state.coinInfo"
      :search="search"
      item-key="market"
    >
      <template class="d-flex flex-row" v-slot:[`item.korean_name`]="{ item }">
        <div>
          <img
            class="coin-logo"
            :src="`https://static.upbit.com/logos/${item.market}.png`"
          />
          {{ item.korean_name }}
        </div>
        <div>
          <small> {{ item.market }}</small>
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      search: '',
      headers: [
        {
          text: '코인',
          align: 'start',
          value: 'korean_name',
        },
        {
          text: '현재가',
          value: '',
        },
        {
          text: '김프',
          value: '',
        },
        {
          text: '전일대비',
          value: '',
        },
      ],
    };
  },
  methods: {},
  computed: {
    parseCoinName(market) {
      return market.substr(3);
    },
  },
  async created() {
    const data = await this.$store.dispatch('COININFO');
    console.log(data);
  },
};
</script>

<style scoped>
.coin-logo {
  width: 0.775rem;
}
</style>
