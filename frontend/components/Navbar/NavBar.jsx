import React from 'react';
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { getIsLoggedIn } from '../../pages/AuthPages/slices/selectors';

import './stylesheets/navbar.scss';

const NavBar = () => {
  const isLoggedIn = useSelector(getIsLoggedIn);

  return (
    <div className="navbar">
      <Link className="navbar__icon" to="/">
        corobo
      </Link>
      <div className="navbar__items">
        <Link
          className="navbar__items--link"
          to={isLoggedIn ? '/logout' : '/login'}
        >
          {isLoggedIn ? 'Sign out' : 'Sign in'}
        </Link>
      </div>
    </div>
  );
};

export default NavBar;
