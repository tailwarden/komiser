import DependencyGraphError from './dependencygrapherror';
import DependencyGraphSkeleton from './dependencygraphskeleton';
import DependencyGraphView from './dependencygrapview';
import { ReactFlowData } from './hooks/useDependencyGraph';

export type DashboardDependencyGraphProps = {
  loading: boolean;
  data: ReactFlowData | undefined;
  error: boolean;
  fetch: () => void;
};

function DependencyGraph({
  loading,
  data,
  error,
  fetch
}: DashboardDependencyGraphProps) {
  if (loading) return <DependencyGraphSkeleton />;

  if (error) return <DependencyGraphError fetch={fetch} />;

  return <DependencyGraphView data={data} />;
}

export default DependencyGraph;
