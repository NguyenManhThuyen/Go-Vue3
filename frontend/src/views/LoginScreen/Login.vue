<template>
  <v-container fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="4">
        <v-card>
          <v-card-title class="text-center">Đăng nhập</v-card-title>
          <v-card-text>
            <v-form @submit="login">
              <v-text-field
                v-model="username"
                label="Tên đăng nhập"
                required
              ></v-text-field>
              <v-text-field
                v-model="password"
                label="Mật khẩu"
                type="password"
                required
              ></v-text-field>
              <v-checkbox v-model="isAdmin" label="Admin"></v-checkbox>
              <v-checkbox v-model="isUser" label="Người dùng"></v-checkbox>
              <v-btn type="submit" color="primary">Đăng nhập</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  data() {
    return {
      username: '',
      password: '',
      isAdmin: false,
      isUser: false
    };
  },
  methods: {
    login(event: { preventDefault: () => void; }) {
      event.preventDefault(); // Ngăn chặn hành vi mặc định của form
      if (this.isAdmin) {
        this.$router.push('/admin');
      } else if (this.isUser) {
        this.$router.push('/users');
      } else {
        this.$router.push('/schoolList'); // Chuyển hướng đến "/schoolList"
      }
}

  },
  watch: {
    isAdmin(val) {
      if (val && this.isUser) {
        this.isUser = false;
      }
    },
    isUser(val) {
      if (val && this.isAdmin) {
        this.isAdmin = false;
      }
    }
  }
});
</script>
