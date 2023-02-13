import Head from 'next/head';

function Dashboard() {
  return (
    <div className="relative">
      <Head>
        <title>Dashboard - Komiser</title>
        <meta name="description" content="Dashboard - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <p className="flex items-center gap-2 text-lg font-medium text-black-900">
        Dashboard overview
      </p>
    </div>
  );
}

export default Dashboard;
