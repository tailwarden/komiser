import '../styles/globals.css';
import type { AppProps } from 'next/app';
import Layout from '../components/layout/Layout';

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
}

export default function App({ Component, pageProps }: AppProps) {
  return (
    <Layout>
      <Component {...pageProps} />
    </Layout>
  );
}
