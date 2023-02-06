import '../styles/globals.css';
import type { AppProps } from 'next/app';
import Layout from '../components/layout/Layout';

export default function App({ Component, pageProps }: AppProps) {
  // If there is no display banner flag in the localStorage, create one:
  if (typeof window !== 'undefined' && !localStorage.displayBanner) {
    localStorage.displayBanner = 'true';
  }

  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  );
}
