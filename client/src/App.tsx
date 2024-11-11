import React, { useState, useEffect } from 'react';
import { getTasks, createTask, deleteTask } from './api/api';
import { Task } from './types';
import { formatDate } from './utils/formatDate';

const App: React.FC = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [newTaskName, setNewTaskName] = useState('');
  const [newCategory, setNewCategory] = useState('');

  const [newTaskAmount, setNewTaskAmount] = useState<string>('');

  useEffect(() => {
    fetchTasks();
  }, []);

  const fetchTasks = async () => {
    try {
      const tasksData = await getTasks();
      setTasks(tasksData);
    } catch (error) {
      console.error("Failed to fetch tasks", error);
    }
  };

  const handleCreateTask = async () => {
    try {
      const newTask = await createTask({ title: newTaskName, price: newTaskAmount , category: newCategory });
      setTasks([...tasks, newTask]);
      setNewTaskName('');
      setNewTaskAmount('');
      setNewCategory('')
    } catch (error) {
      console.error("Failed to create task", error);
    }
  };

  const handleDeleteTask = async (id: number) => {
    try {
      await deleteTask(id);
      setTasks(tasks.filter(task => task.id !== id));
    } catch (error) {
      console.error("Failed to delete task", error);
    }
  };

  return (
    <div className="App">
      <h1>Expense Tracker</h1>
      
      <div>
        <h2>Create Task</h2>
        <input 
          type="text" 
          placeholder="Task Name" 
          value={newTaskName} 
          onChange={(e) => setNewTaskName(e.target.value)} 
        />
        <input 
          type="number" 
          placeholder="Amount" 
          value={newTaskAmount} 
          onChange={(e) => setNewTaskAmount(e.target.value)} 
        />
        <input 
          type="text" 
          placeholder="Task Category" 
          value={newCategory} 
          onChange={(e) => setNewCategory(e.target.value)} 
        />
        <button onClick={handleCreateTask}>Add Task</button>
      </div>

      <h2>Tasks</h2>
      <ul>
        {tasks.map(task => (
          <li key={task.id}>
            {task.title} - {task.price}Є on {formatDate(task.date)} in the {task.category}
            <button onClick={() => handleDeleteTask(task.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
