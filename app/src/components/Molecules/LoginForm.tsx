import React, { useState } from 'react';
import Textfield from '../atoms/Textfield';
import { browserHistory } from "../../history"

const Loginform: React.FC = () => {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  return (
    <form>

      <Textfield
        type="text"
        placeholder="info@example.com"
        id={"email"}
        value={email}
        setValue={setEmail}
        label="Login ID" />

      <Textfield
        type="password"
        placeholder="password"
        id={"password"}
        value={password}
        setValue={setPassword}
        label="Password" />

      <button
        onClick={(e: React.FormEvent) => {
          e.preventDefault();
          console.log(email, password);
          browserHistory.push('/')
        }}>Submit</button>

    </form>
  );
};
export default Loginform;