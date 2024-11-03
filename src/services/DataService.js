import axios from 'axios';
const API_URL = 'http://localhost:3000';
axios.defaults.headers.common['Authorization'] = `Bearer ${localStorage.getItem("token")}`;

export default {
    login(credentials) {
        return axios.post(`${API_URL}/login`, credentials);
    },
    getItems() {
        return axios.get(`${API_URL}/protected/items`);
    },
    createItem(item) {
        return axios.post(`${API_URL}/protected/items`, item);
    },
    updateItem(id, item) {
        return axios.put(`${API_URL}/protected/items/${id}`, item);
    },
    deleteItem(id) {
        return axios.delete(`${API_URL}/protected/items/${id}`);
    }
};
