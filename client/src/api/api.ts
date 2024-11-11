import axios from "axios";

const API_URL = "http://localhost:8080/api";

export const getTasks = async (filterCategory: string) => {
  let url = `${API_URL}/purchase`; // Default URL for fetching all tasks
  if (filterCategory) {
    url = `${API_URL}/purchase?category=${filterCategory}`; // Append filterCategory if provided
  }

  console.log("Request URL: ", url); // Log URL to check if it's correct

  const response = await axios.get(url);
  return response.data;
};

export const createTask = async (taskData: {
  title: string;
  price: string;
  category: string;
}) => {
  const response = await axios.post(`${API_URL}/purchase`, taskData);
  return response.data;
};

export const deleteTask = async (id: number) => {
  await axios.delete(`${API_URL}/purchase/${id}`);
};
