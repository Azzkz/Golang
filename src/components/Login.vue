<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="login">
      <input v-model="username" placeholder="Username" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Login</button>
    </form>
    <p v-if="error">{{ error }}</p>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  data() {
    return {
      username: "",
      password: "",
      error: null
    };
  },
  methods: {
    ...mapActions(["login"]),
    async login() {
      try {
        await this.login({ username: this.username, password: this.password });
        this.$router.push("/items");
      } catch {
        this.error = "Invalid credentials";
      }
    }
  }
};
</script>
