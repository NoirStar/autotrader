<template>
  <div>
    <v-card elevation="3" outlined class="mx-auto mt-16" max-width="700">
      <v-card-title>
        <v-icon left> mdi-chevron-right </v-icon>
        <span class="text-h6">회원가입 정보를 입력해 주세요</span>
      </v-card-title>
      <v-form @submit.prevent="submitForm" ref="form" lazy-validation>
        <v-container>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field
                v-model="id"
                clearable
                label="아이디"
                :rules="idRules"
                :counter="12"
                required
              >
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
                :rules="pwRules"
                label="비밀번호"
                counter
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field
                v-model="pwCopy"
                :append-icon="passwordShow ? 'mdi-eye' : 'mdi-eye-off'"
                :type="passwordShow ? 'text' : 'password'"
                @click:append="passwordShow = !passwordShow"
                :rules="pwCopyRules"
                hint="패스워드가 일치합니다"
                label="비밀번호 확인"
                counter
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field
                v-model="email"
                :rules="emailRules"
                label="이메일"
                clearable
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="9">
              <v-text-field
                :rules="nickRules"
                v-model="nickname"
                label="닉네임"
                clearable
                required
              >
              </v-text-field>
            </v-col>
          </v-row>
          <v-row justify="center">
            <v-col cols="9">
              <v-menu
                ref="menu"
                :close-on-content-click="false"
                transition="scale-transition"
                offset-y
                min-width="auto"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-text-field
                    v-model="birth"
                    :rules="birthRules"
                    label="생년월일"
                    clearable
                    readonly
                    v-bind="attrs"
                    v-on="on"
                  ></v-text-field>
                </template>
                <v-date-picker
                  v-model="birth"
                  :max="
                    new Date(
                      Date.now() - new Date().getTimezoneOffset() * 60000,
                    )
                      .toISOString()
                      .substr(0, 10)
                  "
                  min="1950-01-01"
                  @change="saveDate"
                ></v-date-picker>
              </v-menu>
            </v-col>
          </v-row>
          <v-row class="my-10" justify="center">
            <v-col cols="7">
              <v-btn :disabled="!isFormValid" large block type="submit">
                회원가입
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
import { registerUser } from '@/api/index';
import { validateEmail, validatePassword } from '@/utils/validation';

export default {
  data() {
    return {
      snackbar: false,
      passwordShow: false,
      id: '',
      pw: '',
      pwCopy: '',
      email: '',
      nickname: '',
      birth: '',
      idRules: [
        v => !!v || '아이디를 입력해주세요',
        v => (v && v.length <= 12) || '아이디는 12자 이내여야 합니다',
      ],
      pwRules: [
        v => !!v || '비밀번호를 입력해주세요',
        v => (v && v.length >= 8) || '비밀번호는 8자 이상여야 합니다',
        v =>
          validatePassword(v) ||
          '한글자 이상의 영문자, 숫자, 특수문자를 포함하여야 합니다',
      ],
      pwCopyRules: [
        v => !!v || '비밀번호를 입력해주세요',
        v => v == this.pw || '패스워드가 일치하지 않습니다',
      ],
      emailRules: [
        v => !!v || '이메일을 입력해주세요',
        v => validateEmail(v) || '이메일 형식을 확인해 주세요',
      ],
      birthRules: [v => !!v || '생년월일을 입력해주세요'],
      nickRules: [v => !!v || '닉네임을 입력해주세요'],
      logMessage: '',
    };
  },
  computed: {
    isFormValid() {
      return this.id && this.pw && this.email && this.nickname && this.birth;
    },
  },
  methods: {
    async submitForm() {
      try {
        this.validateForm();
        const userData = {
          id: this.id,
          pw: this.pw,
          email: this.email,
          nickname: this.nickname,
          birth: this.birth,
        };
        const { data } = await registerUser(userData);

        this.logMessage = `${data}님 환영합니다`;
        this.resetForm();
        this.resetValidation();
        this.snackbar = true;
      } catch (error) {
        this.logMessage = error.response.data;
        this.snackbar = true;
      }
    },
    validateForm() {
      this.$refs.form.validate();
    },
    resetForm() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
    saveDate(birth) {
      this.$refs.menu.save(birth);
    },
  },
};
</script>

<style></style>
