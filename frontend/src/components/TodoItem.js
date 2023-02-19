import React from 'react'
import { FaTrash } from 'react-icons/fa'
import { MdEdit } from 'react-icons/md'

const TodoItem = () => {
  return (
      <div className='bg-lightgrey p-3 my-4 rounded-md flex justify-between items-center shadow-md'>
          <div className='flex items-center'>
              <input type="checkbox" className='mr-4 accent-darkpurple' title='Mark as complete' />
                <h2 className='text-white font-semibold'>Go to sleep</h2>
          </div>
          <div className='flex items-center'> 
              <button><MdEdit className='mr-4 text-lg text-lightpurple  hover:text-darkpurple' title='edit' /></button>  
              <button><FaTrash className='text-lightpurple hover:text-darkpurple' title='delete' /></button>
          </div>
    </div>
  )
}

export default TodoItem