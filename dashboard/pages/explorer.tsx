import Head from 'next/head';
import DashboardDependencyGraphWrapper from '@components/explorer/dependency-graph/multi-resource-dependency-graph/DependencyGraphWrapper';
import { useRouter } from 'next/router';
import SingleDependencyGraphWrapper from '@components/explorer/dependency-graph/single-resource-dependency-graph/SingleDependencyGraphWrapper';

function Explorer() {
  const router = useRouter();
  const { resourceId } = router.query;
  return (
    <>
      <Head>
        <title>Explorer - Komiser</title>
        <meta name="description" content="Explorer - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      {resourceId ? (
        <SingleDependencyGraphWrapper
          resourceId={resourceId as string}
          isInExplorer={true}
        />
      ) : (
        <DashboardDependencyGraphWrapper />
      )}
    </>
  );
}

export default Explorer;
