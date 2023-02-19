import React from 'react'

const TodoForm = () => {
  return (
    <form className='flex my-8 '>
      <input type="text" placeholder='New Task' className='border-2 border-inputgrey border-r-0 bg-lightgrey p-2 flex-auto w-80 rounded-l-lg text-white focus:border-lightpurple outline-none'/>
      <button className='text-white bg-lightpurple border-2 border-lightpurple p-2 flex-auto w-20 rounded-r-lg font-semibold hover:bg-darkpurple'>Add</button>
    </form>
  )
}

export default TodoForm