<template>
    <v-main>
        <v-container>
            <h1>Các loại mã hoá dữ liệu</h1>
            <h1>Có 4 kiểu mã hoá chính</h1>
            <h2>Mã hoá cổ điển</h2>
            <h2>Khoá đối xứng và bất đối xứng</h2>
            <p>
                Khoá đối xứng (symmetric key) và khoá bất đối xứng (asymmetric key) là hai khái niệm quan trọng trong mật mã
                học.
            </p>
            <h3>Khoá đối xứng</h3>
            <p>
                Khoá đối xứng sử dụng cùng một khoá để mã hóa và giải mã dữ liệu. Khoá này được gọi là khoá bí mật (secret
                key) và phải được trao đổi một cách an toàn giữa người gửi và người nhận trước khi truyền dữ liệu. Khoá đối
                xứng thường được sử dụng trong các thuật toán mã hóa như AES (Advanced Encryption Standard) và DES (Data
                Encryption Standard).
            </p>

            <h3>Khoá bất đối xứng</h3>
            <p>
                Khoá bất đối xứng sử dụng một cặp khoá, gồm khoá công khai (public key) và khoá bí mật (private key). Khoá
                công khai được chia sẻ công khai và có thể được sử dụng để mã hóa dữ liệu, trong khi khoá bí mật chỉ được
                giữ bí mật và dùng để giải mã. Khoá bất đối xứng thường được sử dụng trong các thuật toán mã hóa như RSA.
            </p>
        </v-container>
    </v-main>
</template>
<script>
export default {
    name: 'ExercisePage'
}
</script>
<style lang="scss" scoped>
/* Custom styling */
$primary-color: #FF9800;
$secondary-color: #2196F3;

h2 {
    color: $primary-color;
    font-size: 24px;
    font-weight: bold;
}

h3 {
    color: $secondary-color;
    font-size: 20px;
    font-weight: bold;
}

p {
    font-size: 16px;
    margin-bottom: 16px;
}</style>