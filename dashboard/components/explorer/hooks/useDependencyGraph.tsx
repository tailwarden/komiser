import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';

import { InventoryFilterData } from '@components/inventory/hooks/useInventory/types/useInventoryTypes';
import settingsService from '@services/settingsService';

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
          provider: ele.provider,
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
            provider: ele.provider,
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
          label: `${rel.relation.toLowerCase()} ${rel.type}`,
          type: rel.type,
          controlPointDistances: [
            Math.floor(Math.random() * 30),
            Math.floor(Math.random() * 31) - 30
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
  const [filters, setFilters] = useState<InventoryFilterData[]>([]);
  const [displayedFilters, setDisplayedFilters] =
    useState<InventoryFilterData[]>();

  const router = useRouter();

  function fetch() {
    if (!loading) {
      setLoading(true);
    }

    if (error) {
      setError(false);
    }

    settingsService.getRelations(filters).then(res => {
      if (res === Error) {
        setLoading(false);
        setError(true);
      } else {
        setLoading(false);
        setData(GetData(res));
      }
    });
  }

  function deleteFilter(idx: number) {
    const updatedFilters: InventoryFilterData[] = [...filters!];
    updatedFilters.splice(idx, 1);
    const url = updatedFilters
      .map(
        filter =>
          `${filter.field}${`:${filter.operator}`}${
            filter.values.length > 0 ? `:${filter.values}` : ''
          }`
      )
      .join('&');
    router.push(url ? `?${url}` : '', undefined, { shallow: true });
  }

  const loadingFilters =
    Object.keys(router.query).length > 0 && !displayedFilters && !error;

  const hasFilters =
    Object.keys(router.query).length > 0 &&
    displayedFilters &&
    displayedFilters.length > 0;

  useEffect(() => {
    fetch();
  }, []);

  useEffect(() => {
    fetch();
  }, [filters, displayedFilters]);

  return {
    loading,
    data,
    error,
    fetch,
    filters,
    displayedFilters,
    setDisplayedFilters,
    deleteFilter,
    setFilters
  };
}

export default useDependencyGraph;
