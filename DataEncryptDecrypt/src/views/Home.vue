<template>
  <v-app>
    <v-main>
      <v-container>
        <v-radio-group v-model="encryptionMethod">
          <v-radio value="rsa">RSA</v-radio>
          <v-radio value="aes">AES</v-radio>
        </v-radio-group>

        <v-file-input v-model="selectedFiles" label="Chọn file hoặc thư mục" chips multiple directory
          @change="handleFileSelection"></v-file-input>

        <v-btn @click="encryptFiles" color="primary">Mã hoá</v-btn>

        <v-list>
          <v-list-item v-for="(file, index) in selectedFiles" :key="index">
            <v-list-item-icon>
              <v-icon v-if="file && file.webkitRelativePath && file.webkitRelativePath.includes('/')">mdi-folder</v-icon>
              <v-icon v-else>mdi-file</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ file && file.name }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-list>

        <v-data-table v-if="encryptedText" :headers="[{ text: 'Encrypted Text', value: 'encryptedText' }]"
          :items="[{ encryptedText }]" hide-default-footer></v-data-table>

        <v-data-table v-if="encryptedFile" :headers="[{ text: 'Encrypted File', value: 'encryptedFile' }]"
          :items="[{ encryptedFile }]" hide-default-footer>
          <template v-slot:item="{ item }">
            <v-btn @click="downloadEncryptedFile(item.encryptedFile)" color="primary">Tải về</v-btn>
          </template>
        </v-data-table>
        <v-btn @click="decrypt" color="primary">Giải mã</v-btn>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { createApp } from 'vue'; // Import Vue from the vue package
import { createVuetify } from 'vuetify'; // Import Vuetify from the vuetify package

// Initialize Vue app
const app = createApp({});

// Initialize Vuetify
const vuetify = createVuetify();

// Use Vuetify in the app
app.use(vuetify);
const encryptionMethod = ref('rsa');
const selectedFiles = ref([] as File[]);
const encryptedText = ref('');
const encryptedFile = ref<File | null>(null);

const handleFileSelection = (): void => {
  // Xử lý các file đã chọn
  selectedFiles.value.forEach((file) => {
    // Kiểm tra nếu là thư mục
    if (file && file.webkitRelativePath) {
      console.log('Thư mục được chọn:', file);
    } else {
      console.log('File được chọn:', file);
    }
  });
};

const encryptFiles = (): void => {
  if (encryptionMethod.value === 'rsa') {
    // Mã hoá file bằng RSA
    console.log('Mã hoá file bằng RSA');
    // Gọi hàm mã hoá RSA từ code backend
    // Assign the encrypted text or file response to the respective variables
    encryptedText.value = 'Encrypted text from backend';
    encryptedFile.value = null; // Set to the encrypted file returned from backend
  } else if (encryptionMethod.value === 'aes') {
    // Mã hoá file bằng AES
    console.log('Mã hoá file bằng AES');
    // Gọi hàm mã hoá AES từ code backend
    // Assign the encrypted text or file response to the respective variables
    encryptedText.value = 'Encrypted text from backend';
    encryptedFile.value = null; // Set to the encrypted file returned from backend
  }
};

const downloadEncryptedFile = (file: File): void => {
  // Implement file download logic for the encrypted file
};

const decrypt = (): void => {
  // Implement decryption logic
};
</script>

<style>
/* Thêm CSS tùy chỉnh nếu cần */
</style>
