import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import api from '../../api';
import { updateAuth } from './slices/appSlice';

const Logout = () => {
  const dispatch = useDispatch();

  useEffect(() => {
    api({
      method: 'PUT',
      url: '/logout',
    }).then((response) => {
      console.log(response.data);
      dispatch(updateAuth({ isLoggedIn: false }));
    });
  });

  return (
    <div className="logout">
      <h1 className="logout__heading">You have successfully logged out</h1>
    </div>
  );
};

export default Logout;
