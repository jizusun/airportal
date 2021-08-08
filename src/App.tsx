import React from 'react';
import './App.css';
import Amplify, { API } from "aws-amplify";
import awsExports from "./aws-exports";
import { useState, useEffect} from 'react';
Amplify.configure(awsExports)


const App = () => {
  const [names, setNames] = useState([])
  useEffect(() => {
    fetchNames()
  }, [])

  async function fetchNames(){
    const response = await API.get('goapi', '/users', null);
  }

  return (
    <div style={styles.container as React.CSSProperties}>
      <h2>Hello Airportal</h2>
    </div>
  );
}
const styles = {
  container: { width: 400, margin: '0 auto', display: 'flex', flexDirection: 'column', justifyContent: 'center', padding: 20 },
  todo: {  marginBottom: 15 },
  input: { border: 'none', backgroundColor: '#ddd', marginBottom: 10, padding: 8, fontSize: 18 },
  todoName: { fontSize: 20, fontWeight: 'bold' },
  todoDescription: { marginBottom: 0 },
  button: { backgroundColor: 'black', color: 'white', outline: 'none', fontSize: 18, padding: '12px 0px' }
}



export default App;
