import { memo } from 'react';
import { ReactFlowData } from '../hooks/useDependencyGraph';
import DependencyGraphError from '../components/DependencyGraphError';
import DependencyGraphSkeleton from '../components/DependencyGraphSkeleton';
import DependencyGraphView from './DependencyGraph';

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
