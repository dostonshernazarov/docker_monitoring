import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [statuses, setStatuses] = useState([]);

  useEffect(() => {
    axios.get('http://localhost:8080/status')
      .then(response => {
        setStatuses(response.data);
      })
      .catch(error => {
        console.error('There was an error fetching the statuses!', error);
      });
  }, []);

  return (
    <div className="App">
      <h1>Container Status</h1>
      <table>
        <thead>
          <tr>
            <th>IP Address</th>
            <th>Ping Time</th>
            <th>Last Success</th>
          </tr>
        </thead>
        <tbody>
          {statuses.map(status => (
            <tr key={status.ip_address}>
              <td>{status.ip_address}</td>
              <td>{status.ping_time} ms</td>
              <td>{new Date(status.last_success).toLocaleString()}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default App;