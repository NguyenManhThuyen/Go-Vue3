<template>

    <v-app-bar flat>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <v-app-bar-title>
        <v-icon>{{ iconPath }}</v-icon>
      </v-app-bar-title>
      <v-spacer></v-spacer>
      <v-btn text to="/">Home</v-btn>
      <v-btn text to="/contact">Contact</v-btn>
      <v-btn text to="/exercise">Exercise</v-btn>
    </v-app-bar>
</template>

<script>
// Import statements
</script>

<style lang="scss">
// Custom SCSS styles
$primary-color: #FF9800;
$secondary-color: #2196F3;


// Custom styling
.v-app {
  .v-app-bar {
    background-color: $primary-color;

    .v-btn {
      color: pink;
      text-transform: uppercase;
      font-weight: bold;
    }
  }

  .v-main {
    padding-top: 20px;
  }
}
</style>
