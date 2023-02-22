import React from 'react'
import TodoItem from './TodoItem'

const TodoList = ({tasks}) => {
  return (
      <div>
      <ul>
        {tasks.map(task => (
          <li key={ task.ID }><TodoItem  task={task} /></li>
        ))}
          </ul>
    </div>
  )
}

export default TodoList