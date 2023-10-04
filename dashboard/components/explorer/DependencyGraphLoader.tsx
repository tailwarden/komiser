import { memo } from 'react';
import DependencyGraphError from './DependencyGraphError';
import DependencyGraphSkeleton from './DependencyGraphSkeleton';
import DependencyGraphView from './DependencyGraph';
import { ReactFlowData } from './hooks/useDependencyGraph';

export type DependencyGraphLoaderProps = {
  loading: boolean;
  data: ReactFlowData | undefined;
  error: boolean;
  fetch: () => void;
};

function DependencyGraphLoader({
  loading,
  data,
  error,
  fetch
}: DependencyGraphLoaderProps) {
  if (loading) return <DependencyGraphSkeleton />;

  if (error) return <DependencyGraphError fetch={fetch} />;

  if (data && !loading) return <DependencyGraphView data={data} />;

  return null;
}

export default memo(DependencyGraphLoader);
