<template>
  <div id="app">
    <router-link to="/login">Login</router-link>
    <router-link to="/items">Items</router-link>
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>
