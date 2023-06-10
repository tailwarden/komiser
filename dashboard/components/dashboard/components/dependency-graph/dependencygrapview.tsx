import { useEffect, useState } from 'react';
import 'reactflow/dist/style.css';
import { ElkNode } from 'elkjs';
import { Background, ReactFlow } from 'reactflow';
import CustomNode from './nodes/nodes';
import elkLayout from './graph';
import { ReactFlowData } from './hooks/useDependencyGraph';

const nodeTypes = {
  customNode: CustomNode
};

export type DashboardDependencyGraphViewProps = {
  data: ReactFlowData | undefined;
};

function DependencyGraphView({ data }: DashboardDependencyGraphViewProps) {
  const [variant, setVar] = useState<any>('cross');
  const [nodes, setNodes] = useState<any>(data?.nodes);
  const [edges, setEdges] = useState<any>(data?.edges);
  const nodesForFlow = (graph: ElkNode) => [
    ...graph.children!.map(node => ({
      ...data?.nodes.find(n => n.id === node.id),
      position: { x: node.x, y: node.y }
    }))
  ];
  const edgesForFlow = (graph: ElkNode) => graph.edges;

  useEffect(() => {
    elkLayout(data?.nodes, data?.edges).then(graph => {
      setNodes(nodesForFlow(graph));
      setEdges(edgesForFlow(graph));
    });
  }, []);

  if (nodes === null) {
    return <></>;
  }
  return (
    <div className={`w-full rounded-lg bg-white px-6 py-4 pb-6`}>
      <div className="-mx-6 flex items-center justify-between border-b border-black-200/40 px-6 pb-4">
        <div>
          <p className="text-sm font-semibold text-black-900">
            Dependency Graph
          </p>
          <div className="mt-1"></div>
          <p className="text-xs text-black-300">
            Analyze account resource associations
          </p>
        </div>
        <div className="flex h-[60px] items-center"></div>
      </div>
      <div className="mt-8"></div>
      <div className="h-[70vh]">
        <ReactFlow nodes={nodes} edges={edges} nodeTypes={nodeTypes}>
          <Background variant={variant} />
        </ReactFlow>
      </div>
      <div className="flex gap-4 text-xs text-black-300"></div>
    </div>
  );
}

export default DependencyGraphView;
