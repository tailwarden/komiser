/* eslint-disable react-hooks/exhaustive-deps */
import React, { useState, useCallback, useLayoutEffect } from 'react';
import ELK, { ElkExtendedEdge, ElkNode } from 'elkjs/lib/elk.bundled.js';
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
  EdgeTypes,
  Edge,
  Node,
  ResizeParamsWithDirection
} from 'reactflow';
import { ReactFlowData } from './hooks/useDependencyGraph';
import Button from '../../button/Button';
import CustomNode from './nodes/nodes';
import 'reactflow/dist/style.css';
import FloatingEdge from './nodes/FloatingEdge';
import FloatingConnectionLine from './nodes/FloatingConnectionLine';

const elk = new ELK();
const elkOptions = {
  // 'elk.algorithm': 'org.eclipse.elk.force',
  // 'elk.layered.spacing.nodeNodeBetweenLayers': '100',
  // 'elk.spacing.nodeNode': '80'
  'elk.algorithm': 'layered',
  'elk.spacing.nodeNode': '2',
  'elk.layered.spacing.nodeNodeBetweenLayers': '50',
  'elk.layered.spacing': '50',
  'elk.layered.mergeEdges': 'true',
  'elk.spacing': '50',
  'elk.spacing.individual': '50',
  'elk.edgeRouting': 'SPLINES'
};

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
} as unknown as EdgeTypes;

const LayoutFlow = ({ data }: LayoutFlowProps) => {
  const [nodes, setNodes, onNodesChange] = useNodesState(data.nodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(data.edges);
  const { fitView } = useReactFlow();

  const handleElkLayoutCallback = (layout: ElkNode) => {
    const nodePositions: { [key: string]: { x?: number; y?: number } } = {};

    layout.children?.forEach(node => {
      nodePositions[node.id] = { x: node.x, y: node.y };
    });

    const newNodes =
      layout.children?.map(node => {
        const nodePosition = nodePositions[node.id];
        return {
          ...node,
          position: {
            x: nodePosition?.x || 200,
            y: nodePosition?.y || 200
          }
        };
      }) || [];

    return { nodes: newNodes, edges: layout.edges };
  };

  const getLayoutedElements = (
    nodes: ElkNode[],
    edges: ElkExtendedEdge[],
    options = {}
  ) => {
    const graph = {
      id: 'root',
      layoutOptions: options,
      children: nodes.map(node => ({
        ...node,
        // Hardcode a width and height for elk to use when layouting.
        width: 200,
        height: 200
      })),
      edges: edges
    };

    return elk.layout(graph).then(handleElkLayoutCallback).catch(console.error);
  };

  useLayoutEffect(() => {
    onLayout({ direction: 'DOWN', useInitialNodes: true });
  }, []);

  const onConnect = useCallback(
    (params: Connection) =>
      setEdges(eds =>
        addEdge(
          {
            ...params,
            style: {
              stroke: '#33CCCC',
              strokeWidth: 1
            }
          },
          eds
        )
      ),
    []
  );
  const onLayout = useCallback(
    ({
      direction,
      useInitialNodes = false
    }: {
      direction: any;
      useInitialNodes: boolean;
    }) => {
      const opts = { 'elk.direction': direction, ...elkOptions };
      const ns = data.nodes;
      const es = data.edges;

      getLayoutedElements(ns, es, opts).then(elements => {
        // @ts-ignore
        setNodes(elements.nodes);
        // @ts-ignore
        setEdges(elements.edges);

        window.requestAnimationFrame(() => fitView());
      });
    },
    [data]
  );

  return (
    <ReactFlowProvider>
      <div className="relative h-full flex-1">
        <div className="mt-8">
          <Panel position="top-right">
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
            onPaneScroll={undefined}
            nodesDraggable={true}
            nodesConnectable={true}
          >
            <Controls />
            <Background color="#ccc" style={{ background: '#F4F9F9' }} />
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
