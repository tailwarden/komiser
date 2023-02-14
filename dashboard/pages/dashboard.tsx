import Head from 'next/head';
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
      </DashboardLayout>
    </>
  );
}

export default Dashboard;
