import React, { useState, useEffect } from "react";
import { getTasks, createTask, deleteTask } from "./api/api";
import { Task } from "./types";
import { formatDate } from "./utils/formatDate";
import "bootstrap/dist/css/bootstrap.min.css";

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [newTaskName, setNewTaskName] = useState("");
  const [newCategory, setNewCategory] = useState("");
  const [newTaskAmount, setNewTaskAmount] = useState<string>("");
  const [filterCategory, setFilterCategory] = useState("");
  const [errorMessage, setErrorMessage] = useState<string>("");

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const tasksData = await getTasks(filterCategory);
        setTasks(tasksData);
      } catch (error) {
        console.error("Failed to fetch tasks", error);
      }
    };
    fetchTasks();
  }, [filterCategory]);

  const handleCreateTask = async () => {
    try {
      const newTask = await createTask({
        title: newTaskName,
        price: newTaskAmount,
        category: newCategory,
      });
      setTasks([...tasks, newTask]);
      setNewTaskName("");
      setNewTaskAmount("");
      setNewCategory("");
    } catch (error) {
      console.error("Failed to create task", error);
    }
  };

  const handleDeleteTask = async (id: number) => {
    try {
      await deleteTask(id);
      setTasks(tasks.filter((task) => task.id !== id));
    } catch (error) {
      console.error("Failed to delete task", error);
    }
  };

  return (
    <div className="App bg-dark text-light p-4" style={{ minHeight: "100vh" }}>
      <h1 className="text-success mb-4">Expense Tracker</h1>

      {/* Create Task Form */}
      <div className="mb-4">
        <h2>Create Task</h2>
        <input
          type="text"
          className="form-control mb-2"
          placeholder="Task Name"
          value={newTaskName}
          onChange={(e) => setNewTaskName(e.target.value)}
        />
        <div className="mb-4">
          <label htmlFor="taskAmount" className="form-label">
            Amount
          </label>
          <input
            id="taskAmount"
            type="text"
            className="form-control"
            placeholder="Enter amount"
            value={newTaskAmount}
            onChange={(e) => {
              const value = e.target.value;
              if (!isNaN(Number(value)) && value.trim() !== "") {
                setNewTaskAmount(value);
                setErrorMessage("");
              } else {
                setErrorMessage("Please enter a valid amount");
              }
            }}
          />
          {errorMessage && (
            <div className="alert alert-warning mt-2" role="alert">
              {errorMessage}
            </div>
          )}
        </div>
        <div className="mb-4">
          <label htmlFor="categorySelect" className="form-label">
            Task Category
          </label>
          <select
            id="categorySelect"
            className="form-select mb-2"
            value={newCategory}
            onChange={(e) => setNewCategory(e.target.value)}
          >
            <option value="">Select a category</option>
            <option value="shop">Shop</option>
            <option value="food">Food</option>
            <option value="entertainment">Entertainment</option>
          </select>
        </div>

        <button className="btn btn-success" onClick={handleCreateTask}>
          Add Task
        </button>
      </div>

      {/* Filter by Category */}
      <div className="mb-4">
        <h2>Filter by Category</h2>
        <div className="btn-group" role="group" aria-label="Category Filter">
          <button
            type="button"
            className={`btn ${
              filterCategory === "" ? "btn-primary" : "btn-secondary"
            }`}
            onClick={() => setFilterCategory("")}
          >
            All
          </button>
          <button
            type="button"
            className={`btn ${
              filterCategory === "shop" ? "btn-primary" : "btn-secondary"
            }`}
            onClick={() => setFilterCategory("shop")}
          >
            Shop
          </button>
          <button
            type="button"
            className={`btn ${
              filterCategory === "food" ? "btn-primary" : "btn-secondary"
            }`}
            onClick={() => setFilterCategory("food")}
          >
            Food
          </button>
          <button
            type="button"
            className={`btn ${
              filterCategory === "entertainment"
                ? "btn-primary"
                : "btn-secondary"
            }`}
            onClick={() => setFilterCategory("entertainment")}
          >
            Entertainment
          </button>
        </div>
      </div>

      {/* Task List */}
      <h2>Tasks</h2>
      <ul className="list-group">
        {tasks.map((task) => (
          <li
            key={task.id}
            className="list-group-item d-flex justify-content-between align-items-center bg-dark text-light border-secondary"
          >
            <div>
              <strong>{task.title}</strong> - {task.price}Є on{" "}
              {formatDate(task.date)} in the{" "}
              <em>
                {task.category.charAt(0).toUpperCase() + task.category.slice(1)}
              </em>
            </div>
            <button
              className="btn btn-outline-danger btn-sm"
              onClick={() => handleDeleteTask(task.id)}
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;
