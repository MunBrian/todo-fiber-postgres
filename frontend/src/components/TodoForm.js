import { useState, useEffect} from "react"

const TodoForm = () => {
  const [input, setInput] = useState('')

  const addTask = async (e) => {
    e.preventDefault()

    fetch("http://localhost:8000/", {
      method: 'POST', // or 'PUT'
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: input
      })
    })
      .then(response => response.json())
      .then(window.location.reload())

    setInput("")
  }

  

  return (
    <form className='flex my-8' onSubmit={addTask}>
      <input value={input}  onChange={(e) => setInput(e.target.value)}  type="text" placeholder='New Task' className='border-2 border-inputgrey border-r-0 bg-lightgrey p-2 flex-auto w-80 rounded-l-lg text-white focus:border-lightpurple outline-none'/>
      <button className='text-white bg-lightpurple border-2 border-lightpurple p-2 flex-auto w-20 rounded-r-lg font-semibold hover:bg-darkpurple'>Add</button>
    </form>
  )
}

export default TodoForm