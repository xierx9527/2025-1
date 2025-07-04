import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const login = async (username, password) => {
    try {
        const response = await axios.post(`${API_URL}/login`, {
            username,
            password
        });
        return response.data;
    } catch (error) {
        throw error.response.data;
    }
};

export const getAuthToken = () => {
    return localStorage.getItem('token');
};
export const logout = () => {
    localStorage.removeItem('token');
}
export const setAuthToken = (token) => {
    localStorage.setItem('token', token);
};

export const removeAuthToken = () => {
    localStorage.removeItem('token');
};