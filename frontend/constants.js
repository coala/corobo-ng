const GITHUB_CLIENT_ID = 'ba86a85db9f2349d1011';
const GITHUB_AUTHORIZE_URL = 'https://github.com/login/oauth/authorize';
const GITHUB_REDIRECT_URI =
  process.env.NODE_ENV === 'production'
    ? 'https://corobo-ng.coala.io'
    : 'http://localhost:8080/login/complete';

const constants = {
  GITHUB_CLIENT_ID,
  GITHUB_AUTHORIZE_URL,
  GITHUB_REDIRECT_URI,
};

export default constants;
