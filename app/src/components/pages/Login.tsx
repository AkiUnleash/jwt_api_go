import * as React from 'react';
import Loginform from '../molecules/LoginForm';
import Footer from '../organisms/Footer';
import Header from '../organisms/Header';

const Login: React.FC = () => {
  return (
    <>
      <Header />
      <Loginform />
      <Footer />
    </>
  );
};
export default Login;