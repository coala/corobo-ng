import React from 'react';
import { Link } from 'react-router-dom';

import '../assets/stylesheets/navbar.scss';

const NavBar = () => {
  return (
    <div className="navbar">
      <div className="navbar__icon">corobo</div>
      <div className="navbar__items">
        <Link className="navbar__items--link" to="/login">
          Sign in
        </Link>
      </div>
    </div>
  );
};

export default NavBar;
