import DependencyGraph from './dependencygraph';
import useDependencyGraph from './hooks/useDependencyGraph';

function DashboardDependencyGraphWrapper() {
  const { loading, data, error, fetch } = useDependencyGraph();
  return (
    <>
      <DependencyGraph
        loading={loading}
        data={data}
        error={error}
        fetch={fetch}
      />
    </>
  );
}

export default DashboardDependencyGraphWrapper;
