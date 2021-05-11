import React, { useEffect } from 'react';
// Componens
import Footer from '../organisms/Footer';
import Header from '../organisms/Header';
import img_log from '../../assets/images/log.png'
import style from '../../assets/scss/style.module.scss'
import { login, logout, selectUser } from '../../common/redux/userSlice'
import { useSelector, useDispatch } from 'react-redux'

const Top: React.FC = () => {
  const user = useSelector(selectUser);
  const dispatch = useDispatch();

  useEffect(() => {
    console.log(process.env.PROJECT_ID)
    dispatch(
      login({
        uid: "test_uid",
        photoUrl: "test_phtho",
        displayName: "test_displayname",
      }))
  }, []);

  return (
    <>
      <Header />
      <div className={style['title']}> トップ-- - ページ </div>
      <img src={img_log} alt="" />
      <div>{user.uid}</div>
      <a href="/login">login</a>
      <Footer />
    </>
  );
};
export default Top;