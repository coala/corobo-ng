import React from 'react';
import GitHubIcon from './Icons/GitHubIcon';

import '../assets/stylesheets/login.scss';

const SignIn = () => {
  return (
    <div className="container">
      <div className="login">
        <h1 className="login__heading">Hey there!</h1>
        <p className="login__description">
          To get started, login or signup via your GitHub or GitLab account.
        </p>
        <div className="login__links">
          <div className="login__links__button">
            <div className="login__links__button__icon">
              <GitHubIcon />
            </div>
            Login via GitHub
          </div>
          <div className="login__links__button">
            <div className="login__links__button__icon">
              <GitHubIcon />
            </div>
            Login via GitLab
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
