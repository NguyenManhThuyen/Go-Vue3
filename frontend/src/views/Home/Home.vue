<template>
  <v-app>
    <v-main>
      <v-container>
        <v-radio-group v-model="dataType">
          <v-radio label="Text" value="text"></v-radio>
          <v-radio label="File" value="file"></v-radio>
          <v-radio label="Folder" value="folder"></v-radio>
        </v-radio-group>

        <v-container v-if="dataType === 'text'">
          <v-textarea v-model="inputText" label="Input Text"></v-textarea>
        </v-container>

        <v-container v-if="dataType === 'file' || dataType === 'folder'">
          <v-file-input
            v-model="selectedFiles"
            label="Select File or Folder"
            @change="handleFileChange"
            multiple
            :directory="dataType === 'folder'"
          ></v-file-input>
          <v-list-item v-for="(file, index) in selectedFiles" :key="index">
            <v-list-item-content>
              <v-list-item-title v-text="file.name"></v-list-item-title>
            </v-list-item-content>
              <v-list-item-action>
                <v-icon @click="removeFile(index)">mdi-minus</v-icon>
              </v-list-item-action>
          </v-list-item>

          <v-btn @click="addFile">+</v-btn>
        </v-container>

        <v-radio-group v-model="encryptionAlgorithm">
          <v-radio label="RSA" value="rsa"></v-radio>
          <v-radio label="AES" value="aes"></v-radio>
        </v-radio-group>

        <v-divider></v-divider>

        <v-textarea v-model="encryptedData" label="Encrypted Data"></v-textarea>
        <v-btn @click="encrypt">Encrypt</v-btn>
        <v-btn @click="decrypt">Decrypt</v-btn>

        <v-snackbar v-model="snackbar" :timeout="3000" multi-line>
          {{ snackbarMessage }}
        </v-snackbar>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      dataType: 'text',
      inputText: '',
      selectedFiles: [], // Array to store selected files
      encryptedData: '',
      snackbar: false,
      snackbarMessage: '',
      encryptionAlgorithm: 'rsa'
    }
  },
  methods: {
    handleFileChange(files) {
      this.selectedFile = Array.from(files)
    },
    addFile() {
      // Open file dialog to add files
      const input = document.createElement('input')
      input.type = 'file'
      input.multiple = true
      input.directory = this.dataType === 'folder'
      input.addEventListener('change', (event) => {
        this.selectedFiles = [...this.selectedFiles, ...Array.from(event.target.files)]
      })
      input.click()
    },
    removeFile(index) {
      this.selectedFiles.splice(index, 1)
    },
    encrypt() {
      let dataToEncrypt = ''
      if (this.dataType === 'text') {
        dataToEncrypt = this.inputText
      } else if (this.dataType === 'file') {
        // Process selected file(s)
        // You can use Axios or the Fetch API to make an HTTP request
        // Example using Axios:
        const formData = new FormData()
        for (let i = 0; i < this.selectedFile.length; i++) {
          formData.append('files', this.selectedFile[i])
        }
        // Send formData to the server for encryption
      } else if (this.dataType === 'folder') {
        // Process selected folder and its files
        // You can use Axios or the Fetch API to make an HTTP request
        // Example using Axios:
        const formData = new FormData()
        for (let i = 0; i < this.selectedFile.length; i++) {
          formData.append('files', this.selectedFile[i])
        }
        // Send formData to the server for encryption
      }

      // Send the dataToEncrypt and encryptionAlgorithm to the server for encryption
      // You can use Axios or the Fetch API to make an HTTP request
      // Example using Axios:
      axios
        .post('/encrypt', {
          data: dataToEncrypt,
          algorithm: this.encryptionAlgorithm
        })
        .then(response => {
          this.encryptedData = response.data.encryptedData
          this.snackbarMessage = 'Data encrypted successfully'
          this.snackbar = true
        })
        .catch(error => {
          console.error(error)
          this.snackbarMessage = 'Error encrypting data'
          this.snackbar = true
        })
    },
    decrypt() {
      // Send the encryptedData and encryptionAlgorithm to the server for decryption
      // You can use Axios or the Fetch API to make an HTTP request
      // Example using Axios:
      axios
        .post('/decrypt', {
          data: this.encryptedData,
          algorithm: this.encryptionAlgorithm
        })
        .then(response => {
          // Handle decrypted data
          // Display or process the decrypted data as needed
        })
        .catch(error => {
          console.error(error)
          this.snackbarMessage = 'Error decrypting data'
          this.snackbar = true
        })
    }
  }
}
</script>
