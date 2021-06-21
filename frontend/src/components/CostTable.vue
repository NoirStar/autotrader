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
      :items="coinInfo"
      :search="search"
      item-key="price"
    >
      <template class="d-flex flex-row" v-slot:[`item.korean_name`]="{ item }">
        <div>
          <img
            class="coin-logo"
            :src="`https://static.upbit.com/logos/${item.marketShort}.png`"
          />
          {{ item.korean_name }}
        </div>
        <div>
          <small> {{ item.marketShort }}</small>
        </div>
      </template>

      <template v-slot:[`item.price`]="{ item }">
        {{ item.price.toLocaleString() }}
      </template>
      <template v-slot:[`item.changeRate`]="{ item }">
        <div>{{ (item.changeRate * 100).toFixed(2) }}%</div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      conn: null,
      search: '',
      reqFormat: [],
      codes: [],
      coinInfo: [],
      headers: [
        {
          text: '코인',
          align: 'start',
          value: 'korean_name',
        },
        {
          text: '현재가',
          value: 'price',
        },
        {
          text: '김프',
          value: '',
        },
        {
          text: '전일대비',
          value: 'changeRate',
        },
        {
          text: '고가대비(52주)',
          value: 'highest52',
        },
        {
          text: '저가대비(52주)',
          value: 'lowest52',
        },
        {
          text: '거래액(일)',
          value: 'accTradePrice',
        },
      ],
    };
  },
  methods: {
    MakeCodes() {
      return this.$store.state.coinInfo.reduce(
        (acc, cur) => (acc.push(`"${cur.market}"`), acc),
        [],
      );
    },
  },
  computed: {},
  async created() {
    await this.$store.dispatch('COININFO');
    this.coinInfo = this.$store.state.coinInfo;
    const conn = await this.$store.dispatch('CONNECTION');
    conn.onmessage = ({ data }) => {
      if (data instanceof Blob) {
        let reader = new FileReader();
        reader.onload = () => {
          let result = JSON.parse(reader.result);
          this.coinInfo.forEach(e => {
            if (e.market === result.cd) {
              this.$set(e, 'price', result.tp);
              this.$set(e, 'changePrice', result.scp);
              this.$set(e, 'changeRate', result.scr);
              this.$set(e, 'accTradePrice', result.atp24h);
              this.$set(e, 'askBid', result.ab);
              this.$set(e, 'highest52', result.h52wp);
              this.$set(e, 'lowest52', result.l52wp);
            }
          });
        };
        reader.readAsText(data);
      }
    };
    conn.send(
      `[{"ticket":"tree"},{"type":"ticker","codes":[${this.MakeCodes()}]},{"format":"SIMPLE"}]`,
    );
  },
};
</script>

<style scoped>
.coin-logo {
  width: 0.775rem;
}
</style>
