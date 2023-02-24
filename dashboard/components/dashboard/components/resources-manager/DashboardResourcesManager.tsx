import DashboardResourcesManagerChart from './DashboardResourcesManagerCard';
import DashboardResourcesManagerError from './DashboardResourcesManagerError';
import DashboardResourcesManagerSkeleton from './DashboardResourcesManagerSkeleton';
import useResourcesManager from './hooks/useResourcesManager';

function DashboardResourcesManager() {
  const { loading, data, error, fetch, query, setQuery, exclude, setExclude } =
    useResourcesManager();

  if (loading) return <DashboardResourcesManagerSkeleton />;

  if (error) return <DashboardResourcesManagerError fetch={fetch} />;

  return (
    <DashboardResourcesManagerChart
      data={data}
      query={query}
      setQuery={setQuery}
      exclude={exclude}
      setExclude={setExclude}
    />
  );
}

export default DashboardResourcesManager;
