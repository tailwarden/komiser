import { memo } from 'react';
import DependencyGraphError from '../components/DependencyGraphError';
import DependencyGraphSkeleton from '../components/DependencyGraphSkeleton';
import { ReactFlowData } from '../hooks/useDependencyGraph';
import SingleDependencyGraphView from './SingleDependencyGraph';

export type SingleDependencyGraphLoaderProps = {
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
}: SingleDependencyGraphLoaderProps) {
  if (loading) return <DependencyGraphSkeleton />;

  if (error) return <DependencyGraphError fetch={fetch} />;

  if (data && !loading) return <SingleDependencyGraphView data={data} />;

  return null;
}

export default memo(DependencyGraphLoader);
