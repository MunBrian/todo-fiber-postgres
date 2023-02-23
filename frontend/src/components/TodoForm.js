import { useState, useEffect} from "react"

const TodoForm = ({ task, addTask, editTask }) => {
  const { name } = task
  const [input, setInput] = useState('')
  const [editInput, setEditInput] = useState(name)


  useEffect(() => setEditInput(name), [name])

  const onEditClick = (e) => {
    e.preventDefault()

    editTask(editInput)

    setEditInput("")
  }

  const onAddClick = (e) => {
    e.preventDefault()

    if (!input) {
      alert("please input task")
      return
    }

    addTask(input)


    setInput("")
  }


  
  return (
    editInput ? (
    <form className='flex my-8' onSubmit={onEditClick}>
      <input value={editInput}  onChange={(e) => setEditInput(e.target.value)}  type="text" placeholder='New Task' className='border-2 border-inputgrey border-r-0 bg-lightgrey p-2 flex-auto w-80 rounded-l-lg text-white focus:border-lightpurple outline-none'/>
      <button className='text-white bg-lightpurple border-2 border-lightpurple p-2 flex-auto w-20 rounded-r-lg font-semibold hover:bg-darkpurple'>Edit</button>
    </form >
    ) : (
    <form className='flex my-8' onSubmit={onAddClick}>
      <input value={input}  onChange={(e) => setInput(e.target.value)}  type="text" placeholder='New Task' className='border-2 border-inputgrey border-r-0 bg-lightgrey p-2 flex-auto w-80 rounded-l-lg text-white focus:border-lightpurple outline-none'/>
      <button className='text-white bg-lightpurple border-2 border-lightpurple p-2 flex-auto w-20 rounded-r-lg font-semibold hover:bg-darkpurple'>Add</button>
    </form >
    )
  )
    
}

export default TodoForm