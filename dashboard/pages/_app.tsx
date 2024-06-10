import '../styles/globals.css';
import type { AppProps } from 'next/app';
import formbricks from '@formbricks/js/website';
import { useRouter } from 'next/router';
import { useEffect } from 'react';
import Layout from '../components/layout/Layout';
import environment from '../environments/environment';

const printHiringMessage = () => {
  // eslint-disable-next-line no-console
  console.log(`                                         
  *   )        (                  (              
' )  /(   ) (  )\\(  (      ) (    )\\ )  (        
 ( )(_)( /( )\\((_)\\))(  ( /( )(  (()/( ))\\ (     
(_(_()))(_)((_)_((_)()\\ )(_)(()\\  ((_)/((_))\\ )  
|_   _((_)_ (_| _(()((_((_)_ ((_) _| (_)) _(_/(  
  | | / _' || | \\ V  V / _' | '_/ _' / -_| ' \\)) 
  |_| \\__,_||_|_|\\_/\\_/\\__,_|_| \\__,_\\___|_||_|  

  WE'RE HIRING REMOTELY IN 🇫🇷, 🇵🇹 and 🇩🇪! 
  
  ---> https://jobs.tailwarden.com <---

  `);
};

if (typeof window !== 'undefined') {
  printHiringMessage();
  formbricks.init({
    environmentId: environment.FORMBRICKS_ENV_ID,
    apiHost: 'https://app.formbricks.com',
  });
}

export default function App({ Component, pageProps }: AppProps) {
  const router = useRouter();

  useEffect(() => {
    // Connect next.js router to Formbricks
    const handleRouteChange = formbricks?.registerRouteChange;
    router.events.on('routeChangeComplete', handleRouteChange);

    return () => {
      router.events.off('routeChangeComplete', handleRouteChange);
    };
  }, []);
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  );
}
