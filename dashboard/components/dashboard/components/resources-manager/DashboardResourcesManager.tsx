import DashboardResourcesManagerChart from './DashboardResourcesManagerCard';
import useResourcesManager from './hooks/useResourcesManager';

function DashboardResourcesManager() {
  const { loading, data, error, fetch } = useResourcesManager();

  if (loading) return <>Loading</>;

  if (error) return <>Error</>;

  return <DashboardResourcesManagerChart data={data} />;
}

export default DashboardResourcesManager;
