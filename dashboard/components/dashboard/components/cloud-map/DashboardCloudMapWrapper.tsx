import DashboardCloudMap from './DashboardCloudMap';
import useCloudMap from './hooks/useCloudMap';

function DashboardCloudMapWrapper() {
  const { loading, data, error, fetch } = useCloudMap();

  return (
    <DashboardCloudMap
      loading={loading}
      data={data}
      error={error}
      fetch={fetch}
    />
  );
}

export default DashboardCloudMapWrapper;
