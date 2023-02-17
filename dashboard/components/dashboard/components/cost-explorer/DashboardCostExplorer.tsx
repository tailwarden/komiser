import DashboardCostExplorerCard from './DashboardCostExplorerCard';
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
    setQueryDate
  } = useCostExplorer();

  if (loading) return <DashboardCostExplorerSkeleton />;

  if (error) return <>Error</>;

  return (
    <DashboardCostExplorerCard
      data={data}
      queryGroup={queryGroup}
      setQueryGroup={setQueryGroup}
      queryGranularity={queryGranularity}
      setQueryGranularity={setQueryGranularity}
      queryDate={queryDate}
      setQueryDate={setQueryDate}
    />
  );
}

export default DashboardCostExplorer;
