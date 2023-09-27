import Head from 'next/head';
import DashboardDependencyGraphWrapper from '../components/explorer/dependency-graph/DependencyGraphWrapper';

function Explorer() {
  return (
    <>
      <Head>
        <title>Explorer - Komiser</title>
        <meta name="description" content="Explorer - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <DashboardDependencyGraphWrapper />
    </>
  );
}

export default Explorer;
