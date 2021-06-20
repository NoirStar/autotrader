<template>
  <v-card elevation="3" outlined class="mx-auto mt-6" max-width="700">
    <v-card-title>
      <v-icon large left> mdi-twitter </v-icon>
      <span class="text-h5">로그인</span>
    </v-card-title>
    <v-form @submit.prevent="submitForm">
      <v-container>
        <v-row>
          <v-col>
            <v-text-field v-model="username" label="아이디" required>
            </v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field v-model="password" label="비밀번호" required>
            </v-text-field>
          </v-col>
        </v-row>
        <v-btn
          :self-align="center"
          :disabled="!isUsernameValid || !password"
          type="submit"
        >
          로그인
        </v-btn>
      </v-container>
    </v-form>
  </v-card>
</template>

<script>
import { validateEmail } from '@/utils/validation';

export default {
  data() {
    return {
      username: '',
      password: '',
      logMessage: '',
    };
  },
  computed: {
    isUsernameValid() {
      return validateEmail(this.username);
    },
  },
  methods: {
    async submitForm() {
      try {
        const userData = {
          username: this.username,
          password: this.password,
        };
        await this.$store.dispatch('LOGIN', userData);
        this.$router.push('/main');
      } catch (error) {
        console.log(error.response);
      } finally {
        this.initForm();
      }
    },
    initForm() {
      this.username = '';
      this.password = '';
    },
  },
};
</script>

<style></style>
