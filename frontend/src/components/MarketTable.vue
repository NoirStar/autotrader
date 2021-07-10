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
    <v-data-table
      :headers="headers"
      :items="marketData"
      :search="search"
      :disable-pagination="true"
      :hide-default-footer="true"
      item-key="code"
    >
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
      marketData: [],
      coinInfo: [],
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
      ],
    };
  },
  async created() {
    try {
      const minute = {
        min: 1,
      };
      const resCoin = await getCoinInfo();
      this.coinInfo = resCoin.data;
      const { data } = await getMarketInfo(minute);

      let temp = JSON.parse(atob(data));

      for (const key of Object.keys(temp)) {
        temp[key]['code'] = key;
      }

      for (const value of Object.values(temp)) {
        this.coinInfo.forEach(e => {
          if (e.market == value['code']) {
            value['codeName'] = e.korean_name;
          }
        });
        value['AskCount'] = value['AskCount'].toFixed(2);
        value['AskTotal'] = parseInt(value['AskTotal']);
        value['BidCount'] = value['BidCount'].toFixed(2);
        value['BidTotal'] = parseInt(value['BidTotal']);
        this.marketData.push(value);
      }
    } catch (error) {
      console.log(error.response.data);
    }
  },
};
</script>

<style scoped>
.coin-logo {
  width: 0.775rem;
}
</style>
