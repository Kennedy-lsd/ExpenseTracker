import axios from 'axios';

const API_URL = 'http://localhost:8080/api';

export const getTasks = async () => {
  const response = await axios.get(`${API_URL}/purchase`);
  return response.data;
};

export const createTask = async (taskData: { title: string; price: string }) => {
  const response = await axios.post(`${API_URL}/purchase`, taskData);
  return response.data;
};

export const deleteTask = async (id: number) => {
  await axios.delete(`${API_URL}/purchase/${id}`);
};
