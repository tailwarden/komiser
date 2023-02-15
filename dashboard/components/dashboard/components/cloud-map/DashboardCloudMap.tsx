import DashboardCloudMapChart from './DashboardCloudMapChart';
import DashboardCloudMapTooltip from './DashboardCloudMapTooltip';
import useCloudMap from './useCloudMap';
import useCloudMapTooltip from './useCloudMapTooltip';

function DashboardCloudMap() {
  const { loading, data, error, fetch } = useCloudMap();
  const { tooltip, setTooltip } = useCloudMapTooltip();

  if (loading) return <>Loading</>;

  if (error) return <>Error loading</>;

  return (
    <>
      <DashboardCloudMapChart regions={data} setTooltip={setTooltip} />
      <DashboardCloudMapTooltip tooltip={tooltip} />
    </>
  );
}

export default DashboardCloudMap;
