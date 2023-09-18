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
  useReactFlow
} from 'reactflow';
import { ReactFlowData } from './hooks/useDependencyGraph';
import Button from '../../button/Button';

const dagreGraph = new dagre.graphlib.Graph();
dagreGraph.setDefaultEdgeLabel(() => ({}));

const nodeExtent: CoordinateExtent = [
  [0, 0],
  [1000, 1000]
];

// In order to keep this example simple the node width and height are hardcoded.
// In a real world app you would use the correct width and height values of
// const nodes = useStoreState(state => state.nodes) and then node.__rf.width, node.__rf.height

const nodeWidth = 172;
const nodeHeight = 36;

const position = { x: 0, y: 0 };
const edgeType = 'smoothstep';

// const getLayoutedElements = (elements, direction = 'TB') => {
//   const isHorizontal = direction === 'LR';
//   dagreGraph.setGraph({ rankdir: 'TB' });

//   elements.forEach((el) => {
//     if (isNode(el)) {
//       dagreGraph.setNode(el.id, { width: nodeWidth, height: nodeHeight });
//     } else {
//       dagreGraph.setEdge(el.source, el.target);
//     }
//   });

//   dagre.layout(dagreGraph);

//   return elements.map((el) => {
//     if (isNode(el)) {
//       const nodeWithPosition = dagreGraph.node(el.id);
//       el.targetPosition = isHorizontal ? 'left' : 'top';
//       el.sourcePosition = isHorizontal ? 'right' : 'bottom';

//       // unfortunately we need this little hack to pass a slightly different position
//       // to notify react flow about the change. Moreover we are shifting the dagre node position
//       // (anchor=center center) to the top left so it matches the react flow node anchor point (top left).
//       el.position = {
//         x: nodeWithPosition.x - nodeWidth / 2 + Math.random() / 1000,
//         y: nodeWithPosition.y - nodeHeight / 2,
//       };
//     }

//     return el;
//   });
// };

type LayoutFlowProps = {
  data: ReactFlowData | undefined;
};

const LayoutFlow = ({ data }: LayoutFlowProps) => {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
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
    dagreGraph.setGraph({ rankdir: direction });

    currentNodes.forEach(node => {
      dagreGraph.setNode(node.id, { width: 150, height: 50 });
    });

    currentEdges.forEach(node => {
      dagreGraph.setEdge(node.source, node.target);
    });

    dagre.layout(dagreGraph);

    const layoutedNodes = currentNodes.map(node => {
      const nodeWithPosition = dagreGraph.node(node.id);
      node.targetPosition = isHorizontal ? Position.Left : Position.Top;
      node.sourcePosition = isHorizontal ? Position.Right : Position.Bottom;
      // we need to pass a slightly different position in order to notify react flow about the change
      // @TODO how can we change the position handling so that we dont need this hack?
      node.position = {
        x: nodeWithPosition.x + Math.random() / 1000,
        y: nodeWithPosition.y
      };

      return node;
    });

    setNodes(layoutedNodes);
  };

  const unselect = () => {
    setNodes(nds => nds.map(n => ({ ...n, selected: false })));
  };

  const changeMarker = () => {
    setEdges(eds =>
      eds.map(e => ({
        ...e,
        markerEnd: {
          type:
            (e.markerEnd as EdgeMarker)?.type === MarkerType.Arrow
              ? MarkerType.ArrowClosed
              : MarkerType.Arrow
        }
      }))
    );
  };

  useEffect(() => {
    console.log(data);
    if (data?.nodes.length > 0) {
      onLayout('LR', data);
    }
  }, [data]);

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
            <Button onClick={() => changeMarker()} size="xxs">
              change marker
            </Button>
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
            nodeExtent={nodeExtent}
            onInit={() => onLayout('TB')}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
          >
            <Controls />
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
