import React from 'react'
import TodoItem from './TodoItem'

const TodoList = ({ tasks, fetchTaskById, deleteTask }) => {  
  return (
      <div>
      <ul>
        {tasks.map(task => (
          <li key={task.ID}><TodoItem { ...{ task, fetchTaskById, deleteTask }} /></li>
        ))}
          </ul>
    </div>
  )
}

export default TodoList