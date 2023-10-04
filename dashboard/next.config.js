/** @type {import('next').NextConfig} */
const withTM = require('next-transpile-modules')(['react-cytoscapejs']);

const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    unoptimized: true
  },
  trailingSlash: true
};

module.exports = withTM(nextConfig);
