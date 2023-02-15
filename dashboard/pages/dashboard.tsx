import Head from 'next/head';
import DashboardCloudMap from '../components/dashboard/components/cloud-map/DashboardCloudMap';
import DashboardLayout from '../components/dashboard/components/DashboardLayout';
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
        </Grid>
      </DashboardLayout>
    </>
  );
}

export default Dashboard;
