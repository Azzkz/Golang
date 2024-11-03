<template>
  <div>
    <h3>Add New Item</h3>
    <form @submit.prevent="addItem">
      <input v-model="title" placeholder="Title" required />
      <button type="submit">Add Item</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      items: [],
      loading: false,
      error: null,
    };
  },
  methods: {
    async fetchItems() {
      this.loading = true;
      try {
        const response = await DataService.getItems();
        this.items = response.data;
      } catch (error) {
        this.error = "Failed to fetch items";
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>
