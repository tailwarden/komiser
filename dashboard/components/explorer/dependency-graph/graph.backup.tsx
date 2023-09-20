/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useCallback, useEffect } from 'react';
import dagre from 'dagre';
import ReactFlow, {
  Controls,
  ReactFlowProvider,
  addEdge,
  Connection,
  CoordinateExtent,
  Position,
  useNodesState,
  useEdgesState,
  MarkerType,
  EdgeMarker,
  Panel,
  useReactFlow,
  Background,
  NodeTypes,
  EdgeTypes
} from 'reactflow';
import { ReactFlowData } from './hooks/useDependencyGraph';
import Button from '../../button/Button';
import CustomNode from './nodes/nodes';
import 'reactflow/dist/base.css';
import 'reactflow/dist/style.css';
import FloatingEdge from './nodes/FloatingEdge';
import FloatingConnectionLine from './nodes/FloatingConnectionLine';

const dagreGraph = new dagre.graphlib.Graph();
dagreGraph.setDefaultEdgeLabel(() => ({}));

const nodeExtent: CoordinateExtent = [
  [0, 0],
  [1000, 1000]
];

// In order to keep this example simple the node width and height are hardcoded.
// In a real world app you would use the correct width and height values of
// const nodes = useStoreState(state => state.nodes) and then node.__rf.width, node.__rf.height

type LayoutFlowProps = {
  data: ReactFlowData;
};

const nodeTypes = {
  customNode: CustomNode
} as NodeTypes;

const edgeTypes = {
  floating: FloatingEdge
} as EdgeTypes;

const LayoutFlow = ({ data }: LayoutFlowProps) => {
  const [nodes, setNodes, onNodesChange] = useNodesState(data.nodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(data.edges);
  const { fitView } = useReactFlow();

  const onConnect = useCallback(
    (connection: Connection) => {
      setEdges(eds => addEdge(connection, eds));
    },
    [setEdges]
  );

  const onLayout = (direction: string, paramNodes?: ReactFlowData) => {
    const currentNodes = paramNodes?.nodes || nodes;
    const currentEdges = paramNodes?.edges || edges;

    const isHorizontal = direction === 'LR';
    dagreGraph.setGraph({ rankdir: 'TB' });

    currentNodes.forEach(node => {
      dagreGraph.setNode(node.id, { width: 150, height: 50 });
    });

    currentEdges.forEach(edge => {
      dagreGraph.setEdge(edge.source, edge.target);
    });

    dagre.layout(dagreGraph);

    const layoutedNodes = currentNodes.map(node => {
      const nodeWithPosition = dagreGraph.node(node.id);
      node.targetPosition = Position.Top;
      node.sourcePosition = Position.Bottom;
      // we need to pass a slightly different position in order to notify react flow about the change
      // @TODO how can we change the position handling so that we dont need this hack?
      node.position = {
        x: nodeWithPosition.x + Math.random() / 1000,
        y: nodeWithPosition.y
      };

      return node;
    });

    setNodes(layoutedNodes);
    setEdges(currentEdges);
  };

  const unselect = () => {
    setNodes(nds => nds.map(n => ({ ...n, selected: false })));
  };

  // const changeMarker = () => {
  //   setEdges(edges =>
  //     edges.map(e => ({
  //       ...e,
  //       markerEnd: {
  //         type:
  //           (e.markerEnd as EdgeMarker)?.type === MarkerType.Arrow
  //             ? MarkerType.ArrowClosed
  //             : MarkerType.Arrow
  //       }
  //     }))
  //   );
  // };

  // useEffect(() => {
  //   console.log(data);
  //   if (data?.nodes.length > 0) {
  //     onLayout('LR', data);
  //   }
  // }, [data]);

  return (
    <ReactFlowProvider>
      <div className="relative h-full flex-1">
        <div className="mt-8">
          <Panel position="top-right">
            <Button onClick={() => onLayout('TB')} size="xxs">
              vertical layout
            </Button>
            <Button onClick={() => onLayout('LR')} size="xxs">
              horizontal layout
            </Button>
            <Button onClick={() => unselect()} size="xxs">
              unselect nodes
            </Button>
            {/*            <Button onClick={() => changeMarker()} size="xxs">
              change marker
            </Button>*/}
            <Button onClick={() => fitView()} size="xxs">
              fitView
            </Button>
            <Button
              onClick={() => fitView({ nodes: nodes.slice(0, 2) })}
              size="xxs"
            >
              fitView partially
            </Button>
          </Panel>
        </div>
        <div className="h-[70vh]">
          <ReactFlow
            nodes={nodes}
            edges={edges}
            onConnect={onConnect}
            nodeTypes={nodeTypes}
            nodeExtent={nodeExtent}
            onInit={() => onLayout('TB')}
            onEdgesChange={onEdgesChange}
            onPaneScroll={undefined}
            nodesDraggable={false}
            nodesConnectable={true}
            edgeTypes={edgeTypes}
            connectionLineComponent={FloatingConnectionLine}
          >
            <Controls />
            {/*<Background variant='lines' gap={24} size={1} />*/}
          </ReactFlow>
        </div>
      </div>
    </ReactFlowProvider>
  );
};

export default ({ data }: LayoutFlowProps) => (
  <ReactFlowProvider>
    <LayoutFlow data={data} />
  </ReactFlowProvider>
);
