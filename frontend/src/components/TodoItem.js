import React from 'react'
import { FaTrash } from 'react-icons/fa'
import { MdEdit } from 'react-icons/md'

const TodoItem = ({ task }) => {
  
  const deleteTask = (id) => {
    let url = `http://localhost:8000/task/${id}`

    fetch(url, {
      method: "DELETE",
      headers: {
        'content-type': 'application/json'
      }
    }).then(response => response.json()).then(window.location.reload())
  }


  return (
      <div className='bg-lightgrey p-3 my-4 rounded-md flex justify-between items-center shadow-md'>
          <div className='flex items-center'>
              <input type="checkbox" className='mr-4 accent-darkpurple' title='Mark as complete' />
        <h2 className='text-white font-semibold'>{task.name}</h2>
          </div>
          <div className='flex items-center'> 
              <button><MdEdit className='mr-4 text-lg text-lightpurple  hover:text-darkpurple' title='edit' /></button>  
              <button onClick={() => deleteTask(task.ID)}><FaTrash className='text-lightpurple hover:text-darkpurple' title='delete' /></button>
          </div>
    </div>
  )
}

export default TodoItem