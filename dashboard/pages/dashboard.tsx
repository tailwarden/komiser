import Head from 'next/head';
import DashboardCloudMapWrapper from '../components/dashboard/components/cloud-map/DashboardCloudMapWrapper';
import DashboardCostExplorer from '../components/dashboard/components/cost-explorer/DashboardCostExplorer';
import DashboardLayout from '../components/dashboard/components/DashboardLayout';
import DashboardResourcesManager from '../components/dashboard/components/resources-manager/DashboardResourcesManager';
import DashboardTopStats from '../components/dashboard/components/top-stats/DashboardTopStats';
import Grid from '../components/grid/Grid';
import DashboardDependencyGraphWrapper from '../components/explorer/DependencyGraphWrapper';

function Dashboard() {
  return (
    <>
      <Head>
        <title>Dashboard - Komiser</title>
        <meta name="description" content="Dashboard - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <DashboardLayout>
        <DashboardTopStats />
        <Grid>
          <DashboardCloudMapWrapper />
          <DashboardResourcesManager />
        </Grid>
        <DashboardCostExplorer />
      </DashboardLayout>
    </>
  );
}

export default Dashboard;
