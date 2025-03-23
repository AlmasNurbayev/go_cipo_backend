/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        hostname: 'cipo_nginx',
        //protocol: '',
        //port: '3201',
      },
      // {
      //   protocol: 'http',
      //   hostname: 'localhost**',
      //   pathname: '/static/news_images/**',
      //   port: '3201', // если не HTTPS
      // },
      // {
      //   protocol: 'http',
      //   hostname: 'localhost',
      //   pathname: '/product_images/**',
      // },
      // {
      //   protocol: 'http',
      //   hostname: 'localhost',
      //   pathname: '/static/product_images/**',
      // },
      // {
      //   protocol: 'http',
      //   hostname: 'localhost',
      //   pathname: '/store_images/**',
      // },
      // {
      //   protocol: 'http',
      //   hostname: 'localhost',
      //   pathname: '/static/store_images/**',
      // },
      // {
      //   protocol: 'https',
      //   hostname: 'cdn.dummyjson.com',
      //   pathname: '**',
      // },
    ],
  },
  // ниже настройка нужна для корректной работы winston
  webpack: (config) => {
    config.resolve.fallback = { fs: false };

    return config;
  },
};

export default nextConfig;
