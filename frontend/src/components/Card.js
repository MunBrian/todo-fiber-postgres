import React from 'react'
import TodoForm from './TodoForm'
import TodoList from './TodoList'

const Card = () => {
  return (
      <card className="bg-darkgrey h-fit w-[28rem] p-4 rounded-lg shadow-lg flex flex-col">
          <h1 className='text-white font-bold text-lg'>Simple Todo App</h1>
      <TodoForm />
      <TodoList />
    </card>
  )
}

export default Card