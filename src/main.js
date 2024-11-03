import Vue from 'vue';
import App from './App.vue';
import store from './store';
import VueRouter from 'vue-router';

import Login from './components/Login.vue';
import ItemList from './components/ItemList.vue';

Vue.config.productionTip = false;

Vue.use(VueRouter);

const routes = [
    { path: "/login", component: Login },
    { path: "/items", component: ItemList },
];

const router = new VueRouter({
    routes,
});

new Vue({
    store,
    router,
    render: (h) => h(App),
}).$mount("#app");
