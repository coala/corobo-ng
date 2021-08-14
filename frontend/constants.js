const IS_PROD = process.env.NODE_ENV === 'production';
const HOST = IS_PROD ? 'https://corobo-ng.coala.io' : 'http://localhost:3000';
const BACKEND_HOST = IS_PROD
  ? 'https://api.corobo-ng.coala.io'
  : 'http://localhost:8080';
const GITHUB_CLIENT_ID = 'ba86a85db9f2349d1011';
const GITHUB_AUTHORIZE_URL = 'https://github.com/login/oauth/authorize';
const GITHUB_REDIRECT_URI = `${HOST}/login/github/complete`;
const GITLAB_CLIENT_ID =
  '9a8f450cd5ab0283eaa51b6a384d07a9c11d17296fdaaedda4fc83e62e9cfb8e';
const GITLAB_AUTHORIZE_URL = 'https://gitlab.com/oauth/authorize';
const GITLAB_REDIRECT_URI = `${HOST}/login/gitlab/complete`;

const constants = {
  IS_PROD,
  HOST,
  BACKEND_HOST,
  GITHUB_CLIENT_ID,
  GITHUB_AUTHORIZE_URL,
  GITHUB_REDIRECT_URI,
  GITLAB_CLIENT_ID,
  GITLAB_AUTHORIZE_URL,
  GITLAB_REDIRECT_URI,
};

export default constants;
