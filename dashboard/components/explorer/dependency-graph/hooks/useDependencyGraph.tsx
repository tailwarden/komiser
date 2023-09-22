import { useEffect, useState } from 'react';
import settingsService from '../../../../services/settingsService';

export type ReactFlowData = {
  nodes: any[];
  edges: any[];
};

// converting the json object into data that reactflow needs
// TODO - based on selected library
function GetData(res: any) {
  const d = {
    nodes: [],
    edges: []
  } as ReactFlowData;
  res.forEach((ele: any) => {
    // check if node exist already
    if (d.nodes.findIndex(element => element.id === ele.resourceId) === -1) {
      const a = {
        data: {
          label: ele.name,
          service: ele.service,
          provider: 'AWS',
          id: ele.resourceId,
          isRoot: true
        }
      };
      d.nodes.push(a);
    }

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: rel.resourceId,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}`,
          source: ele.resourceId,
          target: rel.resourceId,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
      };
      d.edges.push(edge);
    });

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: `${rel.resourceId}-1337`,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}-1337`,
          source: `${ele.resourceId}`,
          target: `${rel.resourceId}-1337`,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
      };
      d.edges.push(edge);
    });

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: `${rel.resourceId}-1338`,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}-1338`,
          source: `${ele.resourceId}`,
          target: `${rel.resourceId}-1338`,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
      };
      d.edges.push(edge);
    });

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: `${rel.resourceId}-1339`,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}-1339`,
          source: `${ele.resourceId}`,
          target: `${rel.resourceId}-1339`,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
      };
      d.edges.push(edge);
    });

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: `${rel.resourceId}-1336`,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic,
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}-1336`,
          source: `${ele.resourceId}`,
          target: `${rel.resourceId}-1336`,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
      };
      d.edges.push(edge);
    });

    ele.relations.forEach((rel: any) => {
      // check for other node exists
      if (d.nodes.findIndex(element => element.id === rel.resourceId) === -1) {
        const a = {
          data: {
            id: `${rel.resourceId}-1335`,
            label: rel.name,
            service: ele.service,
            type: rel.type,
            provider: 'AWS', // when supporting new provider this could be made dynamic
            isRoot: false
          }
        };
        d.nodes.push(a);
      }
      const edge = {
        data: {
          id: `${ele.resourceId}-${rel.resourceId}-1335`,
          source: `${ele.resourceId}`,
          target: `${rel.resourceId}-1335`,
          relation: rel.relation,
          label: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 20),
            Math.floor(Math.random() * 21) - 20
          ]
        }
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
