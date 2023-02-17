import Head from 'next/head';
import DashboardCloudMap from '../components/dashboard/components/cloud-map/DashboardCloudMap';
import DashboardCostExplorer from '../components/dashboard/components/cost-explorer/DashboardCostExplorer';
import DashboardLayout from '../components/dashboard/components/DashboardLayout';
import DashboardResourcesManager from '../components/dashboard/components/resources-manager/DashboardResourcesManager';
import DashboardTopStats from '../components/dashboard/components/top-stats/DashboardTopStats';
import Grid from '../components/grid/Grid';

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
          <DashboardCloudMap />
          <DashboardResourcesManager />
        </Grid>
        {/* to be un-commented on next release <DashboardCostExplorer /> */}
      </DashboardLayout>
    </>
  );
}

export default Dashboard;
