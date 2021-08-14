import React, { useMemo } from 'react';
import PropTypes from 'prop-types';

import '../assets/stylesheets/button.scss';

const Button = ({ icon, iconPosition, buttonText, size }) => {
  const iconElement = useMemo(
    () => <div className="button__icon">{icon}</div>,
    [size, iconPosition, icon, buttonText]
  );

  return (
    <div className={`button button--${size}`}>
      {iconPosition === 'left' && iconElement}
      {buttonText}
      {iconPosition === 'right' && iconElement}
    </div>
  );
};

Button.propTypes = {
  icon: PropTypes.node,
  iconPosition: PropTypes.oneOf(['left', 'right']),
  buttonText: PropTypes.string.isRequired,
  size: PropTypes.oneOf(['small', 'medium', 'large']),
};

Button.defaultProps = {
  icon: null,
  iconPosition: 'left',
  size: 'medium',
};

export default Button;
