import './App.css';
import LoginSignup from "./Components/LoginSignup/LoginSignup";
import ProgramInfo from "./Components/ProgramInfo/ProgramInfo";
import Account from "./Components/Account/Account"
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/" element={<div>
                  <ProgramInfo/>
                  <LoginSignup/>
              </div>}></Route>

              <Route path="/account/:user_id" element={
                  <div><Account/></div>
              }></Route>
          </Routes>
      </Router>

  );
}

export default App;