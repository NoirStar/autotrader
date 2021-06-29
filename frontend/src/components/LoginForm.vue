<template>
  <div>
    <v-card elevation="3" outlined class="mx-auto mt-16" max-width="700">
      <v-card-title class="justify-center">
        <span class="text-h3 my-10">Autotrader</span>
      </v-card-title>
      <v-form @submit.prevent="submitForm">
        <v-container>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field v-model="id" label="아이디" clearable required>
              </v-text-field>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field
                v-model="pw"
                :append-icon="passwordShow ? 'mdi-eye' : 'mdi-eye-off'"
                :type="passwordShow ? 'text' : 'password'"
                @click:append="passwordShow = !passwordShow"
                label="비밀번호"
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row class="my-10" justify="center">
            <v-col cols="7">
              <v-btn
                color="primary"
                large
                block
                :loading="loading"
                :disabled="!id || !pw || loading"
                type="submit"
              >
                로그인
              </v-btn>
            </v-col>
          </v-row>
        </v-container>
      </v-form>
    </v-card>
    <v-snackbar v-model="snackbar" timeout="5000">
      {{ logMessage }}

      <template v-slot:action="{ attrs }">
        <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: '',
      pw: '',
      loading: false,
      snackbar: false,
      passwordShow: false,
      logMessage: '',
    };
  },
  methods: {
    async submitForm() {
      try {
        const userData = {
          id: this.id,
          pw: this.pw,
        };
        this.loading = true;
        await this.$store.dispatch('LOGIN', userData);
        this.loading = false;
        this.$router.push('/main');
      } catch (error) {
        this.logMessage = error.response.data;
        this.snackbar = true;
        this.loading = false;
      } finally {
        this.initForm();
      }
    },
    initForm() {
      this.id = '';
      this.pw = '';
    },
  },
};
</script>

<style></style>
