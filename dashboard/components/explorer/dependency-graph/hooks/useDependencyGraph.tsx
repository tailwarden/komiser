import { useEffect, useState } from 'react';
import settingsService from '../../../../services/settingsService';

export type ReactFlowData = {
  nodes: any[];
  edges: any[];
};

// converting the json object into data that reactflow needs
function GetData(res: any) {
  const d = {
    nodes: [],
    edges: []
  } as ReactFlowData;
  res.forEach((ele: any) => {
    // check if node exist already
    if (d.nodes.findIndex(element => element.id === ele.resourceId) === -1) {
      const a = {
        id: ele.resourceId,
        type: 'customNode',
        data: { label: ele.service, resource: 'AWS' },
        position: { x: 0, y: 0 }
      };
      d.nodes.push(a);
    }
    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.ResourceID) === -1) {
        const a = {
          id: rel.ResourceID,
          type: 'customNode',
          data: { label: rel.Type, resource: 'AWS' }, // when supporting new provider this could be made dynamic
          position: { x: 0, y: 0 }
        };
        d.nodes.push(a);
      }
      const edge = {
        id: `${ele.resourceId}-${rel.ResourceID}`,
        source: ele.resourceId,
        target: rel.ResourceID
      };
      d.edges.push(edge);
    });
  });
  return d;
}

function useDependencyGraph() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<ReactFlowData>();
  const [error, setError] = useState(false);

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getRelations().then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setData(GetData(res));
      }
    });
  }

  useEffect(() => {
    fetch();
  }, []);

  return { loading, data, error, fetch };
}

export default useDependencyGraph;
