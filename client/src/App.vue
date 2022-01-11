<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Drawflow from './components/Drawflow.vue'

</script>

<template>
  <Drawflow />
</template>

<style>

*, *::before, *::after {
  box-sizing: border-box;
}

* {
  margin: 0;
}

*,
:after,
:before {
  box-sizing: border-box;
}

html {
  font-size: 62.5%;
}

html, body, #app {
  height: 100%;
}

body {
  font-family: 'Roboto', sans-serif;
  font-size: 1.6rem;
  color: #1b1b1b;
}

input, button, textarea, select {
  font: inherit;
}

</style>
