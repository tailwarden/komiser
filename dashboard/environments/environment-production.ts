const environment = {
  production: true,
  API_URL: process.env.NEXT_PUBLIC_API_URL
    ? process.env.NEXT_PUBLIC_API_URL
    : '',
  GA_TRACKING_ID: 'G-9HF3HT6S6W'
};

export default environment;
