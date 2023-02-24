import DashboardCostExplorerCard from './DashboardCostExplorerCard';
import DashboardCostExplorerError from './DashboardCostExplorerError';
import DashboardCostExplorerSkeleton from './DashboardCostExplorerSkeleton';
import useCostExplorer from './hooks/useCostExplorer';

function DashboardCostExplorer() {
  const {
    loading,
    data,
    error,
    fetch,
    queryGroup,
    setQueryGroup,
    queryGranularity,
    setQueryGranularity,
    queryDate,
    setQueryDate,
    exclude,
    setExclude
  } = useCostExplorer();

  if (loading) return <DashboardCostExplorerSkeleton />;

  if (error) return <DashboardCostExplorerError fetch={fetch} />;

  return (
    <DashboardCostExplorerCard
      data={data}
      queryGroup={queryGroup}
      setQueryGroup={setQueryGroup}
      queryGranularity={queryGranularity}
      setQueryGranularity={setQueryGranularity}
      queryDate={queryDate}
      setQueryDate={setQueryDate}
      exclude={exclude}
      setExclude={setExclude}
    />
  );
}

export default DashboardCostExplorer;
