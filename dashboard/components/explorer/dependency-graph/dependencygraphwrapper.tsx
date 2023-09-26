import DependencyGraphLoader from './DependencyGraphLoader';
import useDependencyGraph from './hooks/useDependencyGraph';

function DependencyGraphWrapper() {
  const { loading, data, error, fetch } = useDependencyGraph();
  return (
    <>
      <div className="flex h-[calc(100vh-145px)] w-full flex-col">
        <div className="flex flex-row justify-between gap-2">
          <p className="text-lg font-medium text-black-900">Graph View</p>
          <div className="absolute -top-1 right-24 border-x border-b border-black-170 bg-white p-2 text-sm">
            Filters
          </div>
        </div>
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
