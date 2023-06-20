import ELK from 'elkjs';

const elk = new ELK();
const elkLayout = (nodes: any, edges: any) => {
  const nodesForElk = nodes.map((node: any) => ({
    id: node.id,
    width: 150,
    height: 120
  }));
  const graph = {
    id: 'root',
    layoutOptions: {
      'elk.algorithm': 'layered',
      'elk.direction': 'DOWN',
      'nodePlacement.strategy': 'SIMPLE'
    },

    children: nodesForElk,
    edges
  };
  return elk.layout(graph as any);
};

export default elkLayout;
