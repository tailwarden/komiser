import DashboardCloudMapCard from './DashboardCloudMapCard';
import DashboardCloudMapError from './DashboardCloudMapError';
import DashboardCloudMapSkeleton from './DashboardCloudMapSkeleton';
import useCloudMap from './hooks/useCloudMap';

function DashboardCloudMap() {
  const { loading, data, error, fetch } = useCloudMap();

  if (loading) return <DashboardCloudMapSkeleton />;

  if (error) return <DashboardCloudMapError fetch={fetch} />;

  return <DashboardCloudMapCard data={data} />;
}

export default DashboardCloudMap;
