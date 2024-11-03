import Vue from 'vue';
import Vuex from 'vuex';
import DataService from '../services/DataService';

Vue.use(Vuex);

export default new Vuex.Store({
    state: {
        items: [],
        user: null,
        token: localStorage.getItem("token") || null,
    },
    mutations: {
        setItems(state, items) {
            state.items = items;
        },
        setUser(state, user) {
            state.user = user;
        },
        setToken(state, token) {
            state.token = token;
            localStorage.setItem("token", token);
            axios.defaults.headers.common["Authorization"] = `Bearer ${token}`;
        },
        logout(state) {
            state.user = null;
            state.token = null;
            localStorage.removeItem("token");
            delete axios.defaults.headers.common["Authorization"];
        }
    },
    actions: {
        async login({ commit }, credentials) {
            const response = await DataService.login(credentials);
            commit("setToken", response.data.token);
            return response;
        },
        async fetchItems({ commit }) {
            const response = await DataService.getItems();
            commit("setItems", response.data);
        },
        async createItem({ dispatch }, item) {
            await DataService.createItem(item);
            dispatch("fetchItems");
        },
        async updateItem({ dispatch }, { id, item }) {
            await DataService.updateItem(id, item);
            dispatch("fetchItems");
        },
        async deleteItem({ dispatch }, id) {
            await DataService.deleteItem(id);
            dispatch("fetchItems");
        }
    },
});
