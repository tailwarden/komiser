import DashboardCloudMapCard from './DashboardCloudMapCard';
import DashboardCloudMapError from './DashboardCloudMapError';
import DashboardCloudMapSkeleton from './DashboardCloudMapSkeleton';
import { DashboardCloudMapRegions } from './hooks/useCloudMap';

export type DashboardCloudMapProps = {
  loading: boolean;
  data: DashboardCloudMapRegions | undefined;
  error: boolean;
  fetch: () => void;
};

function DashboardCloudMap({
  loading,
  data,
  error,
  fetch
}: DashboardCloudMapProps) {
  if (loading) return <DashboardCloudMapSkeleton />;

  if (error) return <DashboardCloudMapError fetch={fetch} />;

  return <DashboardCloudMapCard data={data} />;
}

export default DashboardCloudMap;
