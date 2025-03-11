import {createElement} from 'react';
import {useState} from 'react';

//to do:

//1. get tasks from DB (GET) What path?

//2. implement task complete button?

//3. bring functions: signup, signout, and taskcompletebutton, out of task

function SignOutButton() {
  async function SignOut(ANum) {
    const url = "//localhost:8078/hours/punch"
    const data = {barcode: ANum}
    await fetch(url, {method: "POST", body: JSON.stringify(data)})
    updateWorkerCount()
  }
  
  function HandleClick() {
    let ANum = prompt("Enter A#:")
    SignOut(ANum)
  }

  return (
    <button onClick={HandleClick}>Sign Out</button>
  )
}

const [numWorkers, setNumWorkers] = useState({});
async function updateWorkerCount() {
  count = {}
  const url = "//localhost:8078/hours/working"
  console.log("update worker count")
  const response = await fetch(url, {method: "GET"})
  const data = await response.json()
  console.log(data)
  let count = 0
  data.forEach((task)=> {
    //for each worker, count num of ID's
    if (count.hasOwnProperty(task.task_id)) {
      count[task.task_id] += 1
    } else {
      count[task.task_id] = 1
    }
  })
  setNumWorkers(count)
}

function Task({taskName, taskDescription, taskID}) { 

  async function SignIn(ANum, task) {
    const url = "//localhost:8078/hours/punch"
    const data = {barcode: ANum, task: task}
    console.log(data)
    await fetch(url, {method: "POST", body: JSON.stringify(data)})
    updateWorkerCount(task)
  }

  function TaskCompleteButton() {
    return (
      <div>
        <button>Task Complete</button>
        {SignOut()}
      </div>
    )
  }

  async function handleClick() {
    let input = prompt("Enter your A#:")
    if (input == null) {
      return
    }
    if (input != "") {
      SignIn(input, taskID)
    }
  }

  let count = 0
  if numWorkers.hasOwnProperty(taskID)) {
    count = numWorkers[taskID]
  }

  const taskStyle = {
    border: 'solid',
    margin: 10,
    flexShrink: 3,
    padding: 2
  }
  
  return (
    <div className="task" style={taskStyle} onClick={handleClick} >
      <h2>{taskName}</h2>
      <p>{taskDescription}</p><br/>
      <p>Number of people working on task: {count}</p>
    </div>
  );
}

updateWorkerCount()

//TO DO: finish function
async function getTasks() {
  const url = ""
  const response = await fetch (url, {method: "GET"})
 
  return (
    response
  );
}

export default function Page() {
  const panelStyle = {
    display: 'flex',
    flexWrap: 'wrap'
  }

  //tasks = getTasks();

  //manually entered tasks:
  let taskInfo = [{name: "Plant Potatoes", description: "plant potatoes", id: 1},
                  {name: "Harvest Tomatoes", description: "harvest, weigh, and wash tomatoes", id: 2},
                  {name: "Sweep Farmhouse", description: "deep sweep main room", id: 3},
                  {name: "Weed Garden", description: "wild mustard is growing behind the shed", id: 4},
                  {name: "Harvest Cucumbers", description: "harvest, weight, and wash cucumbers", id: 5},
                  {name: "Feed Cat", description: "and give them some pets :)", id: 6}]

  const tasks = taskInfo.map((taskInfo) => <Task taskName={taskInfo.name} taskDescription={taskInfo.description} taskID={taskInfo.id}></Task>)

  return (
    <div>
      <h1>Tasks:</h1>
      <div className="task-panel" style={panelStyle}>
        {tasks}
      </div>
      <SignOutButton />
    </div>
  );
}
