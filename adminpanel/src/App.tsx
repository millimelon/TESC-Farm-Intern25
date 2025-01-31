import React from "react";
import DataTable from "./components/dataTable";
import logo from "./logo.svg";
import "./App.css";

const jsondata = [
  {
    ID: 1,
    CreatedAt: "2025-01-23T13:41:50.7536711-08:00",
    UpdatedAt: "2025-01-23T13:41:50.7536711-08:00",
    DeletedAt: null,
    start: "2025-01-22T20:00:00Z",
    duration: 0.18,
    department: "Orchard",
    task: "Picking",
    worker: {
      ID: 1,
      CreatedAt: "2025-01-23T13:26:52.517103547-08:00",
      UpdatedAt: "2025-01-23T13:26:52.517103547-08:00",
      DeletedAt: null,
      name: "Austin Strayer",
      barcode: "69420",
      position: "Harvest Lord",
    },
    worker_id: 1,
    harvest: {
      ID: 1,
      CreatedAt: "2025-01-23T13:37:11.095535553-08:00",
      UpdatedAt: "2025-01-23T13:37:11.095535553-08:00",
      DeletedAt: null,
      weight: 100,
      bin: "123",
      crop: {
        ID: 1,
        CreatedAt: "2025-01-23T13:36:16.436928688-08:00",
        UpdatedAt: "2025-01-23T13:36:16.436928688-08:00",
        DeletedAt: null,
        name: "Honey Crisp Apples",
        season: "Fall",
        type: "Fruit",
      },
      crop_id: 1,
    },
    harvest_id: 1,
  },
];

const coldefs = [
  { field: "worker_id" },
  { field: "crop_id" },
  { field: "start" },
  { field: "duration" },
];

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <main id="main">
        <DataTable rowData={jsondata} colDefs={coldefs} ID={0}></DataTable>
      </main>
    </div>
  );
}

export default App;
