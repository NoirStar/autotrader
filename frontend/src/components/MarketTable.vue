<template>
  <v-card class="mx-auto mt-16">
    <v-card-title>
      마켓
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>

    <v-data-table :headers="headers" :items="marketData" :search="search">
      <template class="d-flex flex-row" v-slot:[`item.codeName`]="{ item }">
        <div>
          <img
            class="coin-logo"
            :src="`https://static.upbit.com/logos/${
              item.code.split('-')[1]
            }.png`"
          />
          {{ item.codeName }}
        </div>
        <div>
          <small> {{ item.code }}</small>
        </div>
      </template>
      <template v-slot:[`item.AskTotal`]="{ item }">
        {{ item.AskTotal.toLocaleString() }}
      </template>
      <template v-slot:[`item.BidTotal`]="{ item }">
        {{ item.BidTotal.toLocaleString() }}
      </template>
      <template v-slot:[`item.AskCount`]="{ item }">
        {{ item.AskCount.toLocaleString() }}
      </template>
      <template v-slot:[`item.BidCount`]="{ item }">
        {{ item.BidCount.toLocaleString() }}
      </template>
      <template v-slot:[`item.change`]="{ item }">
        <template v-if="item.change > 0">
          <div class="price-up">▲ {{ item.change }}</div>
        </template>
        <template v-else-if="item.change < 0">
          <div class="price-down">▼ {{ Math.abs(item.change) }}</div>
        </template>
        <template v-else> - </template>
      </template>
    </v-data-table>
  </v-card>
</template>

<script>
import { getCoinInfo, getMarketInfo } from '@/api/coin';

export default {
  data() {
    return {
      conn: null,
      search: '',
      interval: undefined,
      prevMarketRank: {},
      marketData: [],
      coinInfo: [],
      minute: 5,
      minutes: [1, 3, 5, 15, 30, 60],
      headers: [
        {
          text: '코인',
          align: 'start',
          value: 'codeName',
        },
        {
          text: '매수개수',
          value: 'AskCount',
        },
        {
          text: '매수량',
          value: 'AskTotal',
        },
        {
          text: '매도개수',
          value: 'BidCount',
        },
        {
          text: '매도량',
          value: 'BidTotal',
        },
        {
          text: '매수순위',
          value: 'rank',
        },
        {
          text: '매수순위변동',
          value: 'change',
        },
      ],
    };
  },
  async created() {
    const { data } = await getCoinInfo();
    this.coinInfo = data;
    this.getMarketData(this.minute);
    this.interval = setInterval(() => {
      this.getMarketData(this.minute);
    }, this.minute * 60000);
  },
  beforeDestroy() {
    if (this.interval) {
      clearInterval(this.interval);
      this.interval = undefined;
    }
  },
  methods: {
    async getMarketData(min) {
      try {
        const minute = {
          min: min,
        };
        console.log('get marketdata');
        const { data } = await getMarketInfo(minute);

        let temp = JSON.parse(atob(data));

        for (const key of Object.keys(temp)) {
          temp[key]['code'] = key;
        }
        this.marketData = [];
        for (const value of Object.values(temp)) {
          this.coinInfo.forEach(e => {
            if (e.market == value['code']) {
              value['codeName'] = e.korean_name;
            }
          });
          value['AskTotal'] = parseInt(value['AskTotal']);
          value['BidTotal'] = parseInt(value['BidTotal']);
          this.marketData.push(value);
        }
        this.marketData.sort((a, b) => {
          return b.AskTotal - a.AskTotal;
        });

        if (Object.keys(this.prevMarketRank).length === 0) {
          this.marketData.forEach((e, idx) => {
            e['rank'] = idx + 1;
            this.prevMarketRank[e.code] = idx + 1;
          });
        } else {
          this.marketData.forEach((e, idx) => {
            e['rank'] = idx + 1;
            if (Object.keys(this.prevMarketRank).includes(e['code'])) {
              e['change'] = this.prevMarketRank[e.code] - (idx + 1);
            } else {
              e['change'] = '-';
            }
            this.prevMarketRank[e.code] = idx + 1;
          });
        }
      } catch (error) {
        console.log(error.response.data);
      }
    },
  },
};
</script>

<style scoped>
.coin-logo {
  width: 0.775rem;
}

.price-up {
  color: #f06d6f;
}

.price-down {
  color: #4480da;
}
</style>
