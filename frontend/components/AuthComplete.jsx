import React, { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import PropTypes from 'prop-types';
import axios from 'axios';
import constants from '../constants';

const AuthComplete = ({ provider }) => {
  const location = useLocation();
  const queryParams = new URLSearchParams(location.search);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const code = queryParams.get('code');
    if (code) {
      axios
        .post(`${constants.BACKEND_HOST}/login/${provider}`, {
          code: queryParams.get('code'),
        })
        .then((response) => {
          if (response.data.success) {
            console.log(response.data.token);
          }
        });
    } else {
      setLoading(false);
    }
  });

  return (
    <div>
      {loading ? (
        <h1>Authenticating with {provider}...</h1>
      ) : (
        <h1>Unable to authenticate user, please try again.</h1>
      )}
    </div>
  );
};

AuthComplete.propTypes = {
  provider: PropTypes.oneOf(['github', 'gitlab']).isRequired,
};

export default AuthComplete;
