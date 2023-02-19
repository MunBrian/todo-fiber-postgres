import React from 'react'
import TodoItem from './TodoItem'

const TodoList = () => {
  return (
      <div>
          <ul>     
              <li><TodoItem /></li>
              <li><TodoItem /></li>
              <li><TodoItem /></li>
              <li><TodoItem /></li>
          </ul>
    </div>
  )
}

export default TodoList