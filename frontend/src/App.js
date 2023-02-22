import Card from "./components/Card";
import Footer from "./components/Footer";
import { useState, useEffect } from "react";


function App() {

  const [tasks, setTasks] = useState([])


  useEffect(() => {
    fetchTasks()
  }, [])
  

  const fetchTasks = async () => {
    try {
      const data = await fetch("http://localhost:8000/")
      const tasks = await data.json()
  
      setTasks(tasks)
    } catch (error) {
      console.log(error)
    }
  }
  
  

  return (
    <div className="bg-darkblue flex flex-col justify-center h-screen items-center">
      <Card tasks={tasks} />
       <Footer />
    </div>
  );
}

export default App;
