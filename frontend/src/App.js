import Card from "./components/Card";
import Footer from "./components/Footer";
import { useState, useEffect } from "react";


function App() {

  const [tasks, setTasks] = useState([])
  const [task, setTask] = useState({})


  useEffect(() => {
    fetchTasks()
  }, [])
  

  //fetch tasks
  const fetchTasks = async () => {
    try {
      const data = await fetch("http://localhost:8000/")
      const tasks = await data.json()
  
      setTasks(tasks)
    } catch (error) {
      console.log(error)
    }
  }


  //fetch task by id
  const fetchTaskById = async (id) => {
    try {
      const data = await fetch(`http://localhost:8000/task/${id}`)
      const task = await data.json()

      setTask(task)

    } catch (error) {
      console.log(error)
    }
  }


  //add task
  const addTask = async (inputData) => {
    const res = await fetch("http://localhost:8000/", {
      method: 'POST', // or 'PUT'
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: inputData
      })
    })

    const data = await res.json()

    setTasks([...tasks, data])
  
  }


  //edit task
  const editTask = async (inputData) => {
    let url = `http://localhost:8000/task/${task.ID}`

    const res = await fetch(url, {
      method: "PATCH",
      headers: {
        "content-type": "application/json"
      },
      body: JSON.stringify({
        name: inputData
      })
    })

    const updatedTask = await res.json()

    console.log(updatedTask)
    
    setTasks(tasks.map(task => task.ID === updatedTask.ID ? {...task, name:updatedTask.name} : task ))
  }


  //delete task
   const deleteTask = async (id) => {
    let url = `http://localhost:8000/task/${id}`

    fetch(url, {
      method: "DELETE",
      headers: {
        'content-type': 'application/json'
      }
    }).then(response => {
      const newTasks = tasks.filter(task => task.ID !== id)
      setTasks(newTasks)
    } )
   }
  
  return (
    <div className="bg-darkblue flex flex-col justify-center h-screen items-center">
      <Card {... {tasks, deleteTask, task, addTask, fetchTaskById, editTask}} />
       <Footer />
    </div>
  );
}

export default App;
