import { useContext } from 'react';
import GlobalAppContext from '../../../layout/context/GlobalAppContext';
import DashboardTopStatsCards from './DashboardTopStatsCards';
import DashboardTopStatsError from './DashboardTopStatsError';
import DashboardTopStatsSkeleton from './DashboardTopStatsSkeleton';

function DashboardTopStats() {
  const { loading, data, error, fetch } = useContext(GlobalAppContext);

  if (loading) return <DashboardTopStatsSkeleton />;

  if (error) return <DashboardTopStatsError fetch={fetch} />;

  return <DashboardTopStatsCards data={data} />;
}

export default DashboardTopStats;
