import DashboardResourcesManagerChart from './DashboardResourcesManagerCard';
import DashboardResourcesManagerSkeleton from './DashboardResourcesManagerSkeleton';
import useResourcesManager from './hooks/useResourcesManager';

function DashboardResourcesManager() {
  const { loading, data, error, fetch } = useResourcesManager();

  if (loading) return <DashboardResourcesManagerSkeleton />;

  if (error) return <>Error</>;

  return <DashboardResourcesManagerChart data={data} />;
}

export default DashboardResourcesManager;
