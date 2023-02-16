import DashboardResourcesManagerChart from './DashboardResourcesManagerCard';
import DashboardResourcesManagerError from './DashboardResourcesManagerError';
import DashboardResourcesManagerSkeleton from './DashboardResourcesManagerSkeleton';
import useResourcesManager from './hooks/useResourcesManager';

function DashboardResourcesManager() {
  const { loading, data, error, fetch } = useResourcesManager();

  if (loading) return <DashboardResourcesManagerSkeleton />;

  if (error) return <DashboardResourcesManagerError fetch={fetch} />;

  return <DashboardResourcesManagerChart data={data} />;
}

export default DashboardResourcesManager;
