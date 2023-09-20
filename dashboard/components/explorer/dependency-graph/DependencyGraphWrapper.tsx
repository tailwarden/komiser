import DependencyGraphLoader from './DependencyGraphLoader';
import useDependencyGraph from './hooks/useDependencyGraph';

function DependencyGraphWrapper() {
  const { loading, data, error, fetch } = useDependencyGraph();
  return (
    <>
      <div className="flex flex-col gap-6">
        <p className="flex items-center gap-2 text-lg font-medium text-black-900">
          Graph View
        </p>
        <DependencyGraphLoader
          loading={loading}
          data={data}
          error={error}
          fetch={fetch}
        />
      </div>
    </>
  );
}

export default DependencyGraphWrapper;
