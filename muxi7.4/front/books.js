import axios from 'axios';

const API_URL = 'http://localhost:8080';

const getAuthHeader = () => {
    const token = localStorage.getItem('token');
    return {
        headers: {
            Authorization: `Bearer ${token}`
        }
    };
};
export const logout = () => {
    localStorage.removeItem('token'); // 清除JWT
    localStorage.removeItem('user');  // 清除用户数据（如果有）
};

export const getAllBooks = async () => {
    const response = await axios.get(`${API_URL}/books`);
    return response.data;
};

export const getBookById = async (id) => {
    const response = await axios.get(`${API_URL}/books/${id}`);
    return response.data;
};

export const addBook = async (book) => {
    const response = await axios.post(`${API_URL}/books`, book, getAuthHeader());
    return response.data;
};

export const updateBook = async (id, book) => {
    const response = await axios.put(`${API_URL}/books/${id}`, book, getAuthHeader());
    return response.data;
};

export const deleteBook = async (id) => {
    const response = await axios.delete(`${API_URL}/books/${id}`, getAuthHeader());
    return response.data;
};