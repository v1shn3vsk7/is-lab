import './App.css';
import LoginSignup from "./Components/LoginSignup/LoginSignup";
import ProgramInfo from "./Components/ProgramInfo/ProgramInfo";
import Account from "./Components/Account/Account"
import BlockedPage from "./Components/BlockedPage/BlockedPage"
import AdminPage from "./Components/AdminPage/AdminPage";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import SetupPassword from "./Components/SetupPassword/SetupPassword";

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/" element={<div>
                  <ProgramInfo/>
                  <LoginSignup/>
              </div>}>
              </Route>

              <Route path="/account/:user_id" element={
                  <div><Account/></div>
              }>
              </Route>

              <Route path="/blocked-page" element={
                  <div><BlockedPage/></div>
              }>
              </Route>

              <Route path="/admin-page/:user_id" element={
                  <div><AdminPage/></div>
              }>
              </Route>

              <Route path="/setup-password/:user_id" element={
                  <div><SetupPassword/></div>
              }>
              </Route>

          </Routes>
      </Router>

  );
}

export default App;