import Head from 'next/head';
import DashboardCloudMap from '../components/dashboard/components/cloud-map/DashboardCloudMap';
import DashboardLayout from '../components/dashboard/components/DashboardLayout';
import DashboardTopStats from '../components/dashboard/components/top-stats/DashboardTopStats';

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
        <div className="grid grid-cols-1 gap-8 lg:grid-cols-2">
          <DashboardCloudMap />
        </div>
      </DashboardLayout>
    </>
  );
}

export default Dashboard;
