/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    unoptimized: true
  },
  productionBrowserSourceMaps: false,
  webpack: (
    config,
    { buildId, dev, isServer, defaultLoaders, nextRuntime, webpack }
  ) => {
    config.plugins.push(
      require('unplugin-auto-import/webpack')({
        include: [/\.[tj]sx?$/, /\.md$/],
        imports: [
          'react',
          {
            'next/router': ['useRouter'],
            classnames: ['classNames']
          },
          {
            from: 'react',
            imports: ['ChangeEvent'],
            type: true
          },
          {
            from: 'next',
            imports: ['NextPage'],
            type: true
          }
        ],
        dirs: ['./services', './utils'],
        defaultExportByFilename: true
      })
    );

    return config;
  }
};

module.exports = nextConfig;
