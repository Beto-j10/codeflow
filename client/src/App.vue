<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import Drawflow from './components/Drawflow.vue'

</script>

<template>
  <Drawflow />
</template>

<style>
:root {
  --color-primary: #4ea9ff;

  --border-color: #cacaca;
  --border-color-dark: #c1c1c1;

  --bg-color: #fafafa;
  --bg-color-secondary: #f5f5f5;

  --text-color-primary: #1b1b1b;
  --text-color-secondary: #545454;

  --font-size-base: 1.4rem;
  --font-size-medium: 1.6rem;

  --container-padding: 16px;
}

*,
*::before,
*::after {
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

html,
body,
#app {
  height: 100%;
}

body {
  background-color: var(--bg-color);
  font-family: "Roboto", sans-serif;
  font-size: var(--font-size-medium);
  color: var(--text-color-primary);
}

input,
button,
textarea,
select {
  font: inherit;
}
</style>
