import LayoutFlow from './graph';
import { ReactFlowData } from './hooks/useDependencyGraph';

export type DashboardDependencyGraphViewProps = {
  data: ReactFlowData | undefined;
};

function DependencyGraphView({ data }: DashboardDependencyGraphViewProps) {
  return (
    <div className={`w-full rounded-lg px-6 py-4 pb-6`}>
      {data && <LayoutFlow data={data} />}
      <div className="flex gap-4 text-xs text-black-300"></div>
    </div>
  );
}

export default DependencyGraphView;
