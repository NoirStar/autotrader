<template>
  <v-card class="mx-auto mt-16">
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
      :disable-pagination="true"
      :hide-default-footer="true"
      item-key="korean_name"
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
        <template v-if="item.changeRate >= 0">
          <template v-if="item.askBid == 'ASK'">
            <div class="price-up bold price-border-ask">
              {{ item.price }}
            </div>
          </template>
          <template v-else>
            <div class="price-up bold price-border-bid">
              {{ item.price }}
            </div>
          </template>
        </template>
        <template v-else>
          <div class="price-down bold">{{ item.price }}</div>
        </template>
      </template>

      <template v-slot:[`item.changeRate`]="{ item }">
        <template v-if="item.changeRate >= 0">
          <div class="price-up">+{{ item.changeRate }}%</div>
        </template>
        <template v-else>
          <div class="price-down">{{ item.changeRate }}%</div>
        </template>
      </template>

      <template v-slot:[`item.highest52`]="{ item }">
        {{ item.highest52 }}
      </template>
      <template v-slot:[`item.lowest52`]="{ item }">
        {{ item.lowest52 }}
      </template>
      <template v-slot:[`item.accTradePrice`]="{ item }">
        {{ item.accTradePrice }} 백만
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
    makeCodes() {
      return this.$store.state.coinInfo.reduce(
        (acc, cur) => (acc.push(`"${cur.market}"`), acc),
        [],
      );
    },
  },
  computed: {},
  async created() {
    try {
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
                this.$set(e, 'price', result.tp.toLocaleString());
                this.$set(e, 'changePrice', result.scp.toLocaleString());
                this.$set(e, 'changeRate', (result.scr * 100).toFixed(2));
                this.$set(
                  e,
                  'accTradePrice',
                  Math.floor(result.atp24h / 1000000).toLocaleString(),
                );
                this.$set(e, 'askBid', result.ab);
                this.$set(e, 'highest52', result.h52wp.toLocaleString());
                this.$set(e, 'lowest52', result.l52wp.toLocaleString());
              }
            });
          };
          reader.readAsText(data);
        }
      };
      conn.send(
        `[{"ticket":"tree"},{"type":"ticker","codes":[${this.makeCodes()}]},{"format":"SIMPLE"}]`,
      );
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

.price-up {
  color: #f06d6f;
}

.price-down {
  color: #4480da;
}

.bold {
  font-weight: 600;
}

.price-border-ask {
  border: 0.5px solid rgba(0, 0, 0, 0);
  animation: blink-ask 0.5s;
  -webkit-animation: blink-ask 0.5s;
}

.price-border-bid {
  border: 0.5px solid rgba(0, 0, 0, 0);
  animation: blink-bid 0.5s;
  -webkit-animation: blink-bid 0.5s;
}

@keyframes blink-ask {
  0% {
    border: 0.5px solid rgba(240, 109, 111, 0);
  }
  100% {
    border: 0.5px solid rgba(240, 109, 111, 1);
  }
}

@-webkit-keyframes blink-ask {
  0% {
    border: 0.5px solid rgba(240, 109, 111, 0);
  }
  100% {
    border: 0.5px solid rgba(240, 109, 111, 1);
  }
}

@keyframes blink-bid {
  0% {
    border: 0.5px solid rgba(240, 109, 111, 0);
  }
  100% {
    border: 0.5px solid rgba(68, 128, 218, 1);
  }
}

@-webkit-keyframes blink-bid {
  0% {
    border: 0.5px solid rgba(240, 109, 111, 0);
  }
  100% {
    border: 0.5px solid rgba(68, 128, 218, 1);
  }
}
</style>
