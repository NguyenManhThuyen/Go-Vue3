<template>
    <v-main>
      <v-container>
        <h2>Liên hệ</h2>
        <p>
          Xin chào! Đây là trang liên hệ của tôi. Bạn có thể sử dụng các phương thức sau để liên lạc với tôi hoặc xem thông tin cá nhân của tôi trên các trang mạng xã hội.
        </p>
    <h3>Thông tin cá nhân</h3>
    <ul>
      <li><strong>Tên:</strong> [Tên của bạn]</li>
      <li><strong>Email:</strong> [Địa chỉ email của bạn]</li>
      <li><strong>Số điện thoại:</strong> [Số điện thoại của bạn]</li>
    </ul>
  
    <h3>Mạng xã hội</h3>
    <ul>
      <li>
        <strong>Facebook:</strong>
        <a href="[Đường dẫn Facebook của bạn]" target="_blank">[Tên Facebook của bạn]</a>
      </li>
      <li>
        <strong>GitHub:</strong>
        <a href="[Đường dẫn GitHub của bạn]" target="_blank">[Tên GitHub của bạn]</a>
      </li>
    </ul>
  </v-container>
    </v-main>
  </template>
  <script>
  export default {
    name: 'ContactPage'
  }
  </script>
  <style lang="scss" scoped>
  /* Custom styling */
  h2 {
    font-size: 24px;
    font-weight: bold;
  }
  
  h3 {
    font-size: 20px;
    font-weight: bold;
  }
  
  ul {
    margin-bottom: 16px;
  }
  
  li {
    margin-bottom: 8px;
  }
  
  a {
    color: blue;
    text-decoration: underline;
  }
  </style>