<template>
  <div>
    <h2>Items</h2>
    <button @click="fetchItems">Refresh</button>
    <ul>
      <li v-for="item in items" :key="item.id">
        {{ item.title }}
        <button @click="deleteItem(item.id)">Delete</button>
      </li>
    </ul>
    <ItemForm @add="createItem" />
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import ItemForm from "./ItemForm.vue";

export default {
  components: { ItemForm },
  computed: {
    ...mapState(["items"])
  },
  methods: {
    ...mapActions(["fetchItems", "deleteItem", "createItem"]),
  },
  mounted() {
    this.fetchItems();
  }
};
</script>
